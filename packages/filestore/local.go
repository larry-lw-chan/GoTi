package filestore

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/********************************************************
* Local Filestore
*********************************************************/
// Global Variables
var LocalFolder string = "uploads"

// Main function that uploads a file to the local filestore
func LocalUpload(fu FileUpload) (string, error) {
	// Extract from file upload struct
	file := fu.File
	fh := fu.FileHeader
	username := fu.Username
	dir := fu.Directory

	// Close the file when we finish
	defer file.Close()

	// Print out the file name and mime type
	fmt.Printf("Uploaded File: %+v\n", fh.Filename)
	fmt.Printf("File Size: %+v\n", fh.Size)
	fmt.Printf("MIME Header: %+v\n", fh.Header)

	// Get upload path
	uploadPath := getLocalPath(username, dir)

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

// GetUploadLocation function returns the path where files will be uploaded
func getLocalPath(dir ...string) string {
	uploadPath := LocalFolder

	// Create the uploads directory if it does not exist
	for _, d := range dir {
		uploadPath = uploadPath + "/" + d
		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			os.Mkdir(uploadPath, 0755)
		}
	}

	return uploadPath
}

// Allows the local folder override
func LocalFolderOverride(folder string) {
	LocalFolder = folder

	// Create the uploads directory if it does not exist
	if _, err := os.Stat(LocalFolder); os.IsNotExist(err) {
		os.Mkdir(LocalFolder, 0755)
	}
}
