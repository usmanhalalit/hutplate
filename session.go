package hutplate

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var store *sessions.FilesystemStore

type Session struct {
	request *http.Request
	responseWriter http.ResponseWriter
}

func BootSession () {
	store = sessions.NewFilesystemStore("", []byte(os.Getenv("APP_KEY")))
}

func (session Session) Set(key string, value interface{}) error {
	sessionStore, err := store.Get(session.request, "session")
	if err != nil {
		return err
	}

	sessionStore.Values[key] = value
	if err = store.Save(session.request, session.responseWriter, sessionStore); err != nil {
		return err
	}
	return nil
}

func (session Session) Get(key string) (interface{}, error) {
	sessionStore, err := store.Get(session.request, "session")
	if err != nil {
		return nil, err
	}
	return sessionStore.Values[key], nil
}

func (session Session) SetFlash(key string, value interface{}) error {
	sessionStore, err := store.Get(session.request, "session")
	if err != nil {
		return err
	}
	sessionStore.AddFlash(value, key)
	return store.Save(session.request, session.responseWriter, sessionStore)
}

func (session Session) GetFlash(key string) (interface{}) {
	sessionStore, err := store.Get(session.request, "session")
	if err != nil {
		return nil
	}
	flashes := sessionStore.Flashes(key)
	var flash interface{}

	if len(flashes) == 0 {
		flash = nil
	}  else {
		flash = flashes[0]
	}
	err = store.Save(session.request, session.responseWriter, sessionStore)
	return flash
}
