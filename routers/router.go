package routers

import (
	"api.cb2020.tdedu.org/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/yiqing/get_city_list", &controllers.YiqingController{},"post:GetCityList")
}
