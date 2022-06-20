package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/xfp-discerning/models"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//找到这个文件
	path, _ := os.Getwd()
	home := path + `\template\home.html`
	header := path + `\template\layout\header.html`
	footer := path + `\template\layout\footer.html`
	personal := path + `\template\layout\personal.html`
	post := path + `\template\layout\post.html`
	pagination := path + `\template\layout\pagination.html`
	//访问博客首页时，有多个模板的嵌入，所以需要解析所涉及到的所有模板
	t, err := t.ParseFiles(path+`\template\index.html`, home, header, footer, personal, post, pagination) //此处不是path + "\template\index.html"
	if err != nil {
		log.Println("页面解析失败")
	}
	//页面上所涉及的所有数据，必须有定义
	// viewer := config.Viewer{
	// 	Title: "xfp",
	// 	Description: "Blog",
	// 	Logo: ,
	// }
	var hr = &models.HomeResponse{}
	t.Execute(w, hr)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)

	}
}
