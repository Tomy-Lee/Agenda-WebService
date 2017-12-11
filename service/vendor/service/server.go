package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.Use(&checkLoginHandler{})
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// Group User
	mx.HandleFunc("/api/user/login", checkIsLoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/user/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/user/logout", logoutHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/user/self", deleteAccountHandler(formatter)).Methods("DELETE")

	// Group Users
	mx.HandleFunc("/api/users", registerHandler(formatter)).Methods("POST")
}
