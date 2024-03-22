package filestore

import (
	"errors"
	"mime/multipart"
)

/********************************************************
* Filestore Service
*********************************************************/

// Global Variables
const (
	Local = "local_store"
)

// Global Variables that determines filestore type
var Filestore string = Local

// Struct used to hold the needed file parameters
type FileUpload struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Username   string
	Directory  string
}

// Public Function that uploads a file to the selected filestore
func Upload(fu FileUpload) (url string, err error) {
	if Filestore == Local {
		uploadPath, err := LocalUpload(fu)
		return uploadPath, err
	}
	return "", errors.New("filestore not supported")
}
