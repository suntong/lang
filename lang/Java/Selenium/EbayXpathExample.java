// https://www.browserstack.com/guide/understanding-selenium-timeouts

import java.util.concurrent.TimeUnit;
import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;

public class EbayXpathExample {
  public static void main(String[] args) {

    System.setProperty("webdriver.chrome.driver", "Path of Chrome Driver");
    WebDriver driver = new ChromeDriver();
    driver.get("https://www.ebay.com/");

    // Implicit wait timeout for 20seconds
    driver.manage().timeouts().implicitlyWait(20, TimeUnit.SECONDS);

    driver.findElement(By.xpath("//input[@id='gh-ac']")).sendKeys("Mobile");

    //xpath for search button
    WebElement search = driver.findElement(By.xpath("//input[@id='gh-btn']"));

    search.click();

  }
}