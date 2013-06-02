package controller

import (
	"path/filepath"
	"os"
	"net/http"
	"html/template"
	"fmt"
)


func ExecutableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}

func RedirectUtil(rw http.ResponseWriter, path string, attr map[string]interface {}) {
	binDir, err := ExecutableDir()
	if err != nil {
		panic(err)
	}
	if temp,err := template.ParseFiles(binDir + path); err == nil {
		temp.Execute(rw, attr)
	} else {
		fmt.Println("parse index.html error", err)
	}
}
