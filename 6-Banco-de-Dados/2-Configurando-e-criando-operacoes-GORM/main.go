package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:primaryKey`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// CREATE
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// CREATE BATCH
	//products := []Product{
	//	{Name: "Notebook", Price: 1000.00},
	//	{Name: "Mouse", Price: 50.00},
	//	{Name: "Keyboard", Price: 100.00},
	//}
	//db.Create(products)

	// SELECT ONE
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// SELECT ALL
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// SELECT LIMIT E (OFFSET paginação)
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// SELECT WHERE
	// var products []Product
	// db.Where("price > ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// SELECT LIKE
	var products []Product
	db.Where("name LIKE ?", "%mouse%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

}
