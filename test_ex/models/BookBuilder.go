package models

type BookBuilder struct {
	id       int
	bookName string
	price    int
}

func (b *BookBuilder) SetPrice(price int) *BookBuilder {
	b.price = price
	return b
}
func NewBookBuilder(id int, bookName string) *BookBuilder {
	return &BookBuilder{id: id, bookName: bookName}
}
func (b *BookBuilder) Build() *Book {
	book := &Book{Id: b.id, BookName: b.bookName}
	if b.price > 0 {
		book.Price = b.price
	}
	return book
}
