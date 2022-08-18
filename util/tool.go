package util

import (
	"crypto/md5"
	"io"
	"log"
	"fmt"
)

func Md5String(str string) string {
	m := md5.New()
	_,err := io.WriteString(m,str)
	if err != nil {
		log.Fatal(err)
	}
	b := m.Sum(nil)
	return fmt.Sprintf("%x",b)

}
