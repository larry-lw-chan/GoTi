package filestore

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/********************************************************
* Local Filestore
*********************************************************/
var LocalDefaultFolder string = "uploads"

type LocalStore struct {
	LocalFolder string
}

// Main function that uploads a file to the local filestore
func (ls LocalStore) Upload(fu FileUpload) (string, error) {
	// Get file parameters
	fileBytes := fu.FileBytes
	namePattern := fu.NamePattern
	dir := fu.Directory

	// Get upload path
	dirList := strings.Split(dir, "/")
	uploadPath := ls.getPath(dirList...)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp(uploadPath, namePattern)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// Return the path and file name
	return "/" + tempFile.Name(), nil
}

func (ls LocalStore) Delete(filename string) error {
	// Make sure file is not empty
	if filename == "" {
		return errors.New("no file provided for deletion")
	}

	// Remove the first slash from filename
	filename = filename[1:]

	// Delete the file
	err := os.Remove(filename)
	return err
}

// GetUploadLocation function returns the path where files will be uploaded
func (ls LocalStore) getPath(dir ...string) string {
	// Create main uploads directory if it does not exist
	if ls.LocalFolder == "" {
		ls.LocalFolder = LocalDefaultFolder
	}

	// Create main uploads directory if it does not exist
	if _, err := os.Stat(ls.LocalFolder); os.IsNotExist(err) {
		os.Mkdir(ls.LocalFolder, 0755)
	}

	// Create the sub-directory if it does not exist
	uploadPath := ls.LocalFolder
	for _, d := range dir {
		uploadPath = uploadPath + "/" + d
		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			os.Mkdir(uploadPath, 0755)
		}
	}

	return uploadPath
}
