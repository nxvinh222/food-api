package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	ID   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

type RestaurantUpdate struct {
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	newRestaurant := Restaurant{Name: "KFC", Addr: "nguyen hong"}
	if err := db.Create(&newRestaurant); err != nil {
		fmt.Println(err)
	}

	var restaurants []Restaurant

	db.Where("name = ?", "KFC").Find(&restaurants)
	fmt.Println(restaurants)

	if err := db.Table(Restaurant{}.TableName()).Where("id = 2").Delete(nil); err != nil {
		fmt.Println(err)
	}

	db.Table(Restaurant{}.TableName()).Where("id = 4").Updates(map[string]interface{}{
		"name": "",
		"addr": "nct",
	})

	newName := ""
	db.Table(Restaurant{}.TableName()).Where("id = 5").Updates(&RestaurantUpdate{Addr: &newName})
}
