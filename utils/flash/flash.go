package flash

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Flash types
const NAME = "flash"

const (
	SUCCESS = "success"
	FAILED  = "failed"
)

type Flash struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func Set(w http.ResponseWriter, flashType string, message []byte) {
	flash := &Flash{flashType, string(message)}
	value, err := json.Marshal(flash)
	if err != nil {
		log.Println(err)
	}

	c := &http.Cookie{Name: NAME, Value: encode(value)}
	http.SetCookie(w, c)
}

func Get(w http.ResponseWriter, r *http.Request) *Flash {
	flash := &Flash{}
	value, _ := GetFlashValue(w, r)
	if value != nil {
		flash.Type = value.Type
		flash.Message = value.Message
	}
	return flash
}

func GetFlashValue(w http.ResponseWriter, r *http.Request) (*Flash, error) {
	// See if cookie exists
	c, err := r.Cookie(NAME)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}

	// Decode the value
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}

	// Unmarshal the value
	var flash Flash
	err = json.Unmarshal(value, &flash)
	if err != nil {
		return nil, err
	}

	// Got value from cookie so expiring it not to show again
	dc := &http.Cookie{Name: NAME, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)

	return &flash, nil
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
