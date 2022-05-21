package main

import (
	"food-delivery/component"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
	"food-delivery/modules/restaurantlike/transport/ginrestaurantlike"
	"food-delivery/modules/upload/uploadtransport/gin_upload"
	"food-delivery/modules/user/usermodel"
	"food-delivery/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStrAWS")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3Secret := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3Secret, s3Domain)

	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	err = migrateDB(db)
	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(db, s3Provider, secretKey)
	err = runService(appCtx)
	if err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	v1 := r.Group("/api/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))
	v1.POST("/upload", gin_upload.Upload(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	return r.Run(":5600")
}

func migrateDB(db *gorm.DB) error {
	//db.Migrator().DropTable(usermodel.User{})
	//db.Migrator().DropTable(recipemodel.Recipe{})
	//db.Migrator().DropTable(elementmodel.Element{})

	err := db.AutoMigrate(usermodel.User{})
	if err != nil {
		return err
	}
	//err = db.AutoMigrate(restaurantmodel.Restaurant{})
	//if err != nil {
	//	return err
	//}
	err = db.AutoMigrate(restaurantlikemodel.Like{})
	if err != nil {
		return err
	}
	return nil
}
