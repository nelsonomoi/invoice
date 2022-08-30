package domain


type Currency struct {
	ID int64 `json:"id"`
	Iso_code  string `json:"iso_code"`
	Description string `json:"description"`
}


type Postal_code struct{
	ID int64 `json:"id"`
	Code int64 `json:"code"`
	Town string `json:"town"`
}