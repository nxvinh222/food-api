package restaurantstorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) SoftDeleteRestaurant(ctx context.Context, id int) error {
	db := s.db

	err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
