package router

import (
	"net/http"

	"github.com/xfp-discerning/views"
)

func Router() {
	//根据路由的请求类型可分类
	//1、页面 views 2、数据（json）3、静态资源
	http.HandleFunc("/", views.HTML.Index)
	//添加路由
	//http.FileServer将开启的服务作一个本地的文件服务器
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
