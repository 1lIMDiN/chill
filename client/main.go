package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
    baseURL := "http://localhost:8080/"
    resp, err := http.Get(baseURL+"time")
    if err != nil {
        fmt.Println(err)
        return
    }

    print(resp)

    client := &http.Client{
        Timeout: 1 * time.Second,
    }

    req, err := http.NewRequest(http.MethodGet, baseURL, http.NoBody)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Header.Set("lang", "ru")

    resp, err = client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }

    print(resp)

    params := url.Values{
        "name": {"Aleks"},
        "email": {"2@sds.ru"},
    }

    resp, err = http.PostForm(baseURL+"name/", params)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    print(resp)
}

func print(w *http.Response) {
    body, err := io.ReadAll(w.Body)
    w.Body.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}