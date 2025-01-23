package cart

import (
	"fmt"

	"github.com/AshFire1/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("cart item %d quantity should be greater than zero", i)
		}
		productIDs[i] = item.ProductID
	}
	return productIDs, nil
}
func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}
	//check in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	//total price
	totalPrice := calculateTotalPrice(items, productMap)
	//reduce quantity in db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		if err := h.productStore.UpdateProduct(product); err != nil {
			return 0, 0, err
		}
	}
	//create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "123 Main St",
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	total := 0.0
	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}
func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found in our inventory", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d is out of stock", item.ProductID)
		}
	}
	return nil
}
