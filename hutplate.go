package hutplate

import (
	"net/http"
	"github.com/gorilla/sessions"
)

type Http struct {
	*http.Request
	Response response
	Session  session
	Auth     Auth
}

type configuration struct {
	GetUserWithId    func(id interface{}) interface{}
	GetUserWithCred  func(credential interface{}) (interface{}, string)
	ErrorHandler     func(err error, hut Http)
	SessionStore     *sessions.Store
	SessionSecretKey string
	SessionDirectory string
}

var Config configuration

func NewHttp(responseWriter http.ResponseWriter, request *http.Request) Http {
	session := NewSession(request, responseWriter, Config.SessionStore)
	r := newResponse(responseWriter, request, session)

	newHttp := Http {
		Request: request,
		Response: r,
		Session: session,
		Auth: Auth {
			session,
		},
	}

	return  newHttp
}