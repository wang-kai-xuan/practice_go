package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

type person struct {
	Name string
	Age  int
}

func TestMap(t *testing.T) {
	// test content:
	// init,write, read
	m := map[string]string{"name": "王凯旋", "age": "24", "gender": "男"}
	fmt.Println(m)
	m["name"] = "陈刚容"
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)
}
func TestMap1(t *testing.T) {
	// test content:
	// init,write, read
	m := make(map[string]string)
	m["cgr"] = "陈刚容"
	fmt.Println(m)
	m["wkx"] = "王凯旋"
	fmt.Println(m)
	delete(m, "wkx")
	fmt.Println(m)
	name := m["nihao"]
	fmt.Println(name)
}

func TestMap2(t *testing.T) {
	// test content:
	// read not exist val
	m := make(map[string]string)
	m["wkx"] = "王凯旋"
	name := m["wkx"]
	fmt.Println(name)
	name1 := m["wkx1"]
	if name1 == "" {
		fmt.Println("name not exist")
	} else {
		fmt.Println("name exist")
	}
	fmt.Println("end")
}

func TestMap3(t *testing.T) {
	type person struct {
		name   string
		height int
	}

	m := make(map[string]*person)
	m["wkx"] = &person{"wkx", 170}
	delete(m, "wkx")
	delete(m, "wkx")
	delete(m, "")
}
func TestMap4(t *testing.T) {
	m := make([]string, 0, 10)
	fmt.Println(len(m))
	fmt.Println(m)
}

func TestMap5(t *testing.T) {
	walletCodeBalanceMap := make(map[string]int64)
	walletCodeBalanceMap["hello"] = 7
	b1 := walletCodeBalanceMap["hello"]
	if _, ok := walletCodeBalanceMap["hello"]; ok {
		fmt.Println("exist hello")
	}
	if _, ok := walletCodeBalanceMap["hello1"]; ok {
		fmt.Println("exist hello1")
	}
	fmt.Println(b1)
	// fmt.Println(b2)
}

// 比较字符串的前缀
func TestStringPrefix(t *testing.T) {
	name1 := "wangkaixuan"
	name2 := "wangkx"

	fmt.Println(strings.HasPrefix(name1, "wang"))
	fmt.Println(strings.HasPrefix(name1, "wangkx"))
	fmt.Println(strings.HasPrefix(name2, "wang"))
}

// 比较字符串的前缀
func TestStringPrefix(t *testing.T) {
	name1 := "wangkaixuan"
	name2 := "wangkx"

	fmt.Println(strings.HasPrefix(name1, "wang"))
	fmt.Println(strings.HasPrefix(name1, "wangkx"))
	fmt.Println(strings.HasPrefix(name2, "wang"))
}

// byte转字符串
func TestJson(t *testing.T) {
	p := person{Name: "wangkaixuan", Age: 25}
	var r interface{}
	r, e := json.Marshal(p)
	if e != nil {
		fmt.Println(e.Error())

	}

	fmt.Printf("%s\n", r)
	fmt.Printf("%v\n", r)
	r1 := fmt.Sprintf("%s\n", r)
	fmt.Printf("%s\n", r1)
}

func TestJson1(t *testing.T) {
	str := `["wkx","cgr"]`
	res := gjson.Parse(str)

	for item, key := range res.Array() {
		fmt.Println(item, key)
	}
	fmt.Println(str)
}

// 问题：如何获取interface类型变量的成员
// 回答：fmt.Sprintf
func TestInterfaceMember(t *testing.T) {
	var i interface{}
	i = 1
	i = "nihao"

	i1 := fmt.Sprintf("%s", i)
	i = true
	fmt.Println(i)
	fmt.Println(i1)
}

func TestSet(t *testing.T) {
	// 初始化
	s := New()

	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println("list of all items", s.List())

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
func TestNumConvertToString(t *testing.T) {
	age := 25
	stat := "you are #param"
	fmt.Println(strings.Replace(stat, "#param", strconv.Itoa(age), 1))
}

func TestDeleteArray(t *testing.T) {
	// delete the elem when range a variable
	// Will it cause panic?
	//NOTE delete an array when traverse itself, will miss an item
	name := []string{"wkx", "cgr", "wkx0", "wkx1", "wkx2", "wkx3"}
	for index, val := range name {
		if index == 1 {
			copy(name[1:], name[2:])
		}
		fmt.Println(index, val)
	}
}

func Prinf1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Prinf1:%d\n", i)
	}
}

func Prinf2(cnt int) {
	fmt.Printf("Prinf2:%d\n", cnt)
}

func own_print(flag string, quit chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s:%d\n", flag, i)
		time.Sleep(time.Second * 1)
	}
	quit <- flag
}

func TestGo(t *testing.T) {
	go1 := make(chan string)
	go2 := make(chan string)

	go own_print("go1", go1)
	go own_print("go2", go2)

	quitCnt := 0

	for {
		select {
		case str1 := <-go1:
			fmt.Printf("Process %s\n", str1)
			quitCnt++

		case str2 := <-go2:
			fmt.Printf("Process %s\n", str2)
			quitCnt++

		default:
			if quitCnt >= 2 {
				close(go1)
				close(go2)
				fmt.Printf("exit\n")
				return
			}
		}
	}
}

func TestBigInt(t *testing.T) {
	i1 := new(big.Int)
	i1.SetString("10000000000000000000", 10)

	i2 := new(big.Int)
	i2.SetString("10000000000000100000", 10)

	i3 := new(big.Int)
	i3.SetString("10000000000000100000", 10)

	fmt.Println(i1.Cmp(i2))
	fmt.Println(i2.Cmp(i1))
	fmt.Println(i2.Cmp(i3))

	fmt.Println(i3.String())
}
func TestBigFloat(t *testing.T) {
	i1 := new(big.Float)
	i1.SetString("10000000000000000000")

	i2 := new(big.Float)
	i2.SetString("10000000000000100000")

	i3 := new(big.Float)
	i3.SetString("10000000000000100000")

	fmt.Println(i1.Cmp(i2))
	fmt.Println(i2.Cmp(i1))
	fmt.Println(i2.Cmp(i3))

	fmt.Println(i3.String())
	fmt.Println(i3.Float64())
}

func TestIntConvertFloat(t *testing.T) {
	s1 := "10000"
	f1, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(f1)
}

func TestIoutils(t *testing.T) {
	url := "https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Details_of_the_Object_Model"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	// b, err := ioutil.ReadAll(resp.Body)
	io.Copy(&buf, resp.Body)
	resp.Body.Close()
	str := buf.String()
	fmt.Printf("%s", str)
}

func TestAppendFile(t *testing.T) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("text to append\n"); err != nil {
		log.Println(err)
	}
}

func TestFile(t *testing.T) {
	err := ioutil.WriteFile("testwrite.txt", []byte("Dumping bytes to a file\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadFile("./testwrite.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(buf)
	fmt.Println(content)
}

func TestStr(t *testing.T) {
	words := "nnihao"
	fmt.Println(words)
}

type Person struct {
	alias string
	age   int
}

func TestMapPassParams(t *testing.T) {
	names := make(map[string]Person)
	names["wkx0"] = Person{"wkx", 26}
	names["wkx1"] = Person{"wkx", 26}
	names["wkx2"] = Person{"wkx", 26}
	names["wkx3"] = Person{"wkx", 26}
	modifyName(names)
	fmt.Println(names)
}
func modifyName(names map[string]Person) {
	for key := range names {
		item := names[key]
		item.alias = "====="
		names[key] = item
	}
}

func changeName(names []string) {
	names[0] = "wkx1"
}

func TestSliceParams(t *testing.T) {
	names := []string{"wangkaixuan", "cgr", "wkx"}
	changeName(names)
	assert.True(t, names[0] == "wkx1")
}

func Test_bigNum(t *testing.T) {
	num1 := big.NewInt(0)
	num1.SetString("184467440737095520009", 10)
	fmt.Println(num1.String())
	fmt.Println(num1.Int64())

	num2 := big.NewInt(0)
	num2.SetString("10000000000", 10)
	num1.Div(num1, num2)

	fmt.Println(num1.String())
	fmt.Println(num1.Int64())
}
