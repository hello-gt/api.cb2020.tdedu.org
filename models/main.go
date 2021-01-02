package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type YiQingCity struct {
	Id int
	Countrie_id int
	Coor_id string
	C_name string
	Name string
	Date string
	Confirm_add int
	Confirm int
	Heal int
	Dead int
	Suspect int
	Add_time int
	Is_delete int
}

func init()  {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb   := beego.AppConfig.String("mysqldb")

	// set default database
	orm.RegisterDataBase("default", "mysql",mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurls + ":3306)/" + mysqldb + "?charset=utf8")
	// register model
	orm.RegisterModel(new(YiQingCity))
	// create table
	orm.RunSyncdb("default", false, true)
}