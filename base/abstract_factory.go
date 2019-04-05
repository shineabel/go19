package main

import "fmt"

type OrderDAO interface {
	SaveOrder()
}
type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderDAO() OrderDAO
	CreateOrderDetailDAO() OrderDetailDAO
}


type RDBMSOrderDAO struct {

}

func (d *RDBMSOrderDAO) SaveOrder()  {
	fmt.Println("save order in rdbms...")
}

type RDBMSOrderDetailDAO struct {

}

func (d *RDBMSOrderDetailDAO) SaveOrderDetail()  {
	fmt.Println("save order detail in rdbms...")
}

type RDBMSDAOFactory struct {

}


type FileDAOFactory struct {

}

func (df *RDBMSDAOFactory) CreateOrderDAO()  OrderDAO {
	return &RDBMSOrderDAO{}
}

func (df *RDBMSDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {

	return &RDBMSOrderDetailDAO{}
}

func (df *FileDAOFactory)CreateOrderDAO() OrderDAO {
	return &FileOrderDAO{}
}
func (df *FileDAOFactory) CreateOrderDetailDAO()  OrderDetailDAO{
	return &FileOrderDetailDAO{}
}


type FileOrderDAO struct {

}
type FileOrderDetailDAO struct {

}

func (d *FileOrderDAO) SaveOrder()  {
	fmt.Println("save order in file...")
}

func (d *FileOrderDetailDAO) SaveOrderDetail()  {
	fmt.Println("save order detail in file...")
}