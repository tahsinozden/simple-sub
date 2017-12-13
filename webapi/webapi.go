package webapi

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"io"
	"net/http"
	"tubtitle/utils"
	"tubtitle/subutils"
)

func receiveFile(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("File name %s\n", header.Filename)
	io.Copy(&buffer, file)
	contents := buffer.String()
	fmt.Println(contents)
}

func removeAccentLetters(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()

	f := header.Filename
	fmt.Printf("File name %s\n", f)
	io.Copy(&buffer, file)
	contents := buffer.String()
	fmt.Println(subutils.GetConvertAccentText(contents))

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.accents-removed.srt", f))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	// writes to response, TODO: create a download link
	w.Write([]byte(subutils.GetConvertAccentText(contents)))
}

func Serve(c utils.CommandArgs) {
	p := c.Port
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/upload", removeAccentLetters)
	fmt.Printf("serving on port %s...\n", p)
	http.ListenAndServe(":"+p, nil)
}
