package store

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Repository Repository
}

type Response struct {
	Message string      `json:"message"`
	Result  int         `json:"result"`
	Data    interface{} `json:"data"`
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err.Error())
	}

	Username := r.FormValue("username")
	Password := r.FormValue("password")

	user, err := c.Repository.login(Username, Password)

	var response Response

	if err != nil {
		response = Response{
			Message: err.Error(),
			Result:  http.StatusInternalServerError,
			Data:    struct{}{},
		}
	} else {
		response = Response{
			Message: "",
			Result:  http.StatusOK,
			Data:    user,
		}
	}

	defer printOut(w, response)

	return
}

func (c *Controller) getAllMovie(w http.ResponseWriter, r *http.Request) {
	result := c.Repository.getAllMovie()

	response := Response{
		Message: "",
		Result:  http.StatusOK,
		Data:    result,
	}

	printOut(w, response)

	return
}

func printOut(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(r.Result)
	json.NewEncoder(w).Encode(r)
}
