package service

import (
	"elcorteingles/domain"
	"elcorteingles/telegram"
)


const DiscountNotif = 55

type DefaultProductService struct {
	repo domain.ProductRepository
	tel telegram.Telegram
}


func (s *DefaultProductService) testDb() (err error) {
	return s.repo.TestDb()
}

func (s *DefaultProductService) InsertProduct(pr *domain.Product) (err error) {
	return s.repo.InsertProduct(pr)
}

func (s *DefaultProductService) InsertListProducts(prs []*domain.Product) (err error) {
	for _, pr := range prs{
		if s.repo.ProductIsNew(pr.Link){
			err = s.InsertProduct(pr)
			if pr.Discount>DiscountNotif{
				go s.tel.NotifyProduct(*pr)
			} 
		}
	}
	return err 
}

func (s DefaultProductService)DeleteFromCategory(category string) (err error) {
	return s.repo.DeleteFromCategory(category)
}

func NewProductService(repository domain.ProductRepository) *DefaultProductService {
	tel := telegram.NewTelegram()
	return &DefaultProductService{repository,tel}
}
