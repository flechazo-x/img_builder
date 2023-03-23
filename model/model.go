package model

import (
	"github_project/web/model/img"
	"html/template"
	"net/http"
)

func Tpl(w http.ResponseWriter, _ *http.Request) {
	t := template.New("some template") //创建一个模板
	t, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	image, err := img.GetImage()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	p := struct {
		Img string
	}{}
	p.Img = image
	// 渲染模板
	err = t.Execute(w, p)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func Index(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("index.html")
	_ = t.Execute(w, nil)
}
