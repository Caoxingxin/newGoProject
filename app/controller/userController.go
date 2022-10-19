package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newadmin/Tools"
	"newadmin/model"
	"newadmin/services"
)
//json:"name"数据格式为json格式，并且json字段名为name,form:"name"表单参数名为name
type User struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Tel string `json:"tel" form:"tel"`
}
type Test struct {
	Status int `json:"status"`
	Type int `json:"type"`
	Name string `json:"name"`
}
type Users struct {
	Name string `json:"name"`
	Pwd string `json:"password"`
	CreateTime string `json:"createtime"`
}
func UserTest(content *gin.Context)  {
	user:= User{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	if content.ShouldBind(&user) == nil{
		log.Println(user.Name)
		log.Println(user.Email)
		log.Println(user.Tel)
	}

	test:=&Test{
		Status: 1,
		Type: 2,
		Name: "AA",
	}				//设置json格式

	//name := content.DefaultQuery("name","default")
	//name,ok := content.GetQuery("name");
	nameId := content.Param("id");			//获取url上的参数
	ip:= content.ClientIP();		//获取用户IP
	HtmlType := content.DefaultQuery("type","1")
	//if !ok {
	//	name = "caoxin"
	//}
	// 查询我们之前在日志中间件，注入的键值数据
	example := content.MustGet("example").(string)
	log.Println(example)
	if HtmlType=="1" {
		content.HTML(http.StatusOK,"login.html",gin.H{
			"title":"返回界面",
			"name": "Hello,World!",
		})
	}else{
		content.HTML(http.StatusOK,"adminLogin.html",gin.H{
			"title":"返回界面",
			"name": "Hello,World!",
		})
	}

	content.JSON(200,gin.H{
		"name" : user.Name,
		"email" : user.Email,
		"tel":user.Tel,
		"ip":ip,
		"Test":test,
		"id" : nameId,
	})								//正常接口返回
}

func ContrastFile(content *gin.Context)  {
	result:=services.ContrastFile()
	if result!=0 {
		content.JSON(200,result)
	}else{
		content.JSON(101,"error")
	}

}

func SetCookie(content *gin.Context)  {
	content.SetCookie("site_cookie","cookie_value",3600,"/","localhost",false,true)
}
func GetCookie(content *gin.Context)  {
	data,error := content.Cookie("site_cookie")
	if error != nil {
		content.String(200,data)
		return
	}
	content.String(500,"not,found!")
}

func UploadList(content *gin.Context)  {
	content.HTML(http.StatusOK,"upload.html","")
}
func UploadFile(content *gin.Context)  {
	file,_ :=content.FormFile("file")
	log.Println(file.Filename)
	//保存
	content.SaveUploadedFile(file,"./assets/upload/"+file.Filename)
	content.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
}

func SetSession(content *gin.Context)  {
	sessionData := content.DefaultQuery("session","")
	session := sessions.Default(content);
	session.Set("hello",sessionData)
	session.Save()
	content.String(http.StatusOK,"session设置成功")
}
func GetSession(content *gin.Context)  {
	//初始化session
	session := sessions.Default(content);
	sessionData := session.Get("hello")
	content.JSON(http.StatusOK,gin.H{
		"sessionData":sessionData,
	})
}

func UserGorm(content *gin.Context)  {
	//使用连接池，避免重复读取问题
	db := Tools.GetDb()

	//新增
	//u:= model.Users{
	//	UserName: content.PostForm("name"),
	//	PassWord: content.PostForm("pwd"),
	//	CreateTime: time.Now().Unix(),
	//}
	////插入一跳数据
	//if err:= db.Debug().Create(&u).Error; err!=nil {
	//	log.Println("插入失败",err)
	//}

	//查询
	//u := []model.Users{}
	//	result := db.Model(&u).Debug().Where("username = ?",content.PostForm("name")).Find(&u)		//find所有   ，first一条
	//if errors.Is(result.Error,gorm.ErrRecordNotFound) {			//仅适用于first，take，last三种方法
	//	fmt.Println("找不到记录")
	//}
	//content.JSON(http.StatusOK,u)
	//fmt.Println(u)				//打印查询到的数据

	//定义map类型，key为字符串，value为interface{}类型，方便保存任意值
	data := make(map[string]interface{})
	data["mas"] = "ss"
	data["username"] = "abcd"
	//content.JSON(200,data)
	fmt.Println(data)
	//更新
	//u := []model.Users{}
	//db.Model(&u).Where("username = ?" ,content.PostForm("name")).Update("password",content.PostForm("pwd")).Find(&u)
	//content.JSON(http.StatusOK,u)

	//删除
	//u := []model.Users{}
	//db.Where("username = ?" ,content.PostForm("name")).Delete(&u)
	//content.JSON(http.StatusOK,u)
	//log.Println(host,userName,pwd,dbName,port)

	//事务更新
	//tx := db.Begin()
	//u := []model.Users{}
	//err := tx.Model(&u).Where("password = 1239").Update("password",999).Find(&u).Error
	//if err != nil {
	//	tx.Rollback()
	//}else{
	//	tx.Commit()
	//}

	//关联查询
	u := []model.Users{}
	db.Where("id = ?" ,1).First(&u)
	var profile model.Profile
	db.Model(&u).Preload("Profile").Find(&u)
	content.JSON(200,gin.H{
		"user" :u,
		"profile" : profile,
	})
	fmt.Println(u)
}