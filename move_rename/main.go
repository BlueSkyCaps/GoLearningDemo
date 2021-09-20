package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

func timeSince(start time.Time) {
	fmt.Printf("operation is completed, it takes %v seconds.\n",
		time.Since(start).Seconds())
}

func main() {
	const opRoot = "D:/go files"
	defer cleanEmptyDir(opRoot)
	defer timeSince(time.Now())
	dirInfos, err := ioutil.ReadDir(opRoot)
	if err != nil {
		panic(err.Error())
	}

	sort.Slice(dirInfos, func(i, j int) bool {
		i, _ = strconv.Atoi(dirInfos[i].Name())
		j, _ = strconv.Atoi(dirInfos[j].Name())
		return i < j
	})
	for i := 0; i < len(dirInfos); i++ {
		if dirInfos[i].IsDir() {
			abFileRoot := path.Join(opRoot, dirInfos[i].Name())
			abInfo, err := ioutil.ReadDir(abFileRoot)
			if err != nil {
				panic(err.Error())
			}
			fileName := path.Join(abFileRoot, abInfo[0].Name())
			newFileName := strings.TrimRight(abFileRoot, "/") + ".mp4"
			println(fileName, "->", newFileName)
			err = os.Rename(fileName, newFileName)
			if err != nil {
				panic("rename(move) error and stop current:" + err.Error())
			}
		}
	}
	println("rename(move) successfully!")
}

func cleanEmptyDir(root string) {
	if strings.Trim(root, "") == "" {
		return
	}
	dirInfos, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(dirInfos); i++ {
		if dirInfos[i].IsDir() {
			abFileRoot := path.Join(root, dirInfos[i].Name())
			abInfo, err := ioutil.ReadDir(abFileRoot)
			if err != nil {
				panic(err.Error())
			}
			if len(abInfo) == 0 {
				_ = os.Remove(abFileRoot)
			}
		}
	}
	println("empty folders was cleaned up successfully!")
}
