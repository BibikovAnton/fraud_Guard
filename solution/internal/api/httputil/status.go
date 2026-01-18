package httputil

const (
	StatusOK                  = 200 // Успешный GET/PUT
	StatusCreated             = 201 // Успешный POST (создание ресурса), включая DECLINED транзакции
	StatusNoContent           = 204 // Успешный DELETE
	StatusBadRequest          = 400 // Невалидный JSON, неподдерживаемый Content-Type
	StatusUnauthorized        = 401 // Отсутствует/невалидный/истёкший токен
	StatusForbidden           = 403 // Недостаточно прав (роль или доступ к чужому ресурсу)
	StatusNotFound            = 404 // Ресурс не найден
	StatusConflict            = 409 // Конфликт (email занят, имя правила занято)
	StatusUnprocessableEntity = 422 // Ошибка валидации полей
	StatusLocked              = 423 // Пользователь деактивирован
)
