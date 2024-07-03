package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//	"strconv"
	//"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josgarc0/gambit/models"
	"github.com/josgarc0/gambit/tools"
	//"github.com/josgarc0/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza Registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "Insert into category (Categ_Name, Categ_Path) values ('" + c.CategName + "','" + c.CategPath + "')"

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
	fmt.Println("Insert CAtegory > Ejecución Exitosa")
	return LastInsertId, err2
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza Registro de UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "Update category set "
	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}
	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}

		sentencia += " Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}
	sentencia += " Where Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update CAtegory > Ejecución Exitosa")
	return nil

}

func DeleteCategory(id int) error {
	fmt.Println("Comienza Delete de DeleteCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "Delete from category where Categ_ID = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Delete CAtegory > Ejecución Exitosa")
	return nil

}
