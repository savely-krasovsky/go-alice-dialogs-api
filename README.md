# go-alice-dialogs-api
[![GoDoc](https://godoc.org/github.com/L11R/go-alice-dialogs-api?status.svg)](https://godoc.org/github.com/L11R/go-alice-dialogs-api)

Микро-библиотека для упрощения работы с новой платформой Яндекс.Диалоги. Проще всего показать всё сразу на примере:
```golang
package main

import (
	"go-alice-dialogs-api"
	"net/http"
	"strings"
)

func HandleDialog(req aliceapi.Request) aliceapi.Response {
	// Стандартный ответ
	text := "Что?"
	button := aliceapi.Button{
		Title: "Кнопка!",
		URL:   "https://ru.wikipedia.org/wiki/Hello%2C_world!",
	}

	// Обрабатываем "Привет!" от пользователя
	if strings.Contains(strings.ToLower(req.Command), "привет") {
		text = "Тебе тоже!"
	}

	if req.Command == button.Title {
		text = "Похоже ты затестил кнопку!"
	}

	// Отправляем готовый ответ Алисе
	return aliceapi.Response{
		Text:    text,
		Buttons: []aliceapi.Button{button},
	}
}

func main() {
	// Регистрируем главный обработчик
	aliceapi.Handle("/", HandleDialog)

	// Поднимаем собственный вебсервер и слушаем вебхуки
	http.ListenAndServeTLS("0.0.0.0:7443", "cert.pem", "key.pem", nil)
}
```
