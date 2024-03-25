package flash

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/larry-lw-chan/goti/internal/sessions/cookie"
)

// Flash Notification Types
const (
	SUCCESS string = "success"
	NOTICE  string = "notice"
	ALERT   string = "alert"
	ERROR   string = "error"
)

type Flash struct {
	Type    string
	Message string
}

func init() {
	gob.Register(Flash{})
}

func Get(w http.ResponseWriter, r *http.Request) *Flash {
	session, err := cookie.Store.Get(r, cookie.STORE)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Get flash value
	flashes := session.Flashes()
	if len(flashes) == 0 {
		return nil
	}
	flash := flashes[0].(Flash)

	// Save the session
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}

	return &flash
}

func Set(w http.ResponseWriter, r *http.Request, notice string, message string) {
	// Get the session
	session, err := cookie.Store.Get(r, cookie.STORE)
	if err != nil {
		log.Println(err)
		return
	}

	// Add a flash message
	flash := Flash{Type: notice, Message: message}
	session.AddFlash(flash)

	// Save the session
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
}
