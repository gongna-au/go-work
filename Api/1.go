//发起http请求，在resp中得到访问的请求结果
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

//ioutil.ReadAll 函数从 response 中读取到全部内容;将其结果保存在变量 b 中。resp.Body.Close 关闭 resp 的 Body 流,防止资源泄露,Printf 函数会将结果 b 写出到标准输出流中。
