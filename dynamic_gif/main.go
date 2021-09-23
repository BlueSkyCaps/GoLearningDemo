package main

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"runtime/debug"
)

const separator = string(os.PathSeparator)

var (
	imagesInputRoot string
	gifOutRoot, _   = os.UserHomeDir()
	imgNmeFiles     []string

	// MatchImageFormat 匹配以常见图像文件格式为后缀的正则表达模式
	MatchImageFormat = fmt.Sprintf("%v.*(%v.png|%v.gif|%v.jpg|%v.jpeg)$", separator, separator, separator, separator, separator)
)

func init() {
	if runtime.GOOS == "windows" {
		gifOutRoot = path.Join(gifOutRoot, "desktop")
	}
}
func main() {
	fmt.Println("请输入你想要用于制作动图Gif的图片所在的文件夹路径。(直接粘贴路径并回车即可)")
	_, e := fmt.Scanf("%s", &imagesInputRoot)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(imagesInputRoot)
	fi, e := os.Stat(imagesInputRoot)
	if e != nil || !fi.IsDir() {
		println("不存在输入的这个文件夹哦(⊙o⊙)？检查：你输入的路径是文件夹吗，文件夹存在吗？")
		return
	}
	dirInfo, _ := os.ReadDir(imagesInputRoot)

	for i := 0; i < len(dirInfo); i++ {
		if !dirInfo[i].IsDir() {
			if MatchRegexString(MatchImageFormat, dirInfo[i].Name()) {
				imgNmeFiles = append(imgNmeFiles, dirInfo[i].Name())
			}
		}
	}

	if len(imgNmeFiles) <= 0 {
		println("文件夹中没有一张图片？检查：文件夹有至少一张图片格式的文件吗？制作动态gif,你应该不止一张源图片。")
		return
	}

	inputGifBoss := &gif.GIF{}
	for _, currentGifName := range imgNmeFiles {
		g, _ := os.Open(path.Join(imagesInputRoot, currentGifName))
		currentGifImage, err := gif.Decode(g)
		if err != nil {
			println(err.Error())
			debug.PrintStack()
			return
		}
		_ = g.Close()
		fmt.Printf("%v\n", currentGifImage.Bounds())
		inputGifBoss.Image = append(inputGifBoss.Image, currentGifImage.(*image.Paletted))
		inputGifBoss.Delay = append(inputGifBoss.Delay, 100)
	}
	inputGifBoss.Config = image.Config{ColorModel: inputGifBoss.Config.ColorModel, Width: 200, Height: 200}
	inputGifBoss.LoopCount = -1
	finalGif, _ := os.OpenFile(path.Join(gifOutRoot, UuidGenerator()+"out.gif"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	err := gif.EncodeAll(finalGif, inputGifBoss)
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return
	}
	err = finalGif.Close()
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return
	}
}

func MatchRegexString(p, v string) bool {
	var e error
	var m bool
	if m, e = regexp.MatchString(p, v); m && e == nil {
		return true
	}
	if e != nil {
		println("error in MatchRegexString", e.Error())
	}
	return false
}

func UuidGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
