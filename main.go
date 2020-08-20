package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	logrus "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type Config struct {
	StartPage  int
	EndPage    int
	WalletName string
}

var logger *logrus.Logger

func main() {
	var env string
	flag.StringVar(&env, "conf", "walletexplorer", "-conf filename")
	flag.Parse()

	logger = NewLogger(env)
	configFilePath := "./" + env + ".json"

	configBuf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic("config file error, now process exit")
	}

	// buf, err := ioutil.WriteFile("./BitcoinFogAddresses.html")

	var config Config
	err = json.Unmarshal(configBuf, &config)
	if err != nil {
		panic("parse config file error, now process exit")
	}

	resultFileName := "./" + env + "_result.txt"
	f, err := os.OpenFile(resultFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()
	for i := config.StartPage; i < config.EndPage; i++ {
		logger.Info("query page num:", strconv.Itoa(i))
		url := "https://www.walletexplorer.com/wallet/" + config.WalletName + "/addresses?page=" + strconv.Itoa(i)
		parseAddr(url, f)
	}
}

func parseAddr(url string, f *os.File) {
	// buf, err := ioutil.ReadFile("./BitcoinFogAddresses.html")
	resp, err := http.Get(url)
	if err != nil {
		logger.Errorln(err)
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	s := string(buf)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, addr := range visit(nil, doc) {
		if _, err := f.WriteString(addr); err != nil {
			logger.Error(err)
		}
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && strings.Contains(a.Val, "/address/") {
				links = append(links, a.Val[9:])
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func testSignal() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func fCh1(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Println("will sleep:", num, " second(s).")
	time.Sleep(time.Duration(num) * time.Second)
}

func testGoroutineExit() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go fCh1(&wg, 1)

	wg.Add(1)
	go fCh1(&wg, 2)

	wg.Add(1)
	go fCh1(&wg, 3)

	wg.Wait()
	fmt.Println("normail exit")
	close(ch)
}

func getSelfRealPath() {
	configFilePath := func() string {
		optional := []string{"./go.mod", "./practice_go/go.mod"}
		for _, tmp := range optional {
			_, err := os.Stat(tmp)
			if err == nil {
				return tmp
			} else if os.IsNotExist(err) {
				return tmp
			}
		}
		panic("not found avaiable config file")
	}()

	fmt.Println(configFilePath)
}
