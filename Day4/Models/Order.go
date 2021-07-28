package Models

import (
	"freshers_bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"

)



func GetOrders(order *[]Product) (err error) {
	Mutex.Lock()
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	defer Mutex.Unlock()
	return nil
}

//CreateUser ... Insert New data
func CreateOrder(user *Order) (err error) {
	Mutex.Lock()
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	defer Mutex.Unlock()
	return nil
}

func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}