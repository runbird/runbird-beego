package utils

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"strconv"
)

func SwitchTimeStampToData(timestamp int64) string {
	return strconv.Itoa(int(timestamp))
}

func SwitchMarkDownToHtml(content string) template.HTML {

}

func MD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
