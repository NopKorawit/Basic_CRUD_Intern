package repository

import (
	"customer/handler"
	"customer/model"
	"fmt"
)

type RoomRepository struct{}

func (r *RoomRepository) ReadCustomer(Id string) ([]model.RoomModel, error) {
	db, err := handler.Init_DB()
	if err != nil {
		return nil, err
	}
	defer handler.Close_DB(db)
	returnRoom := make([]model.RoomModel, 0)
	execStr := fmt.Sprintf("exec SP_RoomBookingManagement @Process = N'%s', @Id = '%s' ", "FIND", Id)
	rows, err := db.Query(execStr)
	if err != nil {
		return nil, err
	}
	var Book model.RoomModel
	for rows.Next() {
		if err := rows.Scan(&Book.Id, &Book.ReserveDate, &Book.ReserveStartTime, &Book.ReserveEndTime, &Book.RoomNo); err != nil {
			return returnRoom, err
		}
		returnRoom = append(returnRoom, Book)
	}
	return returnRoom, nil
}
