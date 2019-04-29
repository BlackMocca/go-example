package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func main() {
		driver := agouti.ChromeDriver()
		// driver := agouti.ChromeDriver(
		// 	agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
		// )
		if err := driver.Start(); err != nil {
			log.Fatal("Failed to start driver:", err)
		}
		
		  page, err := driver.NewPage()
		  if err != nil {
			log.Fatal("Failed to open page:", err)
		  }
	
		  if err := page.Navigate("https://www.prakard.com/ucp.php?mode=login"); err != nil {
			log.Fatal("Failed to navigate:", err)
		  }
		
		/* 
		  
		  Sign in

		*/

		if err := page.FindByID(`username`).Fill("stopkung"); err != nil {
			log.Fatal("fill username ", err)
		}
		if err := page.FindByID(`password`).Fill("Stop1234"); err != nil {
			log.Fatal("fill password ", err)
		}
		if err := page.FindByXPath(`//*[@id="login"]/div[1]/div/div/fieldset/dl[4]/dd/input[3]`).Click(); err != nil {
			log.Fatal("Sign in submit : ", err)
		}

		/* 

			post condo

		*/
		page.FindByXPath(`//*[@id="nav-main"]/li[9]/a`).Click()
		
		time.Sleep(5* time.Second)
		
		if err := page.FindByID(`category_post`).Select("Condo / Condominium - คอนโด คอนโดมิเนียม"); err != nil {
			log.Fatal("link to post condo => category ", err)
		}

		if err := page.FindByID(`area_post`).Select("Bangsue / Bangkhen Area"); err != nil {
			log.Fatal("link to post condo => area ", err)
		}

		if err := page.FindByID(`project_post`).Select("333 ริเวอร์ไซด์ (ประชาราษฎร์ สาย1)"); err != nil{
			log.Fatal("link to post condo => project ", err)
		}

		page.FindByID(`btn-post`).Click();

			// detail page post

		if err := page.FindByID(`subject`).Fill("Condominium"); err != nil {
			log.Fatal("enter topic post fail => ", err)
		}
		if err := page.FindByID(`message`).Fill("รายละเอียด ของ condominium"); err != nil {
			log.Fatal("enter topic post fail => ", err)
		}
		
		if err := page.FindByID(`attach-panel-multi`).Find(`input[type=file]`).UploadFile("condo_1.jpg"); err != nil {
			log.Fatal("upload Image fail ", err)
		}

		if err := page.FindByID(`attach-panel-multi`).Find(`input[type=file]`).UploadFile("condo_2.jpg"); err != nil {
			log.Fatal("upload Image fail ", err)
		}

		if err := page.FindByID(`attach-panel-multi`).Find(`input[type=file]`).UploadFile("condo_3.jpg"); err != nil {
			log.Fatal("upload Image fail ", err)
		}
		
			// Submit form

		time.Sleep(30 * time.Second)
		 
		if err := page.FindByXPath(`//*[@id="postform"]/div[2]/div/fieldset/input[4]`).Click(); err != nil {
			log.Fatal("Submit form ", err)
		}
		log.Println("Flow post success")
		



		
}
