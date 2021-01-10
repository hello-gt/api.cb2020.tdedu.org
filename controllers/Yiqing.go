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

	o := orm.NewOrm()
	yiqing := models.YiQingCity{}
	if err != nil {
		this.ResultJson(10105, "参数有问题")
	} else{
		yiqing.Id = id
		//beego.Info("id:",id)
	}
	country_id, err := this.GetInt("countrie_id")

	if err != nil {
		this.ResultJson(10105, "参数有问题")
	} else{
		beego.Info("country_id:",country_id)
		yiqing.Countrie_id = country_id
	}

	yiqingerr := o.Read(&yiqing)
	beego.Info("yiqing:",yiqing)
	if yiqingerr != nil {
		this.ResultJson(10104, "数据库查询失败")
	} else {
		this.ResultJson(0, "成功", yiqing)
	}
	//this.ServeJSON()
}

