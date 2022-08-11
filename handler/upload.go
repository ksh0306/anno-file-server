package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func decompress(f *os.File) error {
	return nil
}

func Upload(c echo.Context) error {

	// 1. receive file
	log.Println("upload")
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("FormFile")
		return err
	}
	src, err := file.Open()
	if err != nil {
		log.Println("Open")
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join("./uploaded", file.Filename))
	if err != nil {
		log.Println("create ", file.Filename)
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println("copy")
		return err
	}

	// decompress file
	decompress(dst)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
