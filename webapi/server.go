package webapi

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"tubtitle/subutils"
	"tubtitle/utils"
)

type SubtitleForm struct {
	Name string
	File bytes.Buffer
	Enc  string
}

func removeAccentLetters(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	//fmt.Println(r)

	fileBottom, headerBottom, err := r.FormFile("fileBottom")
	if err != nil {
		glog.Error(err)
		return
	}
	defer fileBottom.Close()

	r.ParseForm()
	enc := r.Form["encBottom"]
	if len(enc) == 0 {
		glog.Warning("No encoding found!")
		return
	}
	e := enc[0]

	f := headerBottom.Filename
	fmt.Printf("File name %s\n", f)

	buf = utils.CopyWithEncoding(fileBottom, subutils.GetEncoding(e))
	contents := buf.String()
	//fmt.Println(subutils.GetConvertAccentText(contents))

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
