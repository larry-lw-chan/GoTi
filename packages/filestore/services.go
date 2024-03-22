package filestore

import (
	"mime/multipart"
)

/********************************************************
* Filestore Service
*********************************************************/
// Global Variable to hold the selected filestore (default local)
var FS FilestoreService = LocalStore{
	LocalFolder: "uploads",
}

// Interface used to define the filestore service
type FilestoreService interface {
	Upload(fu FileUpload) (string, error)
	Delete(filename string) error
}

// Struct used to hold the needed file parameters
type FileUpload struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Directory  string
}

// Uploads a file to the selected filestore
func Upload(fu FileUpload) (url string, err error) {
	uploadPath, err := FS.Upload(fu)
	return uploadPath, err
}

// Delete a file from the selected filestore
func Delete(filename string) (err error) {
	err = FS.Delete(filename)
	return err
}
