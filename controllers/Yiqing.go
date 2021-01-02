package controllers

import (
	"api.cb2020.tdedu.org/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type YiqingController struct {
	ApiBase
}

func (this *YiqingController) GetCityList() {
	//验证IP
	checkIPRes := checkIP()
	if checkIPRes == false {
		this.ResultJson(10102, "验证ip不通过")
	}

	//验证参数
	secret := this.GetString("secret")
	if secret == "" {
		this.ResultJson(10101, "secret不能为空")

	} else if secret != beego.AppConfig.String("secret") {
		this.ResultJson(10102, "secret错误")
	}

	//接受参数
	id, err := this.GetInt("id")

	if err != nil {
		this.ResultJson(10105, "参数有问题")
	}

	o := orm.NewOrm()
	yiqing := models.YiQingCity{}
	yiqing.Id = id

	yiqingerr := o.Read(&yiqing)
	if yiqingerr != nil {
		this.ResultJson(10104, "数据库查询失败")
	} else {
		this.ResultJson(0, "成功", yiqing)
	}
	//this.ServeJSON()
}

