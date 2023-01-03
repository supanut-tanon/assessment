package expense

type Expense struct {
	ID     	int    		`json:"id"`
	Title   string 		`json:"title"`
	Amount  float32 	`json:"amount"`
	Note 	string 		`json:"note"`
	Tags 	[]string	`json:"tags"`
}
