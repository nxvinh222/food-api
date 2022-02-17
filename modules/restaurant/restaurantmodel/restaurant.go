package restaurantmodel

import (
	"fmt"
	"food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
	LikeCount       int            `json:"like_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name"`
	Addr  *string        `json:"address" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"address" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return fmt.Errorf("restaurant name can't be blank")
	}

	return nil
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}
