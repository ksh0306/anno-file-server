package handler

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}

// https://github.com/mimoo/eureka/blob/master/folders.go
func decompressTar(f *os.File, dst string) error {
	fmt.Println("decompress tar: ", f.Name())

	tr := tar.NewReader(f)

	// uncompress each element
	for {
		header, err := tr.Next()
		if err == io.EOF {
			fmt.Println("io.EOF")
			break // End of archive
		}
		if err != nil {
			return err
		}
		target := header.Name

		// validate name against path traversal
		if !validRelPath(header.Name) {
			return fmt.Errorf("tar contained invalid name error %q", target)
		}

		// add dst + re-format slashes according to system
		target = filepath.Join(dst, header.Name)
		// if no join is needed, replace with ToSlash:
		// target = filepath.ToSlash(header.Name)

		// check the type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it (with 0755 permission)
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		// if it's a file create it (with same permission)
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			// copy over contents
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			fileToWrite.Close()
		}
	}

	if err := os.Remove(f.Name()); err != nil {
		log.Println(err)
	}

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
	decompressTar(dst, "./uploaded")

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
