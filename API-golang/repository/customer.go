package repository

import (
	"customer/handler"
	"customer/model"
	"fmt"
)

type CustomerRepository struct{}

func (r *CustomerRepository) ReadCustomer(Id string) ([]model.CustomerModel, error) {
	db, err := handler.Init_DB()
	if err != nil {
		return nil, err
	}
	defer handler.Close_DB(db)
	returncustomers := make([]model.CustomerModel, 0)
	execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s', @Id = '%s' ", "FIND", Id)
	rows, err := db.Query(execStr)
	if err != nil {
		return nil, err
	}
	var customer model.CustomerModel
	for rows.Next() {
		if err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address, &customer.Birthday); err != nil {
			return returncustomers, err
		}
		returncustomers = append(returncustomers, customer)
	}
	return returncustomers, nil
}

func (r *CustomerRepository) ReadAllCustomer() ([]model.CustomerModel, error) {
	db, err := handler.Init_DB()
	if err != nil {
		return nil, err
	}
	defer handler.Close_DB(db)
	returncustomers := make([]model.CustomerModel, 0)
	execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s'", "View All")
	rows, err := db.Query(execStr)
	if err != nil {
		return nil, err
	}
	var customer model.CustomerModel
	for rows.Next() {
		//ตรงนี้เป็นชื่อไว้เฉยๆไมเกี่ยวกับ db
		if err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address, &customer.Birthday); err != nil {
			return returncustomers, err
		}
		returncustomers = append(returncustomers, customer)
	}
	return returncustomers, nil
}

func (r *CustomerRepository) CreateCustomer(data model.CustomerModel) error {
	db, err := handler.Init_DB()
	if err != nil {
		return err
	}
	defer handler.Close_DB(db)
	execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s', @FirstName = N'%s', @LastName = N'%s' , @Address = N'%s' , @Birthday = N'%v'", "CREATE", data.FirstName, data.LastName, data.Address, data.Birthday.Format("2006-01-02"))
	_, err = db.Exec(execStr)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) UpdateCustomer(data model.CustomerModel) error {
	db, err := handler.Init_DB()
	if err != nil {
		return err
	}
	defer handler.Close_DB(db)
	execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s', @FirstName = N'%s', @LastName = N'%s' , @Address = N'%s' , @Birthday = N'%v', @Id = '%v'", "UPDATE", data.FirstName, data.LastName, data.Address, data.Birthday.Format("2006-01-02"), data.Id)
	_, err = db.Exec(execStr)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) DeleteCustomer(Id string) error {
	db, err := handler.Init_DB()
	if err != nil {
		return err
	}
	defer handler.Close_DB(db)
	execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s', @Id = '%s'", "DELETE", Id)
	_, err = db.Exec(execStr)
	if err != nil {
		return err
	}
	return nil
}
