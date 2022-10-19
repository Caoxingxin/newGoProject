### gin框架

go get报错

go: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 142.251.42.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established con
nection failed because connected host has failed to respond.

在cmd里面输入

 go env -w GO111MODULE=on

  go env -w GOPROXY=https://goproxy.io,direct

之后重启





使用Golang 在配置SDK时会出现 "The selected directory is not a valid home for Go SDK" 的错误

在Go\src\runtime\internal\sys\zversion.go文件, 添加一行版本信息。

	const TheVersion = `go1.17.6`

1.17.6为版本号自行修改

主要系统环境变量GOROOT都需要且路径为go安装路径





##路由

* 1.router使用gin包中的


```
	router := gin.Default()
	//常用请求方式
	router.GET('/test',getting)
	router.POST('/test',posting)
	router.PUT('/test',putting)
	router.DELETE('/test',deleting)
	
	//'/test'后面跟的一般为相应视图或者返回方法，如下
	router.GET("/", func(context *gin.Context) {
		context.String(200,"test,gin")
	})
	//GET获取参数
	context.DefaultQuery('name','default')
	context.Query('name')	
	

	router.POST("/testPostParms", func(context *gin.Context) {
		name:=context.PostForm("name")
		nick:= context.DefaultPostForm("nick","AE86")
		//gin.H为返回一个hash
		context.JSON(200,gin.H{
			"name" : name,
			"message": http.StatusOK,
			"nick" : nick,
		})
	})
	//POST获取参数
	context.PostForm('name')
	context.DefaultPostForm('name','default')
```