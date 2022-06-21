package views

import (
	"net/http"

	"github.com/xfp-discerning/common"
	"github.com/xfp-discerning/config"
	"github.com/xfp-discerning/models"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上所涉及的所有数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "徐发鹏",
			ViewCount:    123,
			CreateAt:     "2022-06-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	index.WriteData(w, hr)
}
