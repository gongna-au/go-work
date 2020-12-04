package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery" // 解析html
	"log"
	"strconv"
	//"io/ioutil"
	"github.com/satori/go.uuid" // 生成图片文件名
	"io/ioutil"
	"net/http"
	"os"
)

func getAllUrls() []string {

	var urls []string
	var url string
	for i := 0; i < 2; i++ {
		url = "http://www.meizitu.com/a/more_" + strconv.Itoa(i+1) + ".html" //网址信息
		urls = append(urls, url)
	}
	return urls
}

func parseHtml(url string) {
	doc, err := goquery.NewDocument(url) //获取将要爬取的html文档信息
	if err != nil {
		log.Fatal(err)
	}
	p := make(chan string)                                              //新开管道
	doc.Find(".pic > a > img").Each(func(i int, s *goquery.Selection) { //遍历整个文档

		img_url, _ := s.Attr("src")
		// 启动协程下载图片

		go download(img_url, p) //将管道传入download函数
		fmt.Println("src = " + <-p + "图片爬取完毕")
	})

}

// 下载图片
func download(img_url string, p chan string) {
	uid, _ := uuid.NewV4() //随机生成四段文件名
	file_name := uid.String() + ".jpg"
	fmt.Println(file_name)
	f, err := os.Create(file_name)
	if err != nil {
		log.Panic("文件创建失败")
	}
	defer f.Close() //结束关闭文件

	resp, err := http.Get(img_url)
	if err != nil {
		fmt.Println("http.get err", err)
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("读取数据失败")
	}
	defer resp.Body.Close() //结束关闭
	f.Write(body)
	p <- file_name //将文件名传入管道内

}

func main() {
	urls := getAllUrls()
	for _, url := range urls {
		parseHtml(url)
	}
}
