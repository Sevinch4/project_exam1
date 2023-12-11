package branch

type Branch struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type BranchProduct struct {
	BranchID  int `json:"branch_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity`
	//Transactions []transaction.Transaction `json:"transactions`
}
