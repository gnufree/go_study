package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
	Sex  string
	Age  int
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Printf("parsefile failed, err:%v\n", err)
			fmt.Fprintf(w, "login login.html failed")
			return
		}
		/*
			user := UserInfo{
				Name: "Mary",
				Sex:  "男",
				Age:  18,
			}
		*/
		m := make(map[string]interface{})
		m["uesrname"] = "Mary"
		m["sex"] = "男"
		m["age"] = 18

		t.Execute(w, m)
		t.Execute(os.Stdout, m)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("username: %s\n", username)
		fmt.Printf("password: %s\n", password)

		if username == "admin" && password == "admin123" {
			fmt.Fprintf(w, "username: %s login success\n", r.FormValue("username"))
		} else {
			fmt.Fprintf(w, "username: %s login failed\n", r.FormValue("username"))

		}

	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
