package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func GetUploadLocation(dir ...string) string {
	uploadDir := os.TempDir()
	return uploadDir
}

func File(file multipart.File, fh *multipart.FileHeader, uploadPath string) (url string, err error) {
	// Print out the file name and mime type
	fmt.Printf("Uploaded File: %+v\n", fh.Filename)
	fmt.Printf("File Size: %+v\n", fh.Size)
	fmt.Printf("MIME Header: %+v\n", fh.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp(uploadPath, fh.Filename)
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
