package main

import (
	"file_operation/file_op"
	"file_operation/view_model"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	println(dir)
	dirp := dir + "\\text_files"
	os.Mkdir(dirp, os.ModePerm)
	filename := dirp + "\\test.txt"
	s := file_op.Say()
	println(s)
	p := view_model.BasePage{Title: "标题H", Body: "ssssssssss\n我问问 sb"}
	f := file_op.OvrrideTextFile(filename, &p)
	println(f)
	file_op.AppendTextFile(filename, "\nnew sb")
	file_op.AppendTextFile(filename, "\nnew sb")
	file_op.ReadTextFile(filename)
	println(file_op.ReadTextFile(filename))
	println("--------------------")
	file_op.LoadPage(filename, &p)
	println(p.Body)
}
