package app

import (
	"net/http"
	c "rest_based_micro_go/src/controller"
)

func StartApp() {
	http.HandleFunc("/users", c.GetUser)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
