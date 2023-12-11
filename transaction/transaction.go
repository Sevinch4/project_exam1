package transaction

type Transaction struct {
	ID        int    `json:"id"`
	BranchID  int    `json:"branch_id"`
	ProductID int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
