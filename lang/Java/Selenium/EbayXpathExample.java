// https://www.browserstack.com/guide/understanding-selenium-timeouts

import java.util.concurrent.TimeUnit;

import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;

public class EbayXpathExample {
  public static void main(String[] args) {
    System.setProperty("webdriver.chrome.driver", "/usr/bin/chromedriver");
    ChromeOptions chromeOptions = new ChromeOptions();
    chromeOptions.addArguments("--headless");
    chromeOptions.addArguments("--no-sandbox");

    WebDriver driver = new ChromeDriver(chromeOptions);

    driver.get("https://www.ebay.com/");

    // Implicit wait timeout for 2 seconds
    driver.manage().timeouts().implicitlyWait(2, TimeUnit.SECONDS);

    driver.findElement(By.xpath("//input[@id='gh-ac']")).sendKeys("Mobile");

    //xpath for search button
    WebElement search = driver.findElement(By.xpath("//input[@id='gh-btn']"));

    search.click();

  }
}
