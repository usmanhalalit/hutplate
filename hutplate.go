package hutplate

import (
	"net/http"
)

type Http struct {
	*http.Request
	Response Response
	Session  session
	Auth     Auth
}

type configuration struct {
	GetUserWithId func(id interface{}) interface{}
	GetUserWithCred func(credential interface{}) (interface{}, string)
	HandleError func(err error, hut Http)
}

var Config configuration

func NewHttp(responseWriter http.ResponseWriter, request *http.Request) Http {
	session := NewSession(request, responseWriter, nil)

	newHttp := Http {
		request,
		Response {
			responseWriter,
			request,
			session,
		},
		session,
		Auth {
			session,
		},
	}

	return  newHttp
}