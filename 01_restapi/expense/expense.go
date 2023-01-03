package expense

type Expense struct {
	ID     	int    		`json:"id"`
	Title   string 		`json:"title"`
	Amount  string 		`json:"amount"`
	Note 	string 		`json:"note"`
	Tag 	[]string	`json:"tag"`
}
