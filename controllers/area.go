package controllers

import (
	"github.com/astaxie/beego"
	"AiJia/models"
	"github.com/astaxie/beego/orm"
	//"encoding/json"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	"time"
	"encoding/json"
)

type AreaController struct {
	beego.Controller
}
func (c *AreaController)RetData(resp map[string]interface{}){

	c.Data["json"]=resp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {
	var area []models.Area

	resp:=make(map[string]interface{})
	defer c.RetData(resp)
	//从 redis缓存中取数据
	//key:用来区分用的哪个redis ，conn：端口，dbnum：redis里面的0-25，
	cache_conn,cache_connErr:=cache.NewCache("redis",`{"key":"redis1","conn":":6379","dbNum":"0","charset":"GBK"}`)
	if cache_connErr!=nil{
		beego.Error("cache.NewCache err=",cache_connErr)
		return
	}

	if areaData := cache_conn.Get("area");areaData!=nil{
		//fmt.Printf("areadata=%s\n",string(areaData.([]byte)))
		json.Unmarshal(areaData.([]byte),&area)
		//beego.Info("Unmarshal is =",area[0])
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]="Get area data for redis"
		resp["data"]=area
		beego.Info("Get area data for redis")
		return
	}

/*	cacheErr:=cache_conn.Put("aa","value aa",time.Second*3600)
	if cacheErr!=nil{
		beego.Error("cache put error=",cacheErr)
		return
	}
	beego.Info("redis get=",cache_conn.Get("redis1:aa"))
	fmt.Printf("printf redis get=%s\n",cache_conn.Get("aa"))
*/






	//在mysql拿到area的数据

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
	resp["data"]=area
	beego.Info("sqlquery is =",area[0])
	json_str,err := json.Marshal(area)
	if err != nil{
		beego.Info("encoding err")
		return
	}


	cache_conn.Put("area",json_str,time.Second*3600)




}
