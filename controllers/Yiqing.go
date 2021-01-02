package controllers

import (
	"api.cb2020.tdedu.org/models"
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type YiqingController struct {
	beego.Controller
}

type JSONStruct struct {
	Code int
	Result string
	time int64
	Msg  string
}


func (this *YiqingController) GetCityList() {
	now := time.Now()
	fmt.Println(now.Unix())
	mystruct := &JSONStruct{200,"success",now.Unix(), "data"}

	//验证IP
	checkIPRes := checkIP()
	if checkIPRes == false {
		mystruct.Code = 10102
		mystruct.Result = "error"
		mystruct.Msg = "验证ip不通过"
		this.Data["json"] = mystruct
		this.ServeJSON()
	}

	//验证参数
	secret := this.GetString("secret")
	if secret == "" {
		mystruct.Code = 10101
		mystruct.Msg = "secret不能为空"
		this.Data["json"] = mystruct
		this.ServeJSON()
	} else if secret != beego.AppConfig.String("secret") {
		mystruct.Code = 10102
		mystruct.Msg = "secret错误"
		this.Data["json"] = mystruct
		this.ServeJSON()
	}

	//接受参数
	id, err := this.GetInt("id")

	if err != nil {
		mystruct.Code = 10105
		mystruct.Msg = "参数有问题"
	}

	o := orm.NewOrm()
	yiqing := models.YiQingCity{}
	yiqing.Id = id

	yiqingerr := o.Read(&yiqing)
	if yiqingerr != nil {
		mystruct.Code = 10104
		mystruct.Msg = "数据库查询失败"
	}

	beego.Info(yiqing)

	mystruct.Result = "success"
	mystruct.Msg = "成功"
	this.Data["json"] = mystruct
	this.ServeJSON()

}

func checkIP() bool {
	return true
}