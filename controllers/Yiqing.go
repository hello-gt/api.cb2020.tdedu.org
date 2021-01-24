package controllers

import (
	"api.cb2020.tdedu.org/common"
	"api.cb2020.tdedu.org/models/yi_qing_city"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
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

/**
 * 获取国家城市对应数据
 * @param int $country 留学国家 1293 美国
 * @param int $option  数据选项 1 全部列表数据 按照累计人数倒叙  2 只要地图坐标数据
 */
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

/**
 * 获取疫情国家坐标
 * @param int $ct_name 留学国家名称
 * @return boolean|array
 * @author ligt
 * @return json
 */
func (c *YiqingController) GetCoordinate() {
	//验证IP
	checkIPRes := checkIP()
	if checkIPRes == false {
		c.ResultJson(10102, "验证ip不通过")
	}
	//验证参数
	secret := c.GetString("secret")
	if secret == "" {
		c.ResultJson(10101, "secret不能为空")

	} else if secret != beego.AppConfig.String("secret") {
		c.ResultJson(10102, "secret错误")
	}
	//接受参数
	ct_name := c.GetString("ct_name")
	if ct_name == "" {
		c.ResultJson(10009, "参数不能为空")
	}
	currentDir := common.GetCurrentDirectory()
	path := currentDir + "/data/" + ct_name + ".json"
	beego.Info("path", path)
	data, err := ioutil.ReadFile(path)
	beego.Info("data", data)
	if err != nil {
		c.ResultJson(10009, "参数不能有问题")
	}
	c.ResultJson(0, "成功", data)
}
