package domain

import (
	"database/sql"
	"time"
)

type ProductRepositoryDb struct {
	client *sql.DB
}

func NewProductRepositoryDb() ProductRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/elcorteingles")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return ProductRepositoryDb{client}
}

func (r *ProductRepositoryDb) TestDb() (err error) {
	return nil
}

func (r *ProductRepositoryDb) ProductIsNew(link string) bool{
	var count int
	query := "SELECT COUNT(title) FROM products WHERE link = ?";
	rows := r.client.QueryRow(query,link)
	err := rows.Scan(&count)
	if err != nil {
		panic(err)
	}
	if count < 1 {
		return true
	}
	return false
}

func (r *ProductRepositoryDb) InsertProduct(pr *Product) (err error){
	//fmt.Println(pr)
	//if not exist
	productQuery := "INSERT INTO products (id, title, originalPrice, finalPrice, discount, category, link) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = r.client.Exec(productQuery, pr.Id, pr.Title, pr.OriginalPrice, pr.FinalPrice, pr.Discount, pr.Category, pr.Link)
	if err != nil {
		panic(err)
	}

	return err
}

func (r *ProductRepositoryDb) DeleteFromCategory(s string) (err error) {
	productQuery := "DELETE FROM products WHERE category = ?"
	_, err = r.client.Exec(productQuery, s)
	if err != nil {
		panic(err)
	}

	return err
}
