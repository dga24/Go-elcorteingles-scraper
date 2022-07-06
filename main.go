package main

import (
	"elcorteingles/domain"
	"elcorteingles/service"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)



func main() {

	//go bot()

	log.Println("Starting aplication...")

	repo := domain.NewProductRepositoryDb()
	service := service.NewProductService(&repo)


	categories := readFile()


	c := make(chan string)

	for _, linkCategory := range categories {
	 parseCategory(linkCategory,service, c)
	}


	for l := range c {
		func (linkCategory string)  {
			time.Sleep(time.Minute*10)
			parseCategory(linkCategory,service,c)
		}(l)
	}
}

func parseCategory(linkCategory string, s *service.DefaultProductService, c chan string) []*domain.Product{
	fmt.Println("updateCategory: ",linkCategory)
	productList := productsParser(linkCategory)
	s.InsertListProducts(productList)
	//c<-linkCategory
	return productList
}







