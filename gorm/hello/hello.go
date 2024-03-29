package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "XXX", Price: 400})

	// Read
	var product Product
	var another Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Println(&product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	db.First(&another, "code = ?", "XXX") // find product with code D42
	fmt.Println(&product)
	fmt.Println("XXX PRODUCT", &another)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Println(&product)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	fmt.Println(&product)

	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println(&product)

	// Delete - delete product
	db.Delete(&product, 1)
}
