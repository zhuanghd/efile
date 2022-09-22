package efile

import (
	"io/ioutil"
	"os"
	"strings"
)

// ReadFileAsString 以字符串方式读取文件内容
func ReadFileAsString(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	str := ""
	if err == nil {
		str = string(content)
	}
	return str, err
}

// WriteStringToFile 将字符串写入文件中，如果文件已存在则覆盖其内容，如果不存在则先创建文件再写入
func WriteStringToFile(path string, content string) error {
	i := strings.LastIndex(path, string(os.PathSeparator))
	if i > 0 {
		dir := path[:i]
		if Exist(dir) == false {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// Exist 判断文件或者文件夹是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
