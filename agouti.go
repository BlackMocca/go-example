package main

import (
	"github.com/sclevine/agouti"
	"log"
)

// http://todomvc.com/examples/react/
func main() {

	driver := agouti.ChromeDriver()
	// driver := agouti.ChromeDriver(
		// agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
	// )
		
	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}
	
	  page, err := driver.NewPage()
	  if err != nil {
		log.Fatal("Failed to open page:", err)
	  }

	  if err := page.Navigate("https://www.prakard.com"); err != nil {
		log.Fatal("Failed to navigate:", err)
	  }

	  // Click Link
	  if err := page.FindByXPath(`//*[@id="nav-main"]/li[10]/a`).Click(); err != nil {
		log.Fatal("Failed to click:", err)
	  }
	  // Input username
	  input := page.FindByXPath(`//*[@id="username"]`)
	  log.Println(input)

	  if err := input.Fill("stopkung"); err != nil {
		  log.Fatal("Input username:", err)
	  }
	  // input password
	  if err := page.FindByXPath(`//*[@id="password"]`).Fill("Stop1234"); err != nil {
		log.Fatal("Input password:", err)
	}
	// click submit
		if err := page.FindByButton(`Login`).Click(); err != nil {
			log.Fatal("Login:", err)
		}


	  
	
	//   if err := driver.Stop(); err != nil {
		// log.Fatal("Failed to close pages and stop WebDriver:", err)
	//   }
}