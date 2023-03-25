package Factory

import "github.com/lbemi/lbemi/test_ex/models"

type ProductType int

const (
	ProductDailyBriefs = iota
	ProductBook
	ProductTest
)

type IProductFactory interface {
	CreateProduct(t int) IProduct
}
type IProduct interface {
	GetInfo() string
}

type TechFactory struct{}

func (*TechFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductBook:
		return &models.Book{}
	}
	return nil
}

type DailyFactory struct{}

func (d *DailyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductDailyBriefs:
		return &models.Briefs{}
	}
	return nil
}

type TestFactory struct{}

func (s *TestFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductTest:
		return &models.Test{}
	}
	return nil
}
