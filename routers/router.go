package routers

import (
	"AiJia/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/api/v1.0/areas",&controllers.AreaController{},"get:GetArea")

	beego.Router("/api/v1.0/houses/index",&controllers.HouseIndexController{},"get:GetHouseIndex")

	beego.Router("/api/v1.0/session",&controllers.SessionController{},"get:GetSessionData;delete:DeleteSessionData")

	beego.Router("/api/v1.0/users",&controllers.UserController{},"post:Reg")
	beego.Router("/api/v1.0/user/avatar",&controllers.UserController{},"post:Postavatar")

	beego.Router("/api/v1.0/user",&controllers.UserController{},"get:GetUserData")
	beego.Router("/api/v1.0/user/name",&controllers.UserController{},"put:UpdateName")
    beego.Router("/api/v1.0/user/auth",&controllers.UserController{},"get:GetAuth;post:SetAuth")

	beego.Router("/api/v1.0/sessions",&controllers.SessionController{},"post:Login")



	beego.Router("/api/v1.0/user/houses",&controllers.HouseController{},"get:GetHousesData")
    beego.Router("/api/v1.0/houses",&controllers.HouseController{},"post:AddHouse")
	beego.Router("/api/v1.0/houses/?:id",&controllers.HouseController{},"get:GetDetailHouseData")


}
