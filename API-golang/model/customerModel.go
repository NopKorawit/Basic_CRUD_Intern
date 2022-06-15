package model

//model customer

type CustomerModel struct {
	Id        int    `db:"Id" json:"Id"`
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	Address   string `db:"Address" json:"address"`
	Birthday  int    `db:"Birthday" json:"birthday"`
}
