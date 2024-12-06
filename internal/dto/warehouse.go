package dto

type WarehouseUri struct {
	Id int `uri:"warehouse_id" binding:"required"`
}

type WarehouseProductInReq struct {
	Supplier  Supplier  `json:"supplier" binding:"required"`
	Warehouse Warehouse `json:"warehouse" binding:"required"`
	Items     []Item    `json:"items" binding:"dive,required"`
	Note      string    `json:"note"`
}

type Warehouse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
