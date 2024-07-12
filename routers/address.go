package routers

import (
	"encoding/json"
	"strconv"

	"github.com/josgarc0/gambit/bd"
	"github.com/josgarc0/gambit/models"
)

func InsertAddress(body string, User string) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos"
	}
	if len(t.AddAddress) == 0 {
		return 400, "Debe especificar el Address"
	}
	if t.AddName == "" {
		return 400, "Debe especificar el Name"
	}
	if t.AddTitle == "" {
		return 400, "Debe especificar el Title"
	}
	if t.AddCity == "" {
		return 400, "Debe especificar el City"
	}
	if t.AddPhone == "" {
		return 400, "Debe especificar el Phone"
	}
	if t.AddPostalCode == "" {
		return 400, "Debe especificar el PostalCode"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del address para el ID del Usuario " + User + " > " + err.Error()
	}

	return 200, "InsertAddress Ok "
}

func UpdateAddress(body string, User string, id int) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}
	t.AddId = id

	var encontrado bool
	err, encontrado = bd.AddressExists(User, t.AddId)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Address para el usuario " + User + " > " + err.Error()
		}
		return 400, "no se encuentra un registro de ID de Usuario asociado a esa ID de Address"
	}

	err2 := bd.UpdateAddress(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del address " + strconv.Itoa(id) + " > " + err2.Error()
	}
	return 200, "UpdateAddress Ok"
}
