package main

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
示例代码目录中存在此结构:
	D:\go developer
		1
			sdadfsef_1.mp4
		2
			jtjttyrj_2.mp4
		......
想要用代码将所有mp4文件统一存放在根目录D:\go developer。且用它上级目录名(数字)命名新文件：
	D:\go developer
		1.mp4
		2.mp4
		......
诸如"D:\go developer\1"的文件夹已没有文件，后续被清理。请看main函数实现此示例。
*/

func main() {
	start := time.Now()
	const opRoot = "D:/go developer"
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
	cleanEmptyDir(opRoot)
	println("empty folders was cleaned up successfully!")
	println(time.Now().Sub(start).Seconds())
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
}
