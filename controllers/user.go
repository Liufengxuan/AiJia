package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"AiJia/models"
	"path"
	"github.com/weilaihui/fdfs_client"
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
	this.SetSession("user",user)
	this.SetSession("user_id",user.Id)
	beego.Info("mobile =",resp["mobile"])
	beego.Info("password =",resp["password"])
	beego.Info("sms_code =",resp["sms_code"])


}



func (this *UserController) Postavatar(){
	 resp:=make(map[string]interface{})
	 defer this.RetData(resp)
	fData,fNamehead,fErr:=this.GetFile("avatar")
	if fErr!=nil{
		resp["errno"]=models.RECODE_SERVERERR
		resp["errmsg"]="图片上传失败"
	}
	//2 得到文件后缀
	suffix:=path.Ext(fNamehead.Filename)//截取文件后缀名称//a.jpg.avi
	fdfsClient,fdfsClientError:=fdfs_client.NewFdfsClient("conf/client.conf")
	if fdfsClientError !=nil{
		beego.Error("fdfs_client.NewFdfsClient  err=",fdfsClientError)
		resp["errno"]=models.RECODE_SERVERERR
		resp["errmsg"]="初始化FastDfs失败"
		return
	}
	fileBuffer:=make([]byte,fNamehead.Size)
	_,err1:=fData.Read(fileBuffer)
	if err1!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="图片没能上传成功1"
		return
	}
	uploadResponse,err2:=fdfsClient.UploadByBuffer(fileBuffer,suffix[1:])
	if err2!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="图片没能上传成功2"
		beego.Info("fdfsClient.UploadByBuffer err=",err2)
		return
	}
	//

	userId:=this.GetSession("user_id")
	o:=orm.NewOrm()

	user:=models.User{Id:userId.(int)}
	err3:=o.Read(&user)
	if err3!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="图片没能上传成功3"
		return
	}
	user.Avatar_url=uploadResponse.RemoteFileId
	if _,err4:=o.Update(&user);err4!=nil{
		beego.Info("头像设置失败",err4)
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="头像设置失败"
		return
	}


	urlMap:=make(map[string]string)
	urlMap["avatar_url"]="http://127.0.0.1:8080/"+uploadResponse.RemoteFileId

	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="图片上传成功"
	resp["data"]=urlMap




}



func (this *UserController) GetUserData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)


	user_id:=this.GetSession("user_id")

	user:=models.User{Id:user_id.(int)}
	o:=orm.NewOrm()
	err3:=o.Read(&user)
	if err3 !=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		beego.Error("GetUserData() o.read is error",err3)
	}
	resp["data"]=user
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
}



func (this *UserController) UpdateName(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)


	user_id:=this.GetSession("user_id")
	beego.Info("id is ",user_id)
	UserName:=make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody,&UserName)
	user:=models.User{Id:user_id.(int)}
	user.Name=UserName["name"]
	beego.Info("update info is :",user)

	o:=orm.NewOrm()
	if _,err:= o.Update(&user,"Name");err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="修改用户名失败"
		return
	}
	if err2:=o.Read(&user);err2!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="修改用户名失败"
		return
	}
	this.SetSession("name",user.Name)
	this.SetSession("user",user)
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="修改成功"
	resp["data"]=UserName
}



func (this *UserController)GetAuth(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)


	user:=this.GetSession("user").(models.User)


	o:=orm.NewOrm()
	if err1:=o.Read(&user,"Id");err1!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="未进行实名认证"
		return
	}

	resp["data"]=user
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	beego.Info("GetAuth()",resp)

}




func(this *UserController)SetAuth(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)

	user:=this.GetSession("user").(models.User)
	rqst:=make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody,&rqst)
	user.Real_name=rqst["real_name"]
	user.Id_card=rqst["id_card"]
	//beego.Info(user)
	o:=orm.NewOrm()
	if _,err:=o.Update(&user,"real_name","id_card");err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="实名认证失败"
		return
	}



	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="实名认证成功"
	resp["data"]=user
	this.SetSession("user",user)


}