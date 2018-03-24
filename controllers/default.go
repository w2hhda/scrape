package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
	"math/rand"
	"strings"
)

type News struct {
	Title   string
	Content string
	Date    string
	Source  string
	Heat    int
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = scrape()
	c.ServeJSON()
}

func scrape() string {
	doc, err := goquery.NewDocument("http://www.jinse.com/lives")
	if err != nil {
		log.Fatal(err)
	}

	var newsArr []News
	doc.Find(".lost > li").Each(func(i int, li *goquery.Selection) {
		liveTime := trim(li.Find(".live-date").Text())
		liveInfo := trim(li.Find(".live-info").Text())

		fmt.Println(liveInfo)
		fmt.Println("===================>>>> s")
		liveTitle := trim(SubString(liveInfo, 0, 30))
		fmt.Println(liveTitle)
		fmt.Println("===================>>>> e")
		item := News{
			liveTitle,
			liveInfo,
			liveTime,
			"ICO 监管",
			rand.Intn(10000),
		}
		newsArr = append(newsArr, item)
	})

	result, err := json.Marshal(newsArr)
	if err != nil {
		fmt.Println("fail")
		return ""
	}

	return string(result)
}

func trim(str string) string {

	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}


func SubString(str string, begin, length int) string {
	fmt.Println("Substring =", str)
	rs := []rune(str)
	lth := len(rs)
	fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	return string(rs[begin:end])
}