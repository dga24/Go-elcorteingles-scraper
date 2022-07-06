package domain

type Product struct {
	Id            string  `json:"id"`
	Title         string  `json:"name"`
	OriginalPrice float32 `json:"o_lPrice"`
	FinalPrice    float32 `json:"f_Price"`
	Discount      int     `json:"discount"`
	Category      string  `json:"Category"`
	Link          string  `json:"Link"`
}

type ProductRepository interface {
	TestDb() (err error)
	InsertProduct(*Product) (err error)
	DeleteFromCategory(s string) (err error)
	ProductIsNew(link string) bool
	// insertListProducts(prs []*product) (err error)
}