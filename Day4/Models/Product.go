package Models



import (
	"fmt"

	"freshers_bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"
)



//CreateUser ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProducts(product *[]Product) (err error) {
	Mutex.Lock()
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	defer Mutex.Unlock()
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	Mutex.Lock()
	defer Mutex.Unlock()
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Mutex.Lock()
	defer Mutex.Unlock()
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}