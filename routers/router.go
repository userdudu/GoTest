package routers

import (
	"shane/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user/profile", &controllers.UserController{}, "get:ProFile")


    beego.Router("/greguser", &controllers.MainController{}, "get:PageJoin")
    //beego.Router("/ploguser", &controllers.MainController{}, "post:Login")
    //beego.Router("/preguser", &controllers.MainController{}, "post:RegUser")
}
