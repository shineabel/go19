package main

import "testing"

func Save(f DAOFactory)  {
	f.CreateOrderDAO().SaveOrder()
	f.CreateOrderDetailDAO().SaveOrderDetail()

}

func TestDAOFactory( t *testing.T)  {



	df := &RDBMSDAOFactory{}
	Save(df)

	df2 := &FileDAOFactory{}
	Save(df2)
}
