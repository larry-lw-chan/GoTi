package cookie

import (
	"encoding/gob"
	"log"
	"net/http"
)

// Flash Selections
const (
	SUCCESS = "success"
	FAILED  = "failed"
)

type Flash struct {
	Type    string
	Message string
}

func init() {
	gob.Register(Flash{})
}

func FlashGet(w http.ResponseWriter, r *http.Request) *Flash {
	session, _ := Store.Get(r, STORE)

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

func FlashSet(w http.ResponseWriter, r *http.Request, notice string, message string) {
	session, _ := Store.Get(r, STORE)
	flash := Flash{Type: notice, Message: message}
	session.AddFlash(flash)

	// Save the session
	err := session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
}
