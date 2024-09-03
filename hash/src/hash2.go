package main

import (
	"fmt"
	"io"
	"os"
	"crypto/md5"
	"crypto/sha1"
)


// 计算文件sha1
func main(){
	ftest := "test.txt"
	f,err := os.Open(ftest)
	if  err == nil{
		md5h := md5.New()
		io.Copy(md5h,f)
		fmt.Printf("%x %s\n",md5h.Sum([]byte("")),ftest)
		sha1h := sha1.New()
		io.Copy(sha1h,f)
		fmt.Printf("%x %s\n",sha1h.Sum([]byte("")),ftest)
	}else{
		fmt.Println(err)
		os.Exit(1)
	}
}