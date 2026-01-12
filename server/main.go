package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time/", timeHandle)
	mux.HandleFunc("/name/", valuesHandle)
	mux.HandleFunc("/", mainHandle)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}

func timeHandle(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
		http.Error(w, "method must be GET", http.StatusBadRequest)
        return
	}
	now := time.Now()
	today := now.Format("20060102 15:04:05")

	io.WriteString(w, today)
}

func valuesHandle(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
    email := r.PostFormValue("email")

	if len(name) == 0 || len(email) == 0 {
		http.Error(w, "provide name and email", http.StatusBadRequest)
        return
	}
	
    // Полезная работа с данными
    // ...
    fmt.Fprintf(w, "Данные успешно получены: %s - %s", name, email)
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method must be GET", http.StatusBadRequest)
        return
	}

	var out string

	if r.Header.Get("lang") == "ru" {
		out = fmt.Sprintf("Хост: %s\nМетод: %s\nURL: %s\n",
			r.Host, r.Method, r.URL.Path)
	} else {
		out = fmt.Sprintf("Host: %s\nMethod: %s\nURL: %s\n",
			r.Host, r.Method, r.URL.Path)
	}

	io.WriteString(w, out)
}
