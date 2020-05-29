package main

import (
	"fmt"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request)  {
	/*
	cookies := r.Cookies()
	for index, cookie := range cookies{
		fmt.Printf("index:%d cookie:%#v\n",index, cookie)
	}*/
	cookie, err  := r.Cookie("sessionId")
	fmt.Printf("cookie:%#v, err:%v \n",cookie,err)

/*
	cookie := &http.Cookie{
		Name: "sessionId",
		Value: "adf121239sdfsdf923r3423",
		MaxAge: 3600,
		Domain: "localhost",
		Path: "/",
	}

 */
	//http.SetCookie(w, cookie)

	w.Write([]byte("hello"))
}

func main()  {
	http.HandleFunc("/",indexHandle)
	http.ListenAndServe(":9090",nil)
}

