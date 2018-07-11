package controllers

import (
	"github.com/astaxie/beego"
	"AiJia/models"
	"github.com/astaxie/beego/orm"
	//"encoding/json"
)

type AreaController struct {
	beego.Controller
}
func (c *AreaController)RetData(resp map[string]interface{}){

	c.Data["json"]=resp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("Get Area is OK")

	resp:=make(map[string]interface{})
	//从 session  取数据



	//在mysql拿到area的数据
	var area []models.Area
	o:=orm.NewOrm()
	num,err:=o.QueryTable("area").All(&area)
	if err!=nil{
		beego.Error("for Read(&area) ERROR=",err)
		//返回错误json包。
		resp["errno"]=400
		resp["errmsg"]="查询失败"
		c.RetData(resp)

		return
	}
	beego.Info("查询到数据=",num)
	resp["errno"]=0
	resp["errmsg"]="OK"
	resp["data"]=&area
	c.RetData(resp)



}
