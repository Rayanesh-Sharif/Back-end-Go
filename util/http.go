package util

import (
	"github.com/go-faster/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var ErrMimeMismatch = errors.New("mime mismatch")

func CheckMimeAndSaveFile(file *multipart.FileHeader, destination, expectedMime string) error {
	// From context.SaveUploadedFile
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Check the mime type. To do so, we read the first 512 bytes into array
	fileHeader := make([]byte, 512)
	n, err := src.Read(fileHeader)
	if err != nil {
		return err
	}
	fileHeader = fileHeader[:n]
	// https://stackoverflow.com/a/38175140/4213397
	if http.DetectContentType(fileHeader) != expectedMime {
		return ErrMimeMismatch
	}
	// Write to file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the buffer at first
	_, err = out.Write(fileHeader)
	if err != nil {
		return err
	}
	// Copy the image to file on disk
	_, err = io.Copy(out, src)
	return err
}
