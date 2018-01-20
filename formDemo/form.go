package formDemo

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体
	// 如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["param"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Golang!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./formDemo/login.html")
		log.Println(t.Execute(w, ""))
	} else {
		r.ParseForm()

		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func init() {
	http.HandleFunc("/", sayHello)   //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由

	fmt.Println("Http server listening at http://localhost:9090")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
