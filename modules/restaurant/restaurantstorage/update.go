package restaurantstorage

import (
	"context"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := *s.db

	err := db.Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}
