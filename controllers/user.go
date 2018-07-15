package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"

	"github.com/astaxie/beego/orm"
	"AiJia/models"
)

type UserController struct {
	beego.Controller
}
func (this *UserController)RetData(resp map[string]interface{}){

	this.Data["json"]=resp
	//转换为json并返回
	this.ServeJSON()
}

func (this *UserController) Reg() {
	//获取前端传过来的json数据
	resp:=make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)
	defer this.RetData(resp)


	o:=orm.NewOrm()
	user:=models.User{}
	user.Password_hash=resp["password"].(string)
	user.Mobile=resp["mobile"].(string)
	user.Name=resp["mobile"].(string)


	id,err:=o.Insert(&user)
	if err!=nil{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAEXIST)
		return
	}



	beego.Info("reg success id:=",id)
	resp["errno"]=0
	resp["errmsg"]="注册成功"
	this.SetSession("name",user)

	beego.Info("mobile =",resp["mobile"])
	beego.Info("password =",resp["password"])
	beego.Info("sms_code =",resp["sms_code"])


}
