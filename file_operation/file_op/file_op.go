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

func ReadTextFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		return string(b), err
	}
	return string("@读取时发生错误,查看返回的error"), err
}

func AppendTextFile(filename string, s string) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return false
	}
	n, err := f.WriteString(s)
	_ = n
	f.Close()
	return err == nil
}

func OvrrideTextFile(filename string, b *view_model.BasePage) bool {

	f, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if e != nil {
		return false
	}
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
