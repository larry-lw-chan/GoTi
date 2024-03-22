package filestore

import (
	"errors"
	"fmt"
	"io"
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
	// Extract from file upload struct
	file := fu.File
	fh := fu.FileHeader
	dir := fu.Directory

	// Close the file when we finish
	defer file.Close()

	// Print out the file name and mime type
	fmt.Printf("Uploaded File: %+v\n", fh.Filename)
	fmt.Printf("File Size: %+v\n", fh.Size)
	fmt.Printf("MIME Header: %+v\n", fh.Header)

	// Get upload path
	uploadPath := ls.getPath(dir)

	// Generate name parttern from file extension
	name := strings.Split(fh.Filename, ".")
	pattern := name[0] + "-*." + name[1]

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp(uploadPath, pattern)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

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
