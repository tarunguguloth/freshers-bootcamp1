package Models

import (
	"freshers_bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"
	"sync"


)
var Mutex sync.Mutex

//CreateCustpmer...Insert New data
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//Delete customer
func DeleteCustomer(customer *Customer, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(customer)
	return nil
}

//GetCustomerByID ... Fetch only one product by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}