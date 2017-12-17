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

var validModes = map[string]func(w http.ResponseWriter, r *http.Request){
	"REMOVE_LETTERS": removeAccentLetters,
	"MERGE_SUB":      mergeSubtitles,
}

func subtitleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		glog.Warning(r.Method + " method is not supported!")

	case "POST":
		params := parseRequest(w, r)
		ops := params["operation"]

		if len(ops) == 0 {
			glog.Error("No operations provided!")
			http.Redirect(w, r, "../error.html", 301)
		}

		if fn, ok := validModes[ops[0]]; ok {
			fn(w, r)

		} else {
			glog.Error(ops[0] + " is not supported!")
			http.Redirect(w, r, "../error.html", 301)
		}
	default:
		fmt.Println("here")
		glog.Warning(r.Method + " method is not supported!")
		http.Redirect(w, r, "../error.html", 301)
	}
}

func removeAccentLetters(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	fileBottom, headerBottom, err := r.FormFile("subBottom")
	if err != nil {
		glog.Error(err)
		return
	}
	defer fileBottom.Close()

	enc := r.Form["encBottom"]
	if len(enc) == 0 {
		glog.Warning("No encoding found!")
		http.Redirect(w, r, "../error.html", 301)
	}
	e := enc[0]
	f := headerBottom.Filename
	fmt.Printf("File name : %s\n", f)

	buf = utils.CopyWithEncoding(fileBottom, subutils.GetEncoding(e))
	contents := buf.String()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.accents-removed.srt", f))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Write([]byte(subutils.GetConvertAccentText(contents)))
}

func mergeSubtitles(w http.ResponseWriter, r *http.Request) {
	// TODO: implement it
}

func parseRequest(w http.ResponseWriter, r *http.Request) map[string][]string {
	r.ParseMultipartForm(10000000)
	fmt.Println("=== Request Body Parameters")
	for key, value := range r.PostForm {
		fmt.Printf("%s : %s\n", key, value)
	}
	return r.Form
}

func Serve(c utils.CommandArgs) {
	p := c.Port
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/subs", subtitleHandler)
	fmt.Printf("serving on port %s...\n", p)
	http.ListenAndServe(":"+p, nil)
}
