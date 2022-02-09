package main

import (
	"food-delivery/component"
	"food-delivery/middleware"
	"food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"food-delivery/modules/upload/uploadtransport/gin_upload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStrAWS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = runService(db)
	if err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	r.POST("/upload", gin_upload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
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
