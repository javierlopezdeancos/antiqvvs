package price

// Structure for price type
type Structure struct {
	ID         string `json:"id"`
	Nickname   string `json:"nickname"`
	Currency   string `json:"currency"`
	Product    string `json:"product"`
	UnitAmount int64  `json:"unitAmount"`
}
