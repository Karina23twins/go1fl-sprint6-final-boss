package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleReturnHtml(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("error: The GET method is expected")
		http.Error(w, "error: The GET method is expected", http.StatusInternalServerError)
		return
	}

	// if _, err := os.Stat("/home/arina/go1fl-sprint6-final-boss/index.html"); os.IsNotExist(err) {
		if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		log.Println("error: file is not exist")
		http.Error(w, "error: file is not exist", http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Println("error checking file: ", err)
		http.Error(w, "error checking file", http.StatusInternalServerError)
		return
	}

	// http.ServeFile(w, r, "/home/arina/go1fl-sprint6-final-boss/index.html")
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Println("error parsing form: ", err)
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}

	// получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("error receiving the file", err)
		http.Error(w, "error receiving the file", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()

	// читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file:", err)
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	// передаем эти данные в функцию автоопределения из пакета service
	convertedString, err := service.ConvertString(string(data))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// создаём локальный файл
	localFileName := time.Now().UTC().String() + filepath.Ext(handler.Filename)
	// localFileName := time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)
	localFile, err := os.Create(localFileName)
	if err != nil {
		log.Println("error creating file:", err)
		http.Error(w, "error creating file", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	// записываем в локальный файл конвертированную строку
	_, err = localFile.Write([]byte(convertedString))
	if err != nil {
		log.Println("error writing to file:", err)
		http.Error(w, "error writing to file", http.StatusInternalServerError)
		return
	}

	//	возврат результата конвертации строки:
	fmt.Fprintln(w, convertedString)
}
