// https://www.selenium.dev/documentation/webdriver/getting_started/first_script/

import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;

public class HelloSelenium {
  public static void main(String[] args) {
    driver = new ChromeDriver();

    driver.get("https://google.com");

    driver.getTitle(); // => "Google"

    driver.manage().timeouts().implicitlyWait(Duration.ofMillis(500));

    WebElement searchBox = driver.findElement(By.name("q"));
    WebElement searchButton = driver.findElement(By.name("btnK"));

    searchBox.sendKeys("Selenium");
    searchButton.click();

    searchBox = driver.findElement(By.name("q"));
    searchBox.getAttribute("value"); // => "Selenium"

    driver.quit();
  }
}
