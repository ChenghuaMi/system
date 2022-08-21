package util

import (
	"path/filepath"
	"testing"
	"fmt"
)

func TestMd5String(t *testing.T) {
	fmt.Println(Md5String("123456"))
}
func TestCreateFold(t *testing.T) {
	p,_ := filepath.Abs("")
	CreateFold(p)
}
func TestRand(t *testing.T) {
	for i := 0;i<5;i++ {
		fmt.Println(Md5String("123456"))
	}
}
func TestReNameFile(t *testing.T) {
	ReNameFile("E:\\go\\src\\system\\uploads\\20220821\\abc.jpg")
}
func TestReplacePath(t *testing.T) {
	ss := ReplacePath("E:\\go\\src\\system","E:\\go\\src\\system\\uploads\\20220821\\abc.jpg")
	fmt.Println("ss:",ss)
}