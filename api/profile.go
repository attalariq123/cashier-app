package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	imageName := "img-avatar.png"                                     // mengambil nama image dari query url
	fileBytes, err := ioutil.ReadFile("./assets/images/" + imageName) // membaca file image menjadi bytes
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("File not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes) // menampilkan image sebagai response
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	alias := "img-avatar"

	uploadedFile, handler, err := r.FormFile("file-avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	//mengambil relative path dari proyek
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//membuat nama file baru yang akan disimpan
	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, ".png")
	}

	//membentuk lokasi tempat menyimpan file
	fileLocation := filepath.Join(dir, "assets/images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	//mengisi file baru dengan data dari file yang ter-upload
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.dashboardView(w, r)
}
