package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type ApiBase struct {
	beego.Controller
}

func (this *ApiBase) init()  {
	checkIP()
}
/**
 * 校验IP(暂未开放)
 */
func checkIP() (bool)  {
	return true
}

/**
 * 秘钥校验
 */

func checkSecret(this *ApiBase) int {
	secret := this.GetString("secret")

	if secret == "" {
		return 20000
	}
	return 20111
}

func (this *ApiBase) ResultJson(code int, msg string, data ...interface{}) interface{} {
	result := make(map[string]interface{})
	result["code"] = code
	result["msg"]  = msg
	result["time"] = time.Now().Unix()

	if len(data) > 0 && data[0] != nil {
		result["data"] = data[0]
	}

	return this.Ctx.Output.JSON(result, true, false)
}