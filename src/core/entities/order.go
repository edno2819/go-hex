package core_entities

type Order struct {
	ID            string      `json:"id"`             // Unique identifier for the order
	CustomerName  string      `json:"customer_name"`  // Name of the customer
	CustomerEmail string      `json:"customer_email"` // Email of the customer
	Items         []OrderItem `json:"items"`          // List of items in the order
	TotalAmount   float64     `json:"total_amount"`   // Total amount for the order
	OrderStatus   string      `json:"order_status"`   // Status of the order (e.g., "Pending", "Shipped", "Delivered")
	PaymentStatus string      `json:"payment_status"` // Payment status (e.g., "Paid", "Unpaid")
	CreatedAt     string      `json:"created_at"`     // Timestamp for when the order was created
	UpdatedAt     string      `json:"updated_at"`     // Timestamp for when the order was last updated
}

// OrderItem represents a single item within an order.
type OrderItem struct {
	ClothesID string  `json:"clothes_id"` // Reference to the Clothes ID
	Quantity  int     `json:"quantity"`   // Quantity of the item ordered
	Price     float64 `json:"price"`      // Price per unit of the item
}
