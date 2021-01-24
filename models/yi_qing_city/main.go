package yi_qing_city

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type YiqingCity struct {
	Id          int    `id`
	Countrie_id int    `国家`
	Coor_id     string `坐标id`
	C_name      string `中文名称`
	Name        string `英文名称`
	Date        string
	Confirm_add int
	Confirm     int
	Heal        int
	Dead        int
	Suspect     int
	Add_time    int
	Is_delete   int
}

func (y *YiqingCity) TableName() string {
	return "td_yiqing_city"
}

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")

	// set default database
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":3306)/"+mysqldb+"?charset=utf8")
	// register model
	orm.RegisterModel(new(YiqingCity))
	// create table
	orm.RunSyncdb("default", false, true)
}
