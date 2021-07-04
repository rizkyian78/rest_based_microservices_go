package controller

import (
	"encoding/json"
	"net/http"
	s "rest_based_micro_go/src/service"
	"rest_based_micro_go/src/utils"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User Must Be Number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(http.StatusNotFound)
		res.Write(jsonValue)
		return
	}

	user, apiErr := s.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
