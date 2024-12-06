package dto

type Stock struct {
	Warehouse Warehouse `json:"warehouse"`
	Product   Product   `json:"product"`
	DusStock  int       `json:"dus_stock"`
	PcsStock  int       `json:"pcs_stock"`
}
