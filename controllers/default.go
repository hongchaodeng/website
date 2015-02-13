package controllers

import (
	"fmt"

	"github.com/gopherchina/website/models"
)

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	name := this.Ctx.Input.Param(":name")
	if name == "" {
		this.Data["indexActive"] = true
		this.TplNames = "index.tpl"
	} else if name == "speaker" {
		this.Data["userActive"] = true
		this.TplNames = "speaker.tpl"
	} else {
		df := models.GetDoc(name, this.Lang)
		this.Data[fmt.Sprintf("%sActive", name)] = true
		this.Data["Section"] = name
		this.Data["Title"] = df.Title
		this.Data["title"] = df.Title + " - "
		this.Data["Data"] = string(df.Data)
		this.TplNames = "detail.tpl"
	}
}