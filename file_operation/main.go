package main

import (
	"file_operation/file_op"
	"file_operation/view_model"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前执行文件的绝对地址
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	println(dir)
	// 定义一个目录并创建这个目录，具有完全权限
	dirp := dir + "\\text_files"
	os.Mkdir(dirp, os.ModePerm)
	// 定义文件路径
	filename := dirp + "\\text.txt"
	s := file_op.Say()
	println(s)
	p := view_model.BasePage{Title: "标题H", Body: "i love girl我爱你 (ノ｀Д)ノ\n"}
	// 写入数据
	f := file_op.OvrrideTextFile(filename, &p)
	println(f)
	// 往这个文件追加数据
	file_op.AppendTextFile(filename, "i love women我爱她 (ノ｀Д)ノ\n")
	file_op.AppendTextFile(filename, "i love book我爱它 (ノ｀Д)ノ\n")
	// 从文件读出数据
	println(file_op.ReadTextFile(filename))
	println("--------------------")
	// 同样从文件读出数据,同时把数据存储到指针变量p中
	file_op.LoadPage(filename, &p)
	println(p.Body)
}
