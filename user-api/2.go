package main

import (
	"database/sql"  //mysql包
	"encoding/json" // 编码为json类型的字符串传输
	"fmt"           //格式化输入输出

	// "io"            //输入输出流
	_ "github.com/go-sql-driver/mysql" //mysql驱动"
	"log"                              //日志包
	"net/http"                         //http包
)

//Result 从后端返回给前端的反馈字符内容(方便修改反馈的内容)
type Result struct {
	Msg string `json:"msg"`
}

//Post 定义传输的数据类型及数据变量名称
type Post struct {
	UserName string //用户名
	Password string //密码
	Nickname string // 用于替换的新名字
	Grade    int
	Age      int // 年龄
}

//1.先注册

func Register(w http.ResponseWriter, r *http.Request) {
	//先在register函数的开始写上允许跨域和序列化json的请求头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	//处理前端传入的数据及数据库的操作
	// 用r.ParseForm解析
	r.ParseForm()

	//用r.form 解析解析传下来的username
	username := r.FormValue("Username")
	r.ParseForm("Nickname")
	r.ParseForm("Password")
	r.ParseForm("Grade")
	r.ParseForm("Age ")

	//打开典型的打开 mysql 驱动的数据库语法 ，
	//驱动的名称就是: mysql，数据源名称就是：root:hellokang@tcp(localhost:3306)/test。

	//登录数据库
	db, err := sql.Open("mysql", "好像是什么源") //登录数据库
	//连接数据库失败
	if err != nil {
		fmt.Println("Fail to connect to mysql!")
	} //在返回前关闭资源
	defer db.Close() //否则连接数据库成功

	fmt.Println("Succeed to connect to mysql!")
	var post Post
	post.UserName = username[0] //每次取第一个值，所以加[0]
	post.Password = Password[0] //每次取第一个值，所以加[0]
	post.Nickname = Nickname[0]
	post.Grade = Grade[0]
	post.Age = Age[0]

	_, err = db.Exec("INSERT INTO user(Username,Password,Nickname,Grade,Age)VALUES(?,?)", post.Username, post.Password, post.Nickname, post.Grade, post.Age)
	//插入数据

	//直接调用encoding/json包中的json.Marshal()函数即可实现用json编码及解码字符串来实现前后端的“沟通”过程
	//不论是否注册成功，都会对返回的文字进行编码（给用户反馈成功或失败等字符）
	if err != nil { //插入数据失败
		res := Result{Msg: "Fail to register"}
		jsonChar, jsonErr := json.Marshal(res) //编码返回信息
		if jsonErr != nil {                    //编码失败
			fmt.Println("Fail to encode!")
		} else {
			w.Write(jsonChar)
		}
	} else { //插入数据成功
		res := Result{Msg: "Succeed to register"}
		jsonChar, jsonErr := json.Marshal(res) //json格式化
		if jsonErr != nil {
			fmt.Println("Fail to encode!") //编码失败
		} else {
			w.Write(jsonChar)
		}
	}
}

//2.登录

func Login(w http.ResponseWriter, r *http.Request) {
	//我们不是写入数据了，我们需要对数据库的数据进行遍历搜索，
	//一一与传入数据进行对比，存在数据且用户名和密码对应则说明登录成功，反之则登录失败

	//register函数的开始写上允许跨域和序列化json的请求头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	//处理前端传入的数据及数据库的操作
	r.ParseForm()
	Username := r.Form("Username")
	Nickname := r.Form("Nickname")
	Password := r.Form("Password")
	Grade := r.Form("Grade")
	Age := r.Form("Age ")

	//打开服务器
	db, err := sql.Open("mysql", "什么什么源")
	//判断是否连接成功服务器
	if err != nil {
		fmt.Println("Fail to connect to mysql!")
		return
	}
	defer db.Close() //返回时关闭资源
	fmt.Println("Succeed to connect to mysql!")

	var post Post
	post.UserName = Username[0] //每次取第一个值，所以加[0]
	post.Password = Password[0] //每次取第一个值，所以加[0]
	post.Nickname = Nickname[0]
	post.Grade = Grade[0]
	post.Age = Age[0]

	var row *sql.Row
	row = db.QueryRow("select *from user where Username=? and Password=? and Nickname=? and Grade=？ and post.Age=？ ", post.Username, post.Password, post.Nickname, post.Grade, post.Age)
	//检索数据
	err = row.Scan(&post.Username, &post.Password, &post.Nickname, &post.Grade, &post.Age)
	//遍历

	//把字符串反馈给前端

	if err != nil {
		res := Result{Msg: "Fail to login"}    //数据库没有这个数据（登录失败）
		jsonChar, jsonErr := json.Marshal(res) //用json.Marshal把结果json 后输出
		if jsonErr != nil {
			fmt.Println("Fail to encode!") //格式化失败
		} else {
			w.Write(jsonChar) //写出json 后的结果
		}
	} else {
		res := Result{Msg: "Succeed to login"}
		jsonChar, jsonErr := json.Marshal(res) //json格式化
		if jsonErr != nil {
			fmt.Println("Fail to encode!") //格式化失败
		} else {
			w.Write(jsonChar) //写出json 后的结果
		}
	}
}

//注册路由
func main() {
	mux := http.NewServeMux()                //定义新路由
	mux.HandleFunc("/login", Login)          // 登录路由
	mux.HandleFunc("/register", Register)    // 注册路由
	err := http.ListenAndServe(":5007", mux) //监听端口
	//可以登录http://localhost:5007/register?username=zwx&password=123 进行数据测试！
	//在命令行输入:curl "http://localhost:5007/register?username=zwx&password=123"
	if err != nil {
		log.Fatal("ListenAndServer:", err.Error())
	}
}
