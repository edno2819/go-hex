package core_entities

type Stock struct {
	ID          string `json:"id"`           // Unique identifier for the stock entry
	ClothesID   string `json:"clothes_id"`   // Reference to the Clothes ID
	Quantity    int    `json:"quantity"`     // Quantity of the item in stock
	WarehouseID string `json:"warehouse_id"` // Reference to the warehouse storing the stock
	UpdatedAt   string `json:"updated_at"`   // Timestamp for when the stock was last updated
}
