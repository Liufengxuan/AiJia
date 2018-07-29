package controllers

import (
	"github.com/astaxie/beego"
	"AiJia/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
)



type HouseController struct {
	beego.Controller
}
func (this *HouseController)RetData(resp map[string]interface{}){

	this.Data["json"]=resp
	//转换为json并返回
	this.ServeJSON()
}


func (this *HouseController) GetHousesData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)

	user:=this.GetSession("user").(models.User)
	houses:=[]models.House{}
	o:=orm.NewOrm()
    o.QueryTable("house").Filter("user__id",user.Id).All(&houses)
	if len(houses)==0{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="没有查询到房屋数据"
		beego.Info("没有查询到房屋数据")
	}
	respData:=make(map[string]interface{})
	respData["houses"]=houses
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="ok le"
	resp["data"]=respData
	beego.Info("GetHousesData() rest =",respData)

}

func (this *HouseController) AddHouse(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)

	reqData:=make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody,&reqData)


	house:=models.House{}
	house.Title=reqData["title"].(string)
	price,_:=strconv.Atoi(reqData["price"].(string))
	house.Price=price
	house.Address=reqData["address"].(string)
	room_count,_:=strconv.Atoi(reqData["room_count"].(string))
	house.Room_count=room_count
	acreage,_:=strconv.Atoi(reqData["acreage"].(string))
	house.Acreage=acreage
	house.Unit=reqData["unit"].(string)
	capacity,_:=strconv.Atoi(reqData["capacity"].(string))
	house.Capacity=capacity
	house.Beds=reqData["beds"].(string)
	deposit,_:=strconv.Atoi(reqData["deposit"].(string))
	house.Deposit=deposit
	min_days,_:=strconv.Atoi(reqData["min_days"].(string))
	house.Min_days=min_days
	max_days,_:=strconv.Atoi(reqData["max_days"].(string))
	house.Max_days=max_days


	facilities:=[]models.Facility{}
	o:=orm.NewOrm()
	beego.Info("facility is :",reqData["facility"])

	for _,fid:=range reqData["facility"].([]interface{}){
		f_id,_:=strconv.Atoi(fid.(string))
		fac:=models.Facility{Id:f_id}
		facilities=append(facilities,fac)
		beego.Info("([]interface{})",fid)
	}
	area_id,_:=strconv.Atoi( reqData["area_id"].(string))
	area:=models.Area{Id:area_id}
	house.Area=&area
	user:=this.GetSession("user").(models.User)
	house.User=&user

	house_id,err2:=o.Insert(&house)
	 if err2!=nil{
	 	beego.Info("o.Insert(&house)",err2,&house)
	 	resp["errno"]=models.RECODE_DBERR
	 	resp["errmsg"]="数据访问出错"
	 	return
	 }
	house.Id=int(house_id)

	m2m:=o.QueryM2M(&house,"Facilities")
	num,err4:=m2m.Add(facilities)
	if err4!=nil ||num==0{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="数据访问出错"
		beego.Info("m2m.Add(&facilities)",err4,&facilities)
		return
	}
	respData:=make(map[string]interface{})
	respData["house_id"]=house_id
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	resp["data"]=respData

}


func (this *HouseController) GetDetailHouseData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	respData:=make(map[string]interface{})
	//1.获取当前用户的user—id
	//user_id:=this.GetSession("user").(models.User).Id


	//2.从url中获取房屋id
	house_id,err :=strconv.Atoi( this.Ctx.Input.Param(":id"))
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="获取house——id失败1"
		return
	}


	//3从缓存中获取房屋数据  redis





	//4关联查询
	o:=orm.NewOrm()
	house:=models.House{Id:house_id}
	if err2:=o.Read(&house);err2!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]="获取house——id失败2"
		return
	}
	//respData["acreage"]=house.Acreage
	//respData["address"]=house.Address
	//respData["beds"]=house.Beds
	//respData["capacity"]=house.Capacity
	//respData["deposit"] =house.Deposit
	//respData["facilities"]=house.Facilities
	//respData["img_urls"]=house.Images
	//respData["min_days"]=house.Min_days
	//respData["max_days"]=house.Max_days
	//respData["price"]=house.Price
	//respData[""]
	o.LoadRelated(&house,"Area")
	o.LoadRelated(&house,"User")
	//o.LoadRelated(&house,"Image")
	o.LoadRelated(&house,"Facilities","fid")



	respData["house"]=house
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="ok"
	resp["data"]=respData



	//5存入缓存。
}