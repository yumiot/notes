package main

import (
	"html/template"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

type P1 struct {
	Name string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm()
	fmt.Println("path:", r.URL.Path)

	p := &Product{}
	p.Name = "get_act"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	fmt.Fprintf(w, string(data))
}

func addUser(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../3.html")
	t.Execute(w, nil)
}

func canshu(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	fmt.Println("-------------这是Value获取的--------------")
	fmt.Println("id是: ",r.PostFormValue("id"))
	fmt.Println("用户名: ",r.PostFormValue("name"))
	fmt.Println("密码是: ",r.PostFormValue("pass"))

	fmt.Println("-------------这是Form获取的--------------")
	data := r.PostForm
	fmt.Println("id是: ",data.Get("id"))
	fmt.Println("用户名: ",data.Get("name"))
	fmt.Println("密码是: ",data.Get("pass"))

	data1, _ := json.Marshal(data)
	fmt.Fprintf(w,string(data1))
}

func a(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("aaaaaa")

	fmt.Fprintf(w,"aaaaaaa")
}

func b(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("bbbbbb")

	fmt.Fprintf(w,"bbbbb")
}

func main() {
	fmt.Println("start!!")
	//http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/1", addUser)
	http.HandleFunc("/cs", canshu)
	http.HandleFunc("/a/a",a)
	http.HandleFunc("/a/b",b)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
