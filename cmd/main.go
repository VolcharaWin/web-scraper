package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// Запуск ChromeDriver
	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)
	if err != nil {
		log.Fatalf("Failed to start ChromeDriver: %v", err)
	}
	defer service.Stop()

	// Настройка capabilities для Chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"goog:chromeOptions": map[string]interface{}{
			"args": []string{
				"--headless", // Включаем headless-режим
				"--disable-gpu",
				"--no-sandbox",
				"--disable-dev-shm-usage",
				"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
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
	err = wd.Get("https://www.ozon.ru")
	if err != nil {
		log.Fatalf("Failed to navigate to the URL: %v", err)
	}

	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Получение заголовка страницы
	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get page title: %v", err)
	}

	fmt.Println("Page title:", title)
}
