package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	fileInfo, _ := ioutil.ReadDir(".")
	dirName := "newDir"
	os.Mkdir(dirName, os.ModePerm)
	for i := 1; i <= len(fileInfo); i++ {
		tmp := fileInfo[i].Name()
		tmp1 := dirName + "/" + strconv.Itoa(i) + tmp
		exec := exec.Command("cp", tmp, tmp1)
		exec.Output()

	}
}
