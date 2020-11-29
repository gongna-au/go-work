package main

import (
"fmt"
"net/http"
"strconv"
"io/ioutil"
"github.com/PuerkitoBio/goquery"
)
//爬取网页内容

func HttpGet(url string) (result string ,err err){
     resp,err := http.Get(url)
    if err != nil {
	  panic(err)
	}
	//读取网页内容
	body, err := ioutil.ReadAll(rp.Body)
         if err != nil {
	panic(err)
        }
           content := string(body)
		   defer rp.Body.Close()
		   return
}





func doWork(start int, end int){
	fmt.Printf("正在爬取从%d到%d的页面",start,end)
			//1.明确爬取范围
    for i:=start,i<=end,i++ {
		url ="http://search.zongheng.com/s?keyword=%E7%81%AB%E5%BD%B1&pageNo=2&sort=&isFromHuayu="
		 strconv.Itoa(i-1)
		 fmt.Println("url=",url)
	
	      //2.把所有的网站内容爬取下来
	    result，err :=HttpGet(url)
	    if err !=nil {
		  fmt.Println("HttpGet err =",err)
		  continue

	    }
	      // 3.请求html页面
	         res, err := http.Get("http://metalsucks.net")
	        if err != nil {
	          // 错误处理
	          log.Fatal(err)
	        }
	          defer res.Body.Close()
	            if res.StatusCode != 200 {
	              log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	            }
	                 // 加载 HTML document对象
	                  //goquery的使用
	                   dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
                    if err != nil {
	                  panic(err)
			        }
			           //选择器语法就是css选择器语法
                        dom.Find(".search-tab").Each(func(i int, selection *goquery.Selection){
	                        selection.Find(".tit").Each(func(i int, title *goquery.Selection) {
		                          fmt.Printf("%3d   ", num)
		                         fmt.Println(title.Text())
		                         num++
	                       })
                        })


	                          //把内容写入文件
	                            fileName :=strconv.Itoa(i)+".html"
	                            f,errl :=os.Create(fileName)
	                            if errl !=nil {
									fmt.Println("os.Create =",errl)
									continue
								}
								f.WriteString(result)//x写内容
								f.Close()//关闭文件

    }

}

func main() {
	var start ,end int 
	fmt.Printf("请输入起始页面：")
	fmt.Scan(&start)
	fmt.Printf("请输入一个终止页：")
	fmt.Scan(&end)
	doWork(start,end)
}