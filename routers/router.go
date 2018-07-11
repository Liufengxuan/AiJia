package routers

import (
	"AiJia/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/v1.0/areas",&controllers.AreaController{},"get:GetArea")
}
