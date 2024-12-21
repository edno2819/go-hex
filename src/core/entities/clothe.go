package core_entities

type Clothes struct {
	ID          string  `json:"id"`          // Unique identifier for the clothing item
	Name        string  `json:"name"`        // Name of the clothing item
	Description string  `json:"description"` // Detailed description of the item
	Category    string  `json:"category"`    // Category (e.g., "Shirts", "Pants")
	Size        string  `json:"size"`        // Size (e.g., "S", "M", "L", "XL")
	Color       string  `json:"color"`       // Color of the clothing item
	Material    string  `json:"material"`    // Material (e.g., "Cotton", "Polyester")
	Price       float64 `json:"price"`       // Price of the item
	SKU         string  `json:"sku"`         // Stock Keeping Unit identifier
	CreatedAt   string  `json:"created_at"`  // Timestamp for when the item was added
	UpdatedAt   string  `json:"updated_at"`  // Timestamp for when the item was last updated
}
