package util

import (
	"crypto/md5"
	"io"
	"log"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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
func CreateFold(basePath string) string {
	foldName := time.Now().Format("20060102")
	foldPath := filepath.Join(basePath,"uploads",foldName)
	fmt.Println(foldPath)
	if _,err := os.Stat(foldPath);os.IsNotExist(err) {
		fmt.Println("::::::::::::::::::::::")
		os.MkdirAll(foldPath,0777)
		os.Chmod(foldPath,0777)
	}
	return foldPath
}
func ReNameFile(filename string) string {
	ss := filepath.Dir(filename)
	base := filepath.Base(filename)
	ext := filepath.Ext(filename)
	xxx := filepath.Join(ss,Md5String(base))
	xxx = strings.Join([]string{xxx,ext},"")
	return xxx
}
func ReplacePath(basePath string,filePath string) string {
	ss := strings.Replace(filePath,basePath,"",-1)
	ss = strings.Replace(ss,"\\","/",-1)
	return ss
}