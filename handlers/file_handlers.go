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

	filePath := filepath.Join(uploadDir, handler.Filename)
	if _, err := os.Stat(filePath); err == nil {
		helpers.ErrorJSON(w, http.StatusConflict, fmt.Errorf("file already exists"))
		return
	}

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

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	http.ServeFile(w, r, filePath)
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		if os.IsNotExist(err) {
			helpers.WriteJSON(w, http.StatusOK, []string{})
			return
		}
		helpers.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error reading upload directory: %w", err))
		return
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			files = append(files, entry.Name())
		}
	}

	helpers.WriteJSON(w, http.StatusOK, files)
}
