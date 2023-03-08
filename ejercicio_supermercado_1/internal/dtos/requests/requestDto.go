package requests

type ProductName struct{
	Name string `json:"name" binding:"required"`
}