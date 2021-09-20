package file_op

import (
	"file_operation/view_model"
	"fmt"
	"io/ioutil"
	"os"
)

func Say() string {
	return fmt.Sprintln("hello file")
}

// ReadTextFile 读取指定文件名的所有数据,返回字符串
func ReadTextFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		return string(b), err
	}
	return string("@读取时发生错误,查看返回的error"), err
}

// AppendTextFile 打开一个文件,然后往里追加数据。文件不存在则创建且则追加数据。
func AppendTextFile(filename string, s string) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return false
	}
	// 追加字符串
	n, err := f.WriteString(s)
	_ = n
	f.Close()
	return err == nil
}

// OvrrideTextFile 打开一个文件,然后往里覆盖数据。文件不存在则创建且则覆盖数据。
func OvrrideTextFile(filename string, b *view_model.BasePage) bool {

	f, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if e != nil {
		return false
	}
	// 追加字节流 可使用WriteString追加字符串
	n, e := f.Write([]byte(b.Body))
	_ = n
	f.Close()

	return e == nil
}

func LoadPage(filename string, b *view_model.BasePage) {
	s, err := ReadTextFile(filename)
	if err != nil {
		return
	}
	b.Body = s
}

func init() {
}
