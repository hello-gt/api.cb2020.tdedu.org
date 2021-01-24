package controllers

import (
	"api.cb2020.tdedu.org/models/yi_qing_city"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type YiqingController struct {
	ApiBase
}

func (this *YiqingController) GetCityId() {
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

	o := orm.NewOrm()
	yiqing := yi_qing_city.YiqingCity{}

	//接受参数
	//id, err := this.GetInt("country")
	//beego.Info("yiqing:",yiqing)
	//if err != nil {
	//	this.ResultJson(10105, "参数有问题")
	//} else{
	//	yiqing.Id = id
	//	//beego.Info("id:",id)10104
	//}
	country_id, err := this.GetInt("countrie_id")

	if err != nil {
		this.ResultJson(10105, "参数有问题")
	} else {
		beego.Info("country_id:", country_id)
		yiqing.Countrie_id = country_id
	}
	yiqingerr := o.Read(&yiqing)
	beego.Info("yiqing:", yiqing)

	if yiqingerr != nil {
		this.ResultJson(10104, "数据库查询失败")
	} else {
		this.ResultJson(0, "成功", yiqing)
	}
	//this.ServeJSON()
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
	beego.Info("table", "td_yiqing_city")

	//接受参数
	country_id, err := this.GetInt("countrie_id")

	if err != nil {
		this.ResultJson(10105, "参数有问题")
	}

	o := orm.NewOrm()
	//yiqing := yi_qing_city.YiqingCity{}
	yiqing := make([]yi_qing_city.YiqingCity, 0)
	// 也可以直接使用 Model 结构体作为表名
	num, err := o.QueryTable("td_yiqing_city").Filter("countrie_id", country_id).All(&yiqing)
	beego.Error("Returned Rows Num: %s, %s", num, err)
	if err != nil {
		this.ResultJson(10104, "数据库查询失败")
	} else {
		this.ResultJson(0, "成功", yiqing)
	}
	//this.ServeJSON()
}
