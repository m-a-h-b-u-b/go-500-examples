package main

import (
    "fmt"
    "net/http"
    "os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        file, header, err := r.FormFile("file")
        if err != nil {
            fmt.Fprintln(w, "Error:", err)
            return
        }
        defer file.Close()

        out, _ := os.Create(header.Filename)
        defer out.Close()
        _, _ = file.Seek(0, 0)
        _, _ = out.ReadFrom(file)

        fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
    } else {
        fmt.Fprintln(w, `<form method="POST" enctype="multipart/form-data">
            <input type="file" name="file">
            <input type="submit">
        </form>`)
    }
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
