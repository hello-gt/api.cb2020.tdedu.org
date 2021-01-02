package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
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
func checkIPs() (bool)  {
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
	fmt.Println(secret)
	return 20111
}
