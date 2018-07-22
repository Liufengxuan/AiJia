package controllers

import (
	"github.com/astaxie/beego"
	"AiJia/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}
func (this *SessionController)RetData(resp map[string]interface{}){

	this.Data["json"]=resp
	beego.Info(resp)
	this.ServeJSON()
}

func (this *SessionController) GetSessionData() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)



	//user:=models.User{}
	//user.Name="wyj"

  	user:=this.GetSession("user")
  	beego.Info("name",user)
  	if user !=nil{
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		resp["data"]=user.(models.User)
		return
		//this.DelSession("name")
	}
	resp["errno"]=models.RECODE_DBERR
	resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)


}
func (this *SessionController) DeleteSessionData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("user")
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
}

func (this *SessionController)Login(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//得到用户信息


	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)
	//beego.Info("name=",resp["mobile"])

	//判断是否合法
	if resp["mobile"]==nil||resp["password"]==nil{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]="请填写用户名密码"
		return
	}
	//与数据库比较验证
	o:=orm.NewOrm()
	user:=models.User{Name:resp["mobile"].(string)}
	if err:=o.Read(&user,"Name");err!=nil{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]="账户不存在"
		return
	}
	if user.Password_hash!=resp["password"].(string){
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]="密码错误"
		return
	}


	//添加Session
	this.SetSession("user",user)
 	this.SetSession("user_id",user.Id)//id留存   取后不删。




	//返回json数据。
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

}
