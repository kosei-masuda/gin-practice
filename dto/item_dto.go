package dto

type CreateItem struct {
	Name        string `json:"name" binding:"required,gte=3"`
	Price       uint   `json:"price" binding:"required,gt=0,lte=999999"`
	Description string `json:"description"`
}

type UpdateItem struct {
	Name        *string `json:"name" binding:"omitnil,gte=3"`
	Price       *uint   `json:"price" binding:"omitnil,gt=0,lte=999999"`
	Description *string `json:"description"`
	Soldout     *bool   `json:"soldout"`
}
