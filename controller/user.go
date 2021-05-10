package controller

import (
	"encoding/json"
	"fmt"
	"github.com/noguchidaisuke/mvc-plus-service/model"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	id, _ := strconv.Atoi(idStr)

	user := model.User{
		ID: id,
		Name: name,
	}

	model.UpdateUser(&user)

	b, err := json.Marshal(model.Users)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}