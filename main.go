package main

import (
	"food-delivery/component"
	"food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main()  {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = runService(db)
	if err != nil {
		log.Fatalln(err)
	}

}

func runService (db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
	}

	return r.Run()
}

type Restaurant struct {
	ID   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

