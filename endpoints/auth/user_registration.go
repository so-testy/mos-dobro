package auth_controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type userRegistrationRequest struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	MiddleName string `json:"middle_name"`
	Birthdate  string `json:"birthdate"`
	Gender     uint   `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

// UserRegistration - эндпоинт регистрирующий пользователя
func (c *Controller) UserRegistration(resp http.ResponseWriter, req *http.Request) {
	// Читаем тело запроса
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Преобразуем тело запроса в структуру
	var reqData userRegistrationRequest
	if err = json.Unmarshal(body, &reqData); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
