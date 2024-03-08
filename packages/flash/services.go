package flash

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/larry-lw-chan/goti/packages/cookie"
)

// Flash Notification Types
const (
	SUCCESS = "success"
	NOTICE  = "notice"
	ALERT   = "alert"
	ERROR   = "error"
)

type Flash struct {
	Type    string
	Message string
}

func init() {
	gob.Register(Flash{})
}

func Get(w http.ResponseWriter, r *http.Request) *Flash {
	session, _ := cookie.Store.Get(r, cookie.STORE)

	// Get flash value
	flashes := session.Flashes()
	if len(flashes) == 0 {
		return &Flash{}
	}
	flash := flashes[0].(Flash)

	// Save the session
	err := session.Save(r, w)
	if err != nil {
		log.Println(err)
	}

	return &flash
}

func Set(w http.ResponseWriter, r *http.Request, notice string, message string) {
	session, _ := cookie.Store.Get(r, cookie.STORE)
	flash := Flash{Type: notice, Message: message}
	session.AddFlash(flash)

	// Save the session
	err := session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
}
