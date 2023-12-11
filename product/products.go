package product

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID int    `json:"category_id"`
	// Categories    Category               `json:"categories"`
	// BrachProducts []branch.BranchProduct `json:"branch_products`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
