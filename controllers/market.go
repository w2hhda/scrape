package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
)

type CoinItem struct {
	Name string
	Icon string

	Price       string //价格
	Circulation string //流通市值
	Quantity    string //流通数量
	Deal        string
	Raise       string
}

type MarketController struct {
	beego.Controller
}

func (c *MarketController) Get() {
	c.Data["json"] = scrapeMarket()
	c.ServeJSON()
}

func scrapeMarket() string {
	doc, err := goquery.NewDocument("https://www.feixiaohao.com")
	if err != nil {
		log.Fatal(err)
	}

	var itemArr []CoinItem

	doc.Find(".maintable > tbody > tr").Slice(0, 30).Each(func(i int, tr *goquery.Selection) {
		item := CoinItem{}
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			fmt.Println(trim(td.Text()))
			if j == 1 {
				item.Name = trim(td.Text())
			} else if j == 2 {
				item.Circulation = trim(td.First().Text())
			} else if j == 3 {
				item.Price = trim(td.Text())
			} else if j == 4 {
				item.Quantity = trim(td.Text())
			} else if j == 5 {
				item.Deal = trim(td.Text())
			} else if j == 6 {
				item.Raise = trim(td.Text())
			}
		})

		itemArr = append(itemArr, item)


	})

	result, err := json.Marshal(itemArr)
	if err != nil {
		fmt.Println("fail")
		return ""
	}

	return string(result)
}
