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
	defer c.RetData(resp)


	//在mysql拿到area的数据
	var area []models.Area
	o:=orm.NewOrm()
	num,err:=o.QueryTable("area").All(&area)
	if err!=nil{
		beego.Error("for Read(&area) ERROR=",err)
		//返回错误json包。
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		c.RetData(resp)

		return
	}
	beego.Info("查询到数据=",num)
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	resp["data"]=&area




}
