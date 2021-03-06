package main
import(
    "net/http")
func main(){
req, err := http.NewRequest("GET", swoop.url, nil)
//初始化
for key ,value ：= range swoop.header {
    req.Header.Add(key, value)
}
//发起请求
resp, err := client.Do(req)
if err != nil {
    log.Fatalf("do client err->%v", err)
}
//接收响应
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    log.Fatalf("read resp err->%v", err)
}
//定义结构体
type Swoop struct {
    url    string
    header map[string]string
}
//评价人数
commentCount := `<span>(.*?)评价</span>`
rp2 := regexp.MustCompile(commentCount)
txt2 := rp2.FindAllStringSubmatch(html, -1)

//评分
pattern3 := `property="v:average">(.*?)</span>`
rp3 := regexp.MustCompile(pattern3)
txt3 := rp3.FindAllStringSubmatch(html, -1)

//电影名称
pattern4 := `img width="(.*?)" alt="(.*?)" src=`
rp4 := regexp.MustCompile(pattern4)
txt4 := rp4.FindAllStringSubmatch(html, -1)
url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25)
swoop := &Swoop{url, header}
html := swoop.get_html_header()
    
  }

