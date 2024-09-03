package main
import(
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main(){
	testStr := "hello world"
	Md5Inst  := md5.New()
	Md5Inst.Write([]byte(testStr))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("md5:%x\n",Result)
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(testStr))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("sha1:%x\n\n",Result)
}