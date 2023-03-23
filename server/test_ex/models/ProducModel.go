package models

type Book struct {
	Id       int
	BookName string
	Price    int
}

func (b *Book) Builder(id int, name string) *BookBuilder {
	return NewBookBuilder(id, name)
}
func (this *Book) GetInfo() string {
	return "book"
}

type Briefs struct {
	Id   int
	Size string
}

func (this *Briefs) GetInfo() string {
	return "内裤"
}

type Test struct {
	Id   int
	Name string
}

func (this *Test) GetInfo() string {
	return "Test"
}
