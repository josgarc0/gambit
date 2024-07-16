package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	//	"strconv"
	//"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josgarc0/gambit/models"
	//"github.com/josgarc0/gambit/tools"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza Registro Orders")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "Insert into orders (Order_UserUUID, Order_Total, Order_AddId) values ('"
	sentencia += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + " )"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	for _, od := range o.OrdersDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentencia += "," + strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', 1, 64) + ")"
		fmt.Println(sentencia)
		_, err = Db.Exec(sentencia)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}
	fmt.Println("Insert Order > Ejecución Exitosa")
	return LastInsertId, nil
}
