package main

import (
	"fmt"
	"github.com/lbemi/lbemi/test_ex/Factory"
	"github.com/lbemi/lbemi/test_ex/models"
)

func main() {
	user := Factory.CreateUser(Factory.AdminUser)(123, "klisi").(*models.Admin)
	fmt.Println(user)
	fmt.Println(new(Factory.TechFactory).
		CreateProduct(Factory.ProductBook).GetInfo())

	fmt.Println(new(Factory.DailyFactory).
		CreateProduct(Factory.ProductDailyBriefs).GetInfo())

	fmt.Println(new(models.Book).
		Builder(101, "hahahh").
		SetPrice(20).Build())
}
