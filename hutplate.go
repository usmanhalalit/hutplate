package hutplate

import (
	"net/http"
)

type Http struct {
	*http.Request
	Response Response
	Session Session
	Auth Auth
}

type Configuration struct {
	GetUserWithId func(id interface{}) interface{}
	GetUserWithCred func(credential interface{}) (interface{}, string)
}

type Handler func(hp Http) interface{}

var Config Configuration

func Boot() {
	BootSession()
}

func NewHttp(responseWriter http.ResponseWriter, request *http.Request) Http {
	session := Session {
		request,
		responseWriter,
	}

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

func (handleFunc Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handleFunc(NewHttp(w, r))
	switch response.(type) {
	case error:
		http.Error(w, response.(error).Error(), 500)
	case string:
		w.Write([]byte(response.(string)))
	default:

	}
}