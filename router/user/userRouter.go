package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"newadmin/app/controller"
	"newadmin/middleware"
)

func InitRouter()  {

	fmt.Println("路由加载")
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")		//加载目录下所有html代码
	//使用中间件
	router.Use(middleware.Logging())
	//设置session
	store := cookie.NewStore([]byte("pwd"))		 		// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	router.Use(sessions.Sessions("mysession",store))

	user:= router.Group("/user")
	{
		user.GET("/test/:id",controller.UserTest)
		//user.GET("/login", func(context *gin.Context) {
		//	context.JSON(200,gin.H{
		//		"name":"ss",
		//	})
		//})
		user.GET("/contrastFile",controller.ContrastFile)

		//cookie操作
		user.GET("/setCookie",controller.SetCookie)
		user.GET("/getCookie",controller.GetCookie)

		//上传文件
		user.GET("/userFormView",controller.UploadList)
		user.POST("/upload",controller.UploadFile)

		//session操作
		user.GET("/setSession",controller.SetSession)
		user.GET("/getSession",controller.GetSession)

		//gorm测试
		user.Any("/userGorm",controller.UserGorm)
	}
	router.Run(":8089")
}