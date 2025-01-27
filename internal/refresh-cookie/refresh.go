package refresh

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func Refresh() string {
	// Укажите путь к ChromeDriver
	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)
	if err != nil {
		log.Fatalf("Failed to start ChromeDriver: %v", err)
	}
	defer service.Stop()

	// Настройка capabilities для Chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"chromeOptions": map[string]interface{}{
			"args": []string{
				"--headless", // Запуск в headless-режиме (без графического интерфейса)
				"--disable-gpu",
				"--no-sandbox",
				"--disable-dev-shm-usage",
			},
		},
	}

	// Подключение к WebDriver
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		log.Fatalf("Failed to connect to WebDriver: %v", err)
	}
	defer wd.Quit()

	// Переход на страницу
	err = wd.Get("https://www.ozon.ru/category/smartfony-15502/apple-26303000")
	if err != nil {
		log.Fatalf("Failed to navigate to the URL: %v", err)
	}

	// Ожидание загрузки страницы
	time.Sleep(10 * time.Second)

	// Получение куки
	cookies, err := wd.GetCookies()
	if err != nil {
		log.Fatalf("Failed to get cookies: %v", err)
	}

	// Форматирование куки в строку
	var cookieStrings []string
	for _, cookie := range cookies {
		cookieStrings = append(cookieStrings, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}
	cookieLine := strings.Join(cookieStrings, "; ")

	return cookieLine
}
