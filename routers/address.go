package routers

import (
	"encoding/json"

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
