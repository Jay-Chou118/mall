package serializer

import "github.com/Jay-Chou118/mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []model.Category) (carousels []Category) {
	for _, item := range items {
		carousel := BuildCategory(&item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
