from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium_stealth import stealth
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
import time
import random

# Конфигурация
USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
OUTPUT_HTML = "page_content.html"
OUTPUT_SCREENSHOT = "page_screenshot.png"

def search_product(driver, query):
    try:
        WebDriverWait(driver, 20).until(
            lambda d: d.execute_script("return document.readyState") == "complete"
        )
        
        search_input = WebDriverWait(driver, 25).until(
            EC.element_to_be_clickable((By.NAME, "text"))
        )
        
        # Имитируем человеческий ввод
        for char in query:
            search_input.send_keys(char)
            time.sleep(random.uniform(0.1, 0.3))
        
        search_input.send_keys(Keys.RETURN)
        
        # Ждем прогрузки результатов
        WebDriverWait(driver, 30).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "div[data-widget*='SearchResults']"))
        )
    except Exception as e:
        print(f"Ошибка поиска: {str(e)}")
        raise

def handle_captcha(driver):
    try:
        WebDriverWait(driver, 10).until(
            EC.frame_to_be_available_and_switch_to_it((By.CSS_SELECTOR, "iframe[title*='captcha']"))
        )
        print("Обнаружена капча! Требуется ручное вмешательство.")
        driver.switch_to.default_content()
        time.sleep(120)  # Ожидание ручного ввода
    except:
        pass

def configure_options():
    options = webdriver.ChromeOptions()
    options.add_argument("--headless=new")
    options.add_argument("--disable-blink-features=AutomationControlled")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--disable-gpu")
    options.add_argument(f"user-agent={USER_AGENT}")
    options.add_argument("--window-size=1920,1080")
    # Добавляем обход Cloudflare
    options.add_argument("--disable-bundled-ppapi-flash")
    options.add_argument("--disable-site-isolation-trials")
    options.add_argument("--disable-logging")
    options.add_argument("--log-level=3")
    return options

def save_content(driver):
    """Сохранение результатов работы"""
    # Сохранение HTML
    with open(OUTPUT_HTML, "w", encoding="utf-8") as f:
        f.write(driver.page_source)
    
    # Сохранение скриншота
    driver.save_screenshot(OUTPUT_SCREENSHOT)
    print(f"\nРезультаты сохранены:\n- HTML: {OUTPUT_HTML}\n- Скриншот: {OUTPUT_SCREENSHOT}")

def parse_prices(driver):
    try:
        price_xpath = (
            "//*[contains(@class, 'tsHeadline500Medium')]"
            "[contains(., '₽')]"
            "/ancestor::div[1]"  # Ищем ближайший родительский div
        )
        
        element = WebDriverWait(driver, 25).until(
            EC.presence_of_element_located((By.XPATH, price_xpath))
        )
        
        price_text = element.text.split('₽')[0].replace(' ', '').strip()
        return int(price_text)
    
    except Exception as e:
        print(f"Ошибка парсинга: {str(e)}")
        return None
    
def get_page():
    try:
        # Инициализация драйвера
        service = Service(ChromeDriverManager().install())
        driver = webdriver.Chrome(
            service=service,
            options=configure_options()
        )

        # Настройка stealth-параметров
        stealth(driver,
                languages=["ru-RU", "ru"],
                vendor="Google Inc.",
                platform="Win32",
                webgl_vendor="Intel Inc.",
                renderer="Intel Iris OpenGL Engine",
                fix_hairline=True)

        # Маскировка WebDriver
        driver.execute_cdp_cmd("Page.addScriptToEvaluateOnNewDocument", {
            "source": """
                Object.defineProperty(navigator, 'webdriver', {
                    get: () => undefined
                })
            """
        })

        # Загрузка страницы
        driver.get("https://www.ozon.ru/")
        
        # Имитация реального пользователя
        time.sleep(random.uniform(2, 5))
        driver.execute_script("window.scrollTo(0, document.body.scrollHeight/3)")
        time.sleep(random.uniform(1, 3))
        driver.execute_script("window.scrollTo(0, document.body.scrollHeight)")
        time.sleep(random.uniform(10, 15))

        search_product(driver,"intel core i5 12400f")
        time.sleep(5)
        # Добавляем парсинг цен
        price = parse_prices(driver)
        if price:
            print(f"Успешно получена цена: {price} руб.")
        else:
            print("Цена не найдена")

    except Exception as e:
        print(f"Ошибка: {str(e)}")
    finally:
        driver.quit()

if __name__ == "__main__":
    get_page()