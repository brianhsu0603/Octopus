import requests
from bs4 import BeautifulSoup
from selenium import webdriver
import time
from threading import Thread
import os
from dotenv import load_dotenv
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By

option = webdriver.ChromeOptions()
#option.add_argument("--headless")
option.add_argument("window-size=2560,1440")
driver = webdriver.Chrome(executable_path="./chromedriver", options = option)
driver.get("https://popcat.click")
#button = driver.find_element_by_xpath('//*[@id="app"]/div/div')
button = driver.findElement(By.xpath('//*[@id="app"]/div/div'))
while(True):
    button.click()
