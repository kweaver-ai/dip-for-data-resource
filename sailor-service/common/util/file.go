package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileContentBuffer(path string) (*bytes.Buffer, error) {
	bs, err := FileContent(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bs), nil
}

func FileContent(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file: %s", path)
	}
	defer func() {
		_ = file.Close()
	}()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file: %s", path)
	}
	return fileData, nil
}

func FileBase(path string) string {
	ps := strings.Split(path, string(os.PathSeparator))
	ss := strings.Split(ps[len(ps)-1], ".")
	return ss[0]
}

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

//ReadDirFiles 读取目录下的文件，不递归
func ReadDirFiles(path string) ([]string, error) {
	path = strings.TrimSuffix(path, string(os.PathSeparator))
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	//判断是不是目录
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("%s is not a dir", path)
	}
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	fs := make([]string, 0)
	for _, info := range infos {
		fs = append(fs, fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), info.Name()))
	}
	return fs, nil
}

//FindValidJsonPart 寻找字符串content中，key是start开始的，合法的json字符串的值
func FindValidJsonPart(content, start string) (string, int, int) {
	contentLength := len(content)
	startIndex := strings.Index(content, start)
	subStr := content[startIndex:]
	firstIndex := strings.Index(subStr, "{")
	stepContent := subStr[firstIndex:]

	//步进初始化条件
	stepStr := stepContent
	begin := startIndex + firstIndex
	progressLen := 0

	for {
		stepEndIndex := strings.Index(stepStr, "}")
		progressLen += stepEndIndex + 1
		str := content[begin : begin+progressLen]
		log.Printf("string: %s", str)
		if json.Valid([]byte(str)) {
			return str, begin, begin + progressLen
		}
		stepStr = stepStr[stepEndIndex+1:]
		if begin+progressLen >= contentLength {
			break
		}
	}
	return "", -1, -1
}
