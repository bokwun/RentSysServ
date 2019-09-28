package main

import (
	"database/sql"
	"io"
	"os"
	"strconv"
	"time"

	//导入mysql前要加“_”才不会报错
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/rs/cors"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Id       string `json:"id" form:"id"`
	Password string `json:"password" form:"password"`
}
type Message struct {
	IdUser    string `json:"idUser" form:"idUser"`
	Xiaoqumc  string `json:"xiaoqumc" form:"xiaoqumc"`
	Shi       int    `json:"shi" form:"shi"`
	Ting      int    `json:"ting" form:"ting"`
	Wei       int    `json:"wei" form:"wei"`
	Mianji    int    `json:"mianji" form:"mianji"`
	Diceng    int    `json:"diceng" form:"diceng"`
	Gongceng  int    `json:"gongceng" form:"gongceng"`
	Chewei    string `json:"chewei" form:"chewei"`
	Zujin     int    `json:"zujin" form:"zujin"`
	Quyu      string `json:"quyu" form:"quyu"`
	Biaoti    string `json:"biaoti" form:"biaoti"`
	Miaoshu   string `json:"miaoshu" form:"miaoshu"`
	Lianxiren string `json:"lianxiren" form:"lianxiren"`
	Lianxidh  string `json:"lianxidh" form:"lianxidh"`
	DateTime  string `json:"dateTime" form:"dateTime"`
	PicName   string `json:"picName" form:"picName"`
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求方法
		method := c.Request.Method
		//允许访问所有域
		c.Header("Access-Control-Allow-Origin", "*")
		//header的类型
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, "+
			"Token,session,X_Requested_With,Accept, Origin, Host, Connection, "+
			"Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, "+
			"X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		//服务器支持所有跨域请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, UPDATE")
		//跨域关键设置，让浏览器可以解析
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,"+
			"Last-Modified,Pragma,FooBar")
		//跨域请求是否需要带cookie信息，默认设置为true
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			c.JSON(http.StatusOK, "options Request!")
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	router := gin.Default()

	//只有当真正数据库通信的时候才创建连接
	db, err := sql.Open(
		"mysql",
		"root:Root.1231@tcp(127.0.0.1:3306)/user?parseTime=true",
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//正常启动时打开的连接数
	db.SetMaxIdleConns(20)
	//最大能打开的连接数
	db.SetMaxOpenConns(20)

	//当我们需要在open之后就知道连接的有效性的时候，可以通过ping()来进行
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	api := router.Group("/api")
	{
		api.Use(cors())
		api.POST("/postmsg", func(c *gin.Context) {
			id := c.PostForm("id")
			password := c.PostForm("password")
			_, err := db.Exec("insert into person(id, password) values (?, ?)", id, password)
			if err != nil {
				//系统退出原因，追踪代码可知os.exit()
				log.Fatalln(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"id":       id,
				"password": password,
			})
		})
		api.POST("/message", func(c *gin.Context) {
			timestamp := time.Now().Unix()
			ts := time.Unix(timestamp, 0)
			t := ts.Format("1234567890123")

			idUser := c.PostForm("idUser")
			xiaoqumc := c.PostForm("xiaoqumc")
			shi := c.PostForm("shi")
			ting := c.PostForm("ting")
			wei := c.PostForm("wei")
			mianji := c.PostForm("mianji")
			diceng := c.PostForm("diceng")
			gongceng := c.PostForm("gongceng")
			chewei := c.PostForm("chewei")
			zujin := c.PostForm("zujin")
			quyu := c.PostForm("quyu")
			biaoti := c.PostForm("biaoti")
			miaoshu := c.PostForm("miaoshu")
			lianxiren := c.PostForm("lianxiren")
			lianxidh := c.PostForm("lianxidh")
			dateTime := c.PostForm("dateTime")
			picName := t + c.PostForm("picName")

			_, err := db.Exec("insert into message(idUser, xiaoqumc, shi, ting, wei, mianji,"+
				"diceng, gongceng, chewei, zujin, quyu, biaoti, miaoshu, lianxiren, lianxidh, dateTime, picName) values "+
				"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", idUser, xiaoqumc, shi, ting, wei, mianji,
				diceng, gongceng, chewei, zujin, quyu, biaoti, miaoshu, lianxiren, lianxidh, dateTime, picName)
			if err != nil {
				log.Fatalln(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"idUser": idUser, "xiaoqumc": xiaoqumc, "shi": shi, "ting": ting,
				"wei": wei, "mianji": mianji, "diceng": diceng, "gongceng": gongceng,
				"chewei": chewei, "zujin": zujin, "quyu": quyu, "biaoti": biaoti,
				"miaoshu": miaoshu, "lianxiren": lianxiren, "lianxidh": lianxidh, "dateTime": dateTime, "picName": picName,
			})
		})

		api.GET("/checkUser/:id", func(c *gin.Context) {
			id := c.Param("id")
			var person Person
			err := db.QueryRow(
				"select id, password from person where id=?", id).Scan(&person.Id, &person.Password)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"person": nil,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"person": person,
			})
		})

		//api.GET("/returnAll", func(c *gin.Context) {
		//	rows, err := db.Query("select Id, Password from person")
		//	if err != nil{
		//		log.Fatalln(err)
		//	}
		//	defer rows.Close()
		//
		//	persons := make([]Person, 0)
		//	for rows.Next() {
		//		var person Person
		//		//赋值
		//		err := rows.Scan(&person.Id, &person.Password)
		//		if err != nil {
		//			log.Println(err)
		//		}
		//		persons = append(persons, person)
		//	}
		//	if err = rows.Err(); err != nil{
		//		log.Println(err)
		//	}
		//	c.JSON(http.StatusOK, gin.H{
		//		//这个persons跟前端获取后台值须一样
		//		"persons": persons,
		//	})
		//})

		api.GET("/returnMsg", func(c *gin.Context) {
			rows, err := db.Query("select * from message")
			if err != nil {
				log.Fatalln(err)
			}
			defer rows.Close()

			messages := make([]Message, 0)
			for rows.Next() {
				var message Message
				err := rows.Scan(&message.IdUser, &message.Xiaoqumc, &message.Shi, &message.Ting,
					&message.Wei, &message.Mianji, &message.Diceng, &message.Gongceng,
					&message.Chewei, &message.Zujin, &message.Quyu, &message.Biaoti,
					&message.Miaoshu, &message.Lianxiren, &message.Lianxidh, &message.DateTime, &message.PicName)
				if err != nil {
					log.Println(err)
				}
				messages = append(messages, message)
			}
			if err = rows.Err(); err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"messages": messages,
			})
		})

		api.GET("/returnUserMsg/:idUser", func(c *gin.Context) {
			idUser := c.Param("idUser")
			rows, err := db.Query("select * from message where idUser=?", idUser)
			if err != nil {
				log.Fatalln(err)
			}
			defer rows.Close()

			messages := make([]Message, 0)
			for rows.Next() {
				var message Message
				err := rows.Scan(&message.IdUser, &message.Xiaoqumc, &message.Shi, &message.Ting,
					&message.Wei, &message.Mianji, &message.Diceng, &message.Gongceng,
					&message.Chewei, &message.Zujin, &message.Quyu, &message.Biaoti,
					&message.Miaoshu, &message.Lianxiren, &message.Lianxidh, &message.DateTime, &message.PicName)
				if err != nil {
					log.Println(err)
				}
				messages = append(messages, message)
			}
			if err = rows.Err(); err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"messages": messages,
			})
		})

		api.GET("/returnMsgWhere", func(c *gin.Context) {

			quyu := c.DefaultQuery("quyu", "unknown")
			zujin := c.Query("zujin")
			shi := c.Query("shi")
			var rows *sql.Rows
			var first, second, third string
			//将zujin从string->
			zujinInt, err := strconv.Atoi(zujin)
			zujinMin := zujinInt - 500
			zujinMax := zujinInt + 500

			//判断，将三个string拼接到sel这个select语句
			if quyu == "不限" {
				first = "quyu!=?"
			} else {
				first = "quyu=?"
			}

			if zujin == "1000" {
				second = "zujin<=?"
			} else if zujin == "5000" {
				second = "zujin>=?"
			} else if zujin == "0" {
				second = "zujin!=?"
			} else {
				second = "zujin>=? and zujin<=?"
			}

			if shi == "0" || shi == "5" {
				third = "shi>=?"
			} else {
				third = "shi=?"
			}
			//拼接后的select语句
			sel := "select * from message where " + first + " and " + second + " and " + third
			//租金这块有时需要两个值，有时只需要一个
			if zujin == "1000" || zujin == "5000" || zujin == "0" {
				rows, err = db.Query(sel, quyu, zujin, shi)
			} else {
				rows, err = db.Query(sel, quyu, zujinMin, zujinMax, shi)
			}

			if err != nil {
				log.Fatalln(err)
			}
			defer rows.Close()

			messages := make([]Message, 0)
			for rows.Next() {
				var message Message
				err := rows.Scan(&message.IdUser, &message.Xiaoqumc, &message.Shi, &message.Ting,
					&message.Wei, &message.Mianji, &message.Diceng, &message.Gongceng,
					&message.Chewei, &message.Zujin, &message.Quyu, &message.Biaoti,
					&message.Miaoshu, &message.Lianxiren, &message.Lianxidh, &message.DateTime, &message.PicName)
				if err != nil {
					log.Println(err)
				}
				messages = append(messages, message)
			}
			if err = rows.Err(); err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"messages": messages,
			})
		})

		//删除数据
		api.GET("/delMessage", func(c *gin.Context) {
			idUser := c.DefaultQuery("idUser", "unknown")
			dateTime := c.Query("dateTime")
			rows, err := db.Query("delete from message where idUser=? and dateTime=?", idUser, dateTime)
			if err != nil {
				log.Fatalln(err)
			}
			defer rows.Close()
		})

		//上传图片
		api.POST("/upload", func(c *gin.Context) {
			file, handler, err := c.Request.FormFile("upload")
			filename := handler.Filename
			//打印在控制台
			log.Println("Received file:", handler.Filename)

			timestamp := time.Now().Unix()
			ts := time.Unix(timestamp, 0)
			t := ts.Format("1234567890123")
			//图片资源提交到niginx服务器上
			//  /usr/local/webserver/nginx/html/img/  是路径，要是后面没有/结尾，则最后一个/的后面为创建的文件的前缀
			//windows
			out, err := os.Create("E:/Program Files/nginx-1.16.1/html/img/" + t + filename)

			//linux
			//out, err := os.Create("/usr/local/webserver/nginx/html/img/" + t +filename)
			if err != nil {
				c.JSON(500, gin.H{
					"status": 500,
					"msg":    0,
				})
				log.Fatalln(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				c.JSON(500, gin.H{
					"status": 500,
					"msg":    0,
				})
				log.Fatalln(err)
			}
			c.String(http.StatusOK, "uploaded...", timestamp)
			c.String(http.StatusOK, "uploaded...", ts)
			c.String(http.StatusOK, "uploaded...", t)
		})

		//api.GET("/", func(c *gin.Context) {
		//	c.String(http.StatusNotFound, "404 Not Found")
		//})

	}

	_ = router.Run(":8081")
}
