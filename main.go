package main

import (
	"bytes"
	"fmt"
	"github.com/tienne/rockets-go-tutorial/seam"
	"io"
	"log"
	"net/http"
)

const (
	IMAGE_URL string = "https://bit.ly/2QGPDkr"
)

func main() {
	fmt.Println("Ready for liftoff! Checkout http://localhost:3000/occupymars")

	http.HandleFunc("/occupymars", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("resize") > "" {
			resized, err := seam.ContentAwareResize(IMAGE_URL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "image/jpeg")
			io.Copy(w, bytes.NewReader(resized))
		} else {
			fmt.Fprintf(w, "<html><div>Original image:</div> <img src=\"%s\" /><br/><a href=\"?resize=1\">Resize using Seam Carving</a></html>", IMAGE_URL)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
