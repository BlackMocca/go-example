package main

import (
	"github.com/sclevine/agouti"
	"log"
)

type CondoList struct {
	Id int
	Name string
	Address string
}

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


				// Login Renthub
				if err := page.Navigate(`https://www.renthub.in.th/login`); err != nil {
					log.Fatal("Link to login page error ", err)
				}
				// fill email password
				if err := page.FindByID(`user_email`).Fill("stopkung@hotmail.com"); err != nil {
					log.Fatal("fill username ", err)
				}
				if err := page.FindByID(`user_password`).Fill("Stop1234"); err != nil {
					log.Fatal("fill username ", err)
				}
				if err := page.FindByButton(`Sign in`).Click(); err != nil {
					log.Fatal("commit:", err)
				}

				if err := page.Navigate(`https://www.renthub.in.th/condo_listings/new`); err != nil {
					log.Fatal("Link to post condominium")
				}

			var data []CondoList
		page.RunScript(`
			console.log("test")
			var test = [
				{
					"id": 241,
					"name": "101 Mansion (101 แมนชั่น)",
					"address": "คลองจั่น บางกะปิ กรุงเทพมหานคร"
				},
				{
					"id": 2034,
					"name": "103 Condominium 1 (103 คอนโดมิเนียม 1)",
					"address": "ศิริมังคลาจารย์ สุเทพ เมืองเชียงใหม่ เชียงใหม่"
				},
			]
				$.ajax({url:"/condo_listings/search_project",data:{name:"the"},type:"GET",error:function(){t()},success:function(e){ return e; }})
			`, map[string]interface{}{}, &data)
		log.Println(data)

		// $.ajax({url:"/condo_listings/search_project",data:{name:e},type:"GET",error:function(){t()},success:function(e){n.clearOptions(),n.refreshOptions(),t(e)}})}
}