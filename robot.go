
// Command simple is a chromedp example demonstrating how to do a simple google
// search.
package main

import (
	"context"
	// "fmt"
	// "io/ioutil"
	"log"
	"time"

	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	var err error
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, click())
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("saved screenshot from search result listing `%s` (%s)", res, site)
}

func googleSearch(q, text string, site, res *string) chromedp.Tasks {
	username := "stopkung"
	password := "Stop1234"
	return chromedp.Tasks{
		// Homepage and click link to login page
		chromedp.Navigate(`https://www.prakard.com/index.php`),
		chromedp.WaitVisible(`//*[@id="nav-main"]/li[10]/a`,chromedp.BySearch),
		chromedp.Click(`//*[@id="nav-main"]/li[10]/a`,chromedp.BySearch),
		// sign in
		chromedp.WaitVisible(`//*[@id="username"]`, chromedp.BySearch),
		chromedp.WaitVisible(`//*[@id="password"]`, chromedp.BySearch),
		// chromedp.WaitVisible(`//*[@id="login"]/div[1]/div/div/fieldset/dl[4]/dd/input[3]`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="username"]`, username, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="password"]`, password, chromedp.BySearch),
		// chromedp.Click(`//*[@id="login"]/div[1]/div/div/fieldset/dl[4]/dd/input[3]`,chromedp.BySearch),

		chromedp.Sleep(150 * time.Second),
	}
}

// http://2captcha.com/in.php?key=19a8c698f9e041b93a2655787bab4360&method=userrecaptcha&googlekey=k=6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-&pageurl=https://www.google.com/recaptcha/api2/demo

// k=6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-