package service

import (
	"encoding/json"
	"net/http"
	"utils"

	"entities"

	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

var (
	nonLoginAPI = map[string]bool{
		"/api/user/login": true,
		"/api/users":      true,
	}
)

type checkLoginHandler struct{}

func (h *checkLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, isNonLogin := nonLoginAPI[r.URL.Path]
	if !isNonLogin && checkIsLogin(r) == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		next(w, r)
	}
}

func checkIsLogin(r *http.Request) *entities.User {
	openidCookie, err := r.Cookie("openid")
	// no cookie
	if err != nil {
		return nil
	}
	openid := openidCookie.Value
	session := entities.SessionServ.FindByOpenid(openid)
	// openid not found
	if session == nil {
		return nil
	}

	user := entities.UserServ.FindByUsername(session.Username)
	// user not found
	if user == nil {
		return nil
	}
	return user
}

func checkIsLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if user := checkIsLogin(r); user != nil {
			formatter.JSON(w, http.StatusOK, user)
		} else {
			formatter.JSON(w, http.StatusUnauthorized, struct{}{})
		}
	}
}

func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var badData struct {
			Msg string `json:"msg"`
		}
		if user := checkIsLogin(r); user != nil {
			badData.Msg = user.Username + ", please log out first"
			formatter.JSON(w, http.StatusBadRequest, badData)
			return
		}

		var reqBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		user := entities.UserServ.FindByUsername(reqBody.Username)
		if user == nil || !IsPwdMatch(user, reqBody.Password) {
			badData.Msg = "Incorrent username and password combination"
			formatter.JSON(w, http.StatusUnauthorized, badData)
			return
		}

		var openid string
		for {
			openid = uuid.NewV4().String()
			if entities.SessionServ.FindByOpenid(openid) == nil {
				break
			}
		}
		session := entities.Session{
			Openid:   openid,
			Username: reqBody.Username,
		}
		entities.SessionServ.Add(&session)
		formatter.JSON(w, http.StatusOK, struct {
			Openid string `json:"openid"`
			Msg    string `json:"msg"`
		}{
			Openid: openid,
			Msg:    "Logged in successfully",
		})
	}
}

func logoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		openidCookie, _ := r.Cookie("openid")
		openid := openidCookie.Value
		entities.SessionServ.Delete(openid)
		w.WriteHeader(http.StatusOK)
	}
}

// IsPwdMatch checks whether provided password matches that of the user
func IsPwdMatch(user *entities.User, password string) bool {
	return utils.MD5(password) == user.Password
}
