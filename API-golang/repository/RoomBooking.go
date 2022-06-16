package repository

import (
	"customer/handler"
	"customer/model"
	"fmt"
)

type RoomRepository struct{}

func (r *RoomRepository) ReadRoom(Id string) ([]model.RoomModel, error) {
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

func (r *RoomRepository) ReadAllRoom() ([]model.RoomModel, error) {
	db, err := handler.Init_DB()
	if err != nil {
		return nil, err
	}
	defer handler.Close_DB(db)
	returnRooms := make([]model.RoomModel, 0)
	execStr := fmt.Sprintf("exec SP_RoomBookingManagement @Process = N'%s'", "View All")
	rows, err := db.Query(execStr)
	if err != nil {
		return nil, err
	}
	var Book model.RoomModel
	for rows.Next() {
		//ตรงนี้เป็นชื่อไว้เฉยๆไมเกี่ยวกับ db
		if err := rows.Scan(&Book.Id, &Book.ReserveDate, &Book.ReserveStartTime, &Book.ReserveEndTime, &Book.RoomNo); err != nil {
			return returnRooms, err
		}
		returnRooms = append(returnRooms, Book)
	}
	return returnRooms, nil
}
