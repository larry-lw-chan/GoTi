package flash

import (
	"encoding/base64"
	"net/http"
	"time"
)

// Flash types
const (
	SUCCESS = "success"
	FAILED  = "failed"
)

type Flash struct {
	Type    string
	Message string
}

func Set(w http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{Name: name, Value: encode(value)}
	http.SetCookie(w, c)
}

func Get(w http.ResponseWriter, r *http.Request, name string) *Flash {
	flash := &Flash{}
	value, _ := GetFlashValue(w, r, name)
	if value != nil {
		flash.Type = name
		flash.Message = string(value)
	}
	return flash
}

func GetFlashValue(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	// Got value from cookie so expiring it not to show again
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)

	return value, nil
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
