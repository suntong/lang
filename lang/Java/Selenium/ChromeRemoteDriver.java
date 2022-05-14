// http://selftechy.com/2011/08/17/running-selenium-tests-with-chromedriver-on-linux

import java.io.File;
import java.io.IOException;
import java.net.*;

//import org.junit.*;
//import org.junit.runner.RunWith;
//import org.junit.runners.BlockJUnit4ClassRunner;
import org.openqa.selenium.*;
import org.openqa.selenium.chrome.ChromeDriverService;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.remote.RemoteWebDriver;

public class ChromeRemoteDriver {

  public static void main(String[] args) throws MalformedURLException {
    new DesiredCapabilities();
    URL serverurl = new URL("http://localhost:9515");
    DesiredCapabilities capabilities = DesiredCapabilities.chrome();
    WebDriver driver = new RemoteWebDriver(serverurl, capabilities);
    driver.get("http://www.google.com");
    WebElement searchEdit = driver.findElement(By.name("q"));
    searchEdit.sendKeys("Selftechy on google");
    searchEdit.submit();

  }
}