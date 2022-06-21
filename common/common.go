package common

import (
	"github.com/xfp-discerning/config"
	"github.com/xfp-discerning/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir+"/template/")
}