package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var find bool
var listFind bool
var buf bytes.Buffer
var num int = 10

const  server  = "https://www.xsbiquge.com"
func getBody(url string)  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get failed,err:%v\n", err)
	}
	rb, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("ioutil readall failed.err:%v\n",err)
	}
	fmt.Printf("%s",rb)
}

func main()  {
	/*
	// 定义解析字符串
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// 解析字符串类型
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	// 定义一个函数变量
	var f func(*html.Node)
	// 为函数变量赋值
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	// 调用函数
	f(doc)
	 */
	//getBody("https://www.xsbiquge.com/91_91600/")

	//Get_Book("https://www.xsbiquge.com/91_91600/")
	buf.Write([]byte("Hello "))
	fmt.Fprintf(&buf, "world!")
	buf.WriteTo(os.Stdout)
}


func Get_Book(target string) error {
	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("get book use http get failed,err:%v\n", err)
		return err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("get book html parse failed,err:%v\n",err)
		return err
	}
	parseList(doc)
	return nil

}

func parseList(n *html.Node)  {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "list" {
				listFind = true
				parseTitle(n)
				break
			}
		}
	}
	if !listFind {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseList(c)
		}
	}
}

func parseTitle(n *html.Node)  {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// 获取文章title
				for c := n.FirstChild; c != nil; c=c.NextSibling {
					buf.WriteString(c.Data + "\n")
				}
				url := a.Val
				target := server + url
				everyChapter(target)
				num--
			}
		}
	}
	if num <= 0 {
		return
	} else {
		for c := n.FirstChild; c != nil; c=c.NextSibling {
			parseTitle(c)
		}
	}
}

func everyChapter(target string)  {
	fmt.Println(target)
	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("http get failed, err:%v\n",err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	find = false
	parse(doc)
	text, err := os.Create("./三国之他们非要打种地的我.txt")
	if err != nil {
		fmt.Printf("create file failed, err:%v\n",err)
	}
	file := strings.NewReader(buf.String())
	file.WriteTo(text)
}

func parse(n *html.Node)  {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "content" {
				find = true
				parseTxt(&buf,n)
				break
			}
		}
	}
	if !find {
		for c := n.FirstChild; c!=nil; c=c.NextSibling {
			parse(c)
		}
	}
}

func parseTxt(buf *bytes.Buffer, n *html.Node)  {
	for c := n.FirstChild; c!= nil; c=c.NextSibling {
		if c.Data != "br" {
			buf.WriteString(c.Data + "\n")
		}
	}
}
// 结论: 结构体没掌握好，buffer 、strings 不会使用，文件创建没掌握好，借口没掌握好