package store

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/structs"
)

type Controller struct {
	Repository Repository
}

type Response struct {
	Message string                 `json:"message"`
	Result  int                    `json:"result"`
	Data    map[string]interface{} `json:"data"`
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err.Error())
	}

	Username := r.FormValue("username")
	Password := r.FormValue("password")

	user := c.Repository.login(Username, Password)

	var response Response
	response = Response{
		Message: "",
		Result:  http.StatusOK,
		Data:    structs.Map(user),
	}
	defer printOutput(w, response)

	return
}

func printOutput(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(r.Result)
	json.NewEncoder(w).Encode(r)
}
