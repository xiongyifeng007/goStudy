package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/rInfo", rInfo)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

func handler(w http.ResponseWriter, r *http.Request) { //http.ResponseWriter is a interface ,can not use "w * http.ResponseWriter
	mu.Lock()

	// http.SetCookie(w, &http.Cookie{Name: " X-Xrsftoken", Value: "ad54fffs32addd422 ", HttpOnly: true})  no use
	// r.AddCookie(&http.Cookie{Name: " X-Xrsftoken", Value: "ad54fffs32addd422 ", HttpOnly: true})
	// w.Header().Add("X-Xrsftoken", "aaaaaaaaaaaa") // half use
	ck := &http.Cookie{
		Name:   "myCookie",  //名字
		Value:  "hello",     //值
		Path:   "/",         //路径
		Domain: "localhost", //域名
		MaxAge: 120,         //存活时间
	}

	http.SetCookie(w, ck) //设置cookie

	// ch2, err := r.Cookie("myCookie")
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// 	return
	// }
	// io.WriteString(w, ch2.Value)
	count++

	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	// token := r.Header.Get("myCookie")
	token, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Count %d\ntoken%s", count, token.Value)
	mu.Unlock()

}

func rInfo(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q]=%q", k, v)
	// }
	// fmt.Fprintf(w, "Host=%q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)
	// if err := r.ParseForm(); err != nil {
	// 	log.Print(err)
	// }
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "Form[%q]=%q", k, v)
	// // }
	http.Redirect(w, r, "/count", 200)

}
