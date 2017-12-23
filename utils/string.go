package utils

import (
	"strconv"
	"dix975.com/basic/logger"
	"regexp"
	"bytes"
	"strings"
)

type StringUtils struct {
	String string
}

func (s StringUtils) ToInt() int{

	if i, err := strconv.Atoi(s.String); err == nil{
		return i
	} else {
		logger.Error.Println(err)
		panic(err)
	}

}


var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string)(string){
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 { chunks[idx] = bytes.Title(val) }
	}
	return string(bytes.Join(chunks, nil))
}

func MakeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}