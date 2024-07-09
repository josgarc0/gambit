package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josgarc0/gambit/models"
	//"github.com/josgarc0/gambit/tools"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza el registro InsertAddress")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name) "
	sentencia += " VALUES ('" + User + "','" + addr.AddAddress + "','" + addr.AddCity + "','" + addr.AddState + "','"
	sentencia += addr.AddPostalCode + "','" + addr.AddPhone + "','" + addr.AddTitle + "','" + addr.AddName + "')"

	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Insert Addresses > Ejecución Exitosa")
	return nil
}
