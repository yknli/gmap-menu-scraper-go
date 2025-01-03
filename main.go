package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func FindElement(driver selenium.WebDriver, by, value string) (selenium.WebElement, error) {
	var element selenium.WebElement
	var err error
	for i := 0; i < 5; i++ {
		element, err = driver.FindElement(by, value)
		if err == nil {
			return element, nil
		}
		time.Sleep(2 * time.Second)
	}
	return nil, err
}

func main() {
	startTime := time.Now()

	const (
		seleniumPath     = "./selenium/selenium-server-standalone.jar"
		chromeDriverPath = "./chrome/chromedriver"
		port             = 8081
	)

	opts := []selenium.ServiceOption{
		selenium.Output(nil), // Output debug information to STDERR
	}

	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome",
		"goog:chromeOptions": map[string]interface{}{
			"args": []string{
				"--headless",
				"--disable-gpu",
				"--window-size=1920,1080",
			},
		}}

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer driver.Quit()

	restaurantName := "三奇壹號咖啡館x築甜製菓"
	driver.Get(fmt.Sprintf("http://maps.google.com/?q=%s", restaurantName))

	photosHeader, err := FindElement(driver, selenium.ByXPATH, `//h2[contains(string(), "相片和影片")]`)
	if err != nil {
		time.Sleep(10 * time.Second)
		panic(err)
	}

	_, err = driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{photosHeader})
	if err != nil {
		panic(err)
	}

	nextButton, err := FindElement(driver, selenium.ByXPATH, `//button[@aria-label="下一張相片"]`)
	if err != nil {
		time.Sleep(10 * time.Second)
		panic(err)
	}
	_, err = driver.ExecuteScript("arguments[0].click();", []interface{}{nextButton})
	if err != nil {
		panic(err)
	}

	menuButton, err := FindElement(driver, selenium.ByXPATH, `//button[@aria-label="菜單"]/img`)
	if err != nil {
		time.Sleep(10 * time.Second)
		panic(err)
	}

	err = menuButton.Click()
	if err != nil {
		panic(err)
	}

	var menuPhotoDivs []selenium.WebElement
	for {
		menuPhotoDivs, err = driver.FindElements(selenium.ByXPATH, `//a[@data-photo-index]/div[1]/div[1]`)
		if err == nil && len(menuPhotoDivs) == 20 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	photoDir := "photos"
	err = os.MkdirAll(fmt.Sprintf("%s/%s", photoDir, restaurantName), os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, photoDiv := range menuPhotoDivs {

		for {
			displayed, err := photoDiv.IsDisplayed()
			if err == nil {
				if displayed {
					driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{photoDiv})
					break
				}
				time.Sleep(1 * time.Second)
			} else {
				panic(err)
			}
		}

		photoDivStyle, err := photoDiv.GetAttribute("style")
		if err != nil {
			panic(err)
		}

		startPoint := "https"
		urlStartIndex := strings.Index(photoDivStyle, startPoint)
		if urlStartIndex == -1 {
			panic("photo url not found")
		}
		urlEndIndex := strings.Index(photoDivStyle, `");`)
		photoURL := photoDivStyle[urlStartIndex:urlEndIndex]
		fmt.Println(photoURL)

		photoName := fmt.Sprintf("%s/%s/%s.jpg", photoDir, restaurantName, photoURL[strings.LastIndex(photoURL, "/")+1:])
		resp, err := http.Get(photoURL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(photoName, body, 0644)
		if err != nil {
			panic(err)
		}
	}

	endTime := time.Now()
	fmt.Println("Duration: ", endTime.Sub(startTime).String())
}
