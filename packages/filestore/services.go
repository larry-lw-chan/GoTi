package filestore

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// UploadPath is the path where files will be uploaded
var UploadPath = "uploads"

func init() {
	// Override the default port with user configuration
	if os.Getenv("UPLOAD_PATH") != "" {
		UploadPath = os.Getenv("UPLOAD_PATH")
	}

	// Create the uploads directory if it does not exist
	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		os.Mkdir(UploadPath, 0755)
	}
}

// GetUploadLocation function returns the path where files will be uploaded
func GetUploadPath(dir ...string) string {
	uploadPath := UploadPath

	// Create the uploads directory if it does not exist
	for _, d := range dir {
		uploadPath = uploadPath + "/" + d
		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			os.Mkdir(uploadPath, 0755)
		}
	}

	return uploadPath
}

// Upload a file to the server
func Upload(file multipart.File, fh *multipart.FileHeader, uploadPath string) (url string, err error) {
	// Close the file when we finish
	defer file.Close()

	// Print out the file name and mime type
	fmt.Printf("Uploaded File: %+v\n", fh.Filename)
	fmt.Printf("File Size: %+v\n", fh.Size)
	fmt.Printf("MIME Header: %+v\n", fh.Header)

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
