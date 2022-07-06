package main

import (
	"bufio"
	"elcorteingles/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const urlDomain string = "www.elcorteingles.es"

func newCollyConfig() (*colly.Collector) {
	c := colly.NewCollector(
		colly.AllowedDomains(urlDomain),
		colly.MaxDepth(9),
		colly.UserAgent("ext/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Accept", "*/*")
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println(r.Headers)
		fmt.Println("Something went wrong:", err)
	})

	return c
}

func productsParser(linkCategory string) (productList []*domain.Product){
	var pr *domain.Product
	c := newCollyConfig()
	c.OnHTML("div", func(e *colly.HTMLElement) {
		if e.Attr("id") == "products-list" {
			fmt.Println("products-list")
			e.ForEach("li", func(i int, h *colly.HTMLElement) {
				product := h.ChildAttr("span","data-json")
				pr = jsonToProduct(product)
				link := h.ChildAttr("a","href")
				if link != ""{
					pr.Link = urlDomain+link
				}
				if pr.FinalPrice != 0{
					pr.Category = getCategoryFromLink(linkCategory)
					productList = append(productList, pr)
				}
			})
		}
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		if e.Attr("class") == "event _pagination_link" {
			e.Request.Visit(e.Attr("href"))
		}
	})
	c.Visit(linkCategory)
	fmt.Printf("c: %v\n", c)
	return productList
}


func jsonToProduct(prjson string) (*domain.Product){
	pr := new(domain.Product)
	if prjson == ""{
		return pr
	}
	datamap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(prjson), &datamap); err != nil{
		log.Println("error 78")
	}
	if datamap["code_a"].(string)==""{
		log.Println("error 81")
		return pr
	}
	pr.Id = datamap["code_a"].(string)
	pr.Title = datamap["name"].(string)
	jsonunit8Price,_ := json.Marshal(datamap["price"])
	stringPrice := string(jsonunit8Price)
	priceMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(stringPrice), &priceMap); err != nil{
		log.Println("linea93",err)
		return pr
	}
	aux,_ := strconv.ParseFloat(fmt.Sprint(priceMap["o_price"]),32)
	pr.OriginalPrice = float32(aux)
	aux2,_ := strconv.ParseFloat(fmt.Sprint(priceMap["f_price"]),32)
	pr.FinalPrice = float32(aux2)
	aux3,_ := strconv.Atoi(fmt.Sprint(priceMap["discount_percent"]))
	pr.Discount = aux3
	return pr
}



func getIdFromLink(link string) (id string) {
	link = strings.Split(link, "/")[len(strings.Split(link, "/"))-2]
	return strings.Split(link, "-")[0]
}


func readFile() []string{
	var links []string
	f, err := os.Open("categoriListTXTElec.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		links = append(links, scanner.Text())
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return links
}

func getCategoryFromLink(link string) string{
	if len(strings.Split(link, "/"))>2 {
		r := strings.Split(link, "/")[len(strings.Split(link, "/"))-2]
		return r
	}
	return ""
}