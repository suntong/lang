// https://www.selenium.dev/documentation/webdriver/getting_started/first_script/

//import java.time.Duration;
import java.util.concurrent.TimeUnit;

import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;

public class HelloSelenium {
  public static void main(String[] args) {
    System.setProperty("webdriver.chrome.driver", "/usr/bin/chromedriver");
    ChromeOptions chromeOptions = new ChromeOptions();
    chromeOptions.addArguments("--headless");
    chromeOptions.addArguments("--no-sandbox");

    WebDriver driver = new ChromeDriver(chromeOptions);

    driver.get("https://google.com");

    driver.getTitle(); // => "Google"

    driver.manage().timeouts().implicitlyWait(1, TimeUnit.SECONDS);

    WebElement searchBox = driver.findElement(By.name("q"));
    WebElement searchButton = driver.findElement(By.name("btnK"));

    searchBox.sendKeys("Selenium");
    searchButton.click();

    searchBox = driver.findElement(By.name("q"));
    searchBox.getAttribute("value"); // => "Selenium"

    driver.quit();
  }
}
