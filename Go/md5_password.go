package main
import (
    "crypto/md5"
    "fmt"
    "encoding/hex"
)

func main(){
   fmt.Println(encrypt("yumiot,,.."))
	fmt.Println(encrypt("yumiot123,,.."))
}

func encrypt(encryptedData string) string {
    //fmt.Println("这是要加密的数据： ",encryptedData)
	md5Ctx := md5.New()
    md5Ctx.Write([]byte(encryptedData))
    //cipherStr := md5Ctx.Sum(nil)
    //fmt.Println("这是加密后的初始数据： ",cipherStr)
    //fmt.Println("这是加密后的十六进制数据： ",hex.EncodeToString(cipherStr))
	//查看数据类型
	// fmt.Println(reflect.TypeOf(hex.EncodeToString(cipherStr)))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}