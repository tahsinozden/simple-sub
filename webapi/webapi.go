package webapi

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"io"
	"net/http"
	"simple-sub/utils"
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

func Serve(c utils.CommandArgs) {
	p := c.Port
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/upload", receiveFile)
	fmt.Printf("serving on port %s...\n", p)
	http.ListenAndServe(":"+p, nil)
}
