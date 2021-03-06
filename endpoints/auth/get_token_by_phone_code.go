package auth_controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type getTokenByPhoneCodeRequest struct {
	SmsCode string `json:"sms_code"`
}

// GetTokenByPhoneCode - эндпоинт получения JWT токена по sms коду
func (c *Controller) GetTokenByPhoneCode(resp http.ResponseWriter, req *http.Request) {
	// Читаем тело запроса
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Преобразуем тело запроса в структуру
	var reqData getTokenByPhoneCodeRequest
	if err = json.Unmarshal(body, &reqData); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Аутентифицируем пользователя по номеру телефона
	jwt, err := c.authS.GetTokenByPhoneCode(req.Context(), reqData.SmsCode)
	if err != nil {
		// TODO: Тут обработчик ошибок
	}

	// Добавляем токен в Header
	resp.Header().Add("Authorization", jwt)
	resp.WriteHeader(http.StatusOK)
}
