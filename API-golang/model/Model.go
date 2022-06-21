package model

import "time"

//model customer

type CustomerModel struct {
	Id        int       `db:"Id" json:"Id"`
	FirstName string    `db:"FirstName" json:"firstName"`
	LastName  string    `db:"LastName" json:"lastName"`
	Address   string    `db:"Address" json:"address"`
	Birthday  time.Time `db:"Birthday" json:"birthday"`
}

// type RoomModel struct {
// 	Id               int
// 	ReserveDate      NullTime
// 	ReserveStartTime NullTime
// 	ReserveEndTime   NullTime
// 	RoomNo           string
// }

type RoomModel struct {
	Id               int
	ReserveDate      time.Time
	ReserveStartTime time.Time
	ReserveEndTime   time.Time
	RoomNo           string
}
