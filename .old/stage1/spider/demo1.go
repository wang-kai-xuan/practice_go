package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var readNum = make(chan int)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		res += string(buf[:n])
	}
	return
}

func getAndSave(url string, i int) {
	// 爬取数据
	fmt.Println("正在爬取第", i, "页,url:", url)
	res, err := HttpGet(url)

	fileName := strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	f.WriteString(res)
	f.Close()
	readNum <- i
}
func working(start, end int) {

	for i := start; i < end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=1" + strconv.Itoa((i-1)*50)
		go getAndSave(url, i)
	}
	for i := start; i < end; i++ {
		fmt.Println("第", <-readNum, "页读取完成")
	}
}
func main() {
	//var start, end int = 1, 3
	//fmt.Println("输入开始抓取的起始页：")
	//fmt.Scan(&start)
	//fmt.Println("输入结束抓取的终止页：")
	//fmt.Scan(&end)
	working(1, 4)

}
