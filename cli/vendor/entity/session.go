package entity

import (
	"fmt"
	"net/http"
)

// Session model for one login session
type Session struct {
	Openid string `json:"openid"`
}

type sessionDb struct {
	CurUser Session `json:"curUser"`
}

type sessionModel struct {
	storage
	session *Session
}

var (
	// CurSessionModel model for current session
	CurSessionModel sessionModel
)

func init() {
	addModel(&CurSessionModel, "curUser")
}

// Init initialize a session model
func (model *sessionModel) Init(path string) {
	logger.Println("[sessionmodel] initializing")
	model.path = path
	model.session = &Session{}

	model.load()
	logger.Println("[sessionmodel] initialized")
}

// SetCurOpenid sets current openid in the session
func (model *sessionModel) SetCurOpenid(openid string) {
	logger.Printf("[sessionmodel] try setting openid '%s' to current session\n", openid)
	model.session.Openid = openid
	model.dump()
	logger.Printf("[sessionmodel] set openid '%s' to current session\n", openid)
}

// GetCurOpenid get current openid
func (model *sessionModel) GetCurOpenid() string {
	return model.session.Openid
}

// GetCurUser is kept for compilation
func (model *sessionModel) GetCurUser() string {
	return ""
}

// SetCurUser is kept for compilation
func (model *sessionModel) SetCurUser() {
	return
}

func (model *sessionModel) load() {
	var sessionDb sessionDb
	model.storage.load(&sessionDb)
	model.session = &sessionDb.CurUser
}

func (model *sessionModel) dump() {
	var sessionDb sessionDb
	sessionDb.CurUser = *model.session
	model.storage.dump(&sessionDb)
}

// Login performs login request
func Login(username string, password string) (openid string, err error) {
	logger.Println("[sessionentity] try performing login request with username = ", username)
	reqBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{username, password}

	var resBody struct {
		Openid string `json:"openid"`
		Msg    string `json:"msg"`
	}

	var code int
	if code, err = request("POST", "/api/user/login", &reqBody, &resBody); err != nil {
		return
	}
	if code == http.StatusOK {
		openid = resBody.Openid
		return
	}
	err = fmt.Errorf("%s", resBody.Msg)
	return
}

// Logout performs logout request
func Logout() (err error) {
	logger.Println("[sessionentity] try performing logout request")
	var code int
	if code, err = request("POST", "/api/user/logout", nil, nil); err != nil {
		return
	}
	if code == http.StatusOK {
		return
	}
	err = fmt.Errorf("%d", code)
	return
}

// GetCurUsername performs request to fetch the username associated with current session
func GetCurUsername() (username string, err error) {
	logger.Println("[sessionentity] try performing request for whether current session is logged in")
	var body struct {
		Username string `json:"username"`
	}
	var code int
	if code, err = request("GET", "/api/user/login", nil, &body); err != nil {
		return
	}
	if code == http.StatusOK {
		username = body.Username
	}
	return
}
