package dao

import (
	"fmt"

	"github.com/Jay-Chou118/mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Admin{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Notice{},
			&model.Product{},
			&model.ProductImg{},
			&model.Order{},
			&model.Favorite{},
		)
	if err != nil {
		fmt.Println("err", err)
	}

}
