package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Fazendo transaction com lock pessimista
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronic"
	tx.Debug().Save(&c)
	tx.Commit()
	//category := Category{Name: "Eletronicos"}
	//db.Create(&category)
	//category2 := Category{Name: "Cozinha"}
	//db.Create(&category2)
	//product := Product{
	//	Name:       "Batedeira",
	//	Price:      1000,
	//	Categories: []Category{category, category2},
	//}
	//db.Create(&product)
	//db.Create(&SerialNumber{
	//	Number:    "123456",
	//	ProductID: product.ID,
	//})
	//product2 := Product{
	//	Name:       "Mouse",
	//	Price:      1000,
	//	CategoryID: category.ID,
	//}
	//db.Create(&product2)
	//var categories []Category
	//err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	//if err != nil {
	//	panic(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.Name, ":")
	//	for _, product := range category.Products {
	//		fmt.Println("- ", product.Name)
	//	}
	//}
	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, p := range products {
	//	fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	//}
}

func createOne(db *gorm.DB) {
	db.Create(&Product{
		Name:  "Notebook 3",
		Price: 3000,
	})
}

func createMany(db *gorm.DB) {
	products := []Product{
		{Name: "Notebook", Price: 1000},
		{Name: "Mouse", Price: 50},
		{Name: "Keyboard", Price: 100},
	}
	db.Create(&products)
}

func getOneById(db *gorm.DB) {
	var product Product
	db.First(&product, 3)
	fmt.Println(product)
}

func getOneByParam(db *gorm.DB) {
	var product Product
	db.First(&product, "name = ?", "Keyboard")
	fmt.Println(product)
}

func getAll(db *gorm.DB) {
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}

func getWhere(db *gorm.DB) {
	var products []Product
	db.Where("price > ?", 110).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}

func update(db *gorm.DB) {
	var product Product
	db.First(&product, 5)
	product.Name = "New Notebook 3"
	db.Save(&product)
	var product2 Product
	db.First(&product2, 5)
	fmt.Println(product2)
}

func delete(db *gorm.DB) {
	var product Product
	db.First(&product, 5)
	fmt.Println(product)
	db.Delete(&product)
}
