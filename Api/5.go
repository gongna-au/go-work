
//页面登录
function  login() {
	var formData = new FormData();
	formData.append("username",$('input[name="username"]').val())
	formData.append("password",$('input[name="password"]').val())
	$.ajax({
		url:"http://127.0.0.1:8888/user/login",
		type:'POST',
		data:formData,
		processData: false,
		contentType: false,
		dataType:'json',
		success:function(data){
			console.log(data)
			if(data.code==200) {
				location.href = 'index.html'
				localStorage.setItem("myproject_username",$('input[name="username"]').val())
				localStorage.setItem("myproject_password",$('input[name="password"]').val())
			}else{
				alert(data.msg) //提示错误！！
			}
		}
	});
}
//添加controllers 包管理接口
func (m UserC) Login(g *gin.Context) {
	fmt.Println("login.........")
	rsp := new(Rsp)
	name := g.PostForm("username")
	pass := g.PostForm("password")
 
	//var gerr *gvalid.Error
	//gerr = gvalid.Check(g.PostForm("username"), "required", nil)
	//if gerr != nil {
	//	rsp.Msg = "faild"
	//	rsp.Code = 201
	//	rsp.Data = gerr.Maps()
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
	//gerr = gvalid.Check(g.PostForm("password"), "required", nil)
	//if gerr != nil {
	//	rsp.Msg = "faild"
	//	rsp.Code = 201
	//	rsp.Data = gerr.Maps()
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
 
	findfilter := bson.D{{"username", g.PostForm("username")}, {"password", g.PostForm("password")}}
	cur, err := m.Mgo.Collection(db.User).Find(context.Background(), findfilter)
	if err != nil {
		rsp.Msg = "faild"
		rsp.Code = 201
		rsp.Data = err.Error()
		g.JSON(http.StatusOK, rsp)
		return
	}
	for cur.Next(context.Background()) {
		elme := new(models.User)
		err := cur.Decode(elme)
		if err == nil {
			if elme.Username == name && elme.Password == pass {
				var info = new(LoginInfo)
				info.User = elme
				token, err := util.GenerateToken(g.PostForm("username"), g.PostForm("password"))
				if err == nil {
					info.Token = token
				}
				rsp.Msg = "success"
				rsp.Code = 200
				rsp.Data = info
				g.JSON(http.StatusOK, rsp)
				return
			}
		}
	}
 
	rsp.Msg = "user is null"
	rsp.Code = 201
	rsp.Data = err
	g.JSON(http.StatusOK, rsp)
}
//对用户名和密码在数据库中检验是否存在
