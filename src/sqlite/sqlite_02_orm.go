package main


import (
	"fmt"
	"gorm.io/gorm"
"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Phone struct {
	gorm.Model
	Num string
	PersonId uint
}
type Person struct {
	gorm.Model
	Name string
	Phones []Phone `gorm:"foreignKey:PersonId"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Phone{})
	db.AutoMigrate(&Person{})


	db.Create(&Person{Name:"ht", Phones:[]Phone{ {Num:"130"}, Phone{Num:"131"}}})
	var p Person
	db.First(&p, "name = ?", "ht")
	fmt.Println(p)

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

}
