package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

var fileName = ""
var root = os.Getenv("HOME") + "/juejin"
var post = flag.String("post", "6859538537830858759", "文章编号")
var rootDir = flag.String("root", root, "文件保存的根目录")

func main() {
	fmt.Println(os.Getenv("JUEJINPATH"))
	flag.Parse()
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(".article-title", func(e *colly.HTMLElement) {
		fileName = e.Text
	})

	c.OnHTML(".markdown-body", func(e *colly.HTMLElement) {
		markdown := convertHTMLToMarkdown(e.DOM)
		writeFile(markdown)
	})

	c.Visit("https://juejin.im/post/" + *post)
	c.Wait()
}

// 将Html转为Markdown
func convertHTMLToMarkdown(selection *goquery.Selection) string {
	converter := md.NewConverter("", true, nil)
	markdown := converter.Convert(selection)
	return markdown
}

// 写入文件
func writeFile(content string) {
	filePath := *rootDir + "/" + fileName + ".md"
	var file *os.File

	if !checkFileIsExist(root) {
		err := os.Mkdir(root, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	if checkFileIsExist(filePath) {
		// 如果文件存在，则删除
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 创建文件并写入内容
	fmt.Println("《" + fileName + "》" + " is download on " + *rootDir)
	file, _ = os.Create(filePath)
	n, _ := io.WriteString(file, "## "+fileName+"\n\n"+content)
	// 关闭文件
	file.Close()
	if n == 0 {
		return
	}
}

// 检查文件是否存在
func checkFileIsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}
