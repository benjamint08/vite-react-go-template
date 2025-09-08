package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/benjamint08/vite-react-go-template/helpers"
)

const uploadDir = "uploads"

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("could not create upload directory: %w", err))
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		helpers.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("error retrieving the file: %w", err))
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
	if err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error creating the file: %w", err))
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error saving the file: %w", err))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"message": "File uploaded successfully"})
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		helpers.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("file parameter is required"))
		return
	}

	filePath := filepath.Join(uploadDir, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
