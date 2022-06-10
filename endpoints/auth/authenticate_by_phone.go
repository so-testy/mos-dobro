package auth_controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type authenticateByPhoneRequest struct {
	Phone string `json:"phone"`
}

// AuthenticateByPhone - эндпоинт аутентификации пользователя по номеру телефона
func (c *Controller) AuthenticateByPhone(resp http.ResponseWriter, req *http.Request) {
	// Читаем тело запроса
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Преобразуем тело запроса в структуру
	var reqData authenticateByPhoneRequest
	if err = json.Unmarshal(body, &reqData); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Аутентифицируем пользователя по номеру телефона
	if err = c.authS.AuthenticateByPhone(req.Context(), reqData.Phone); err != nil {
		// TODO: Тут обработчик ошибок
	}

	// Выстовляем корректный статус ответа
	resp.WriteHeader(http.StatusOK)
}
