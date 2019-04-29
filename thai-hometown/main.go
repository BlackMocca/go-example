package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
	"strconv"
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
	
		  if err := page.Navigate("https://www.thaihometown.com/member/"); err != nil {
			log.Fatal("Failed to navigate:", err)
		  }
		
		/* 
		  
		  Sign in

		*/

		
	
		if err := page.FindByXPath(`//*[@id="loginform"]/tbody/tr[2]/td[2]/input`).Fill("stopkung@hotmail.com"); err != nil {
			log.Fatal("fill username ", err)
		}
		if err := page.FindByXPath(`//*[@id="loginform"]/tbody/tr[3]/td[2]/input`).Fill("Stop1234"); err != nil {
			log.Fatal("fill password ", err)
		}
		if err := page.FindByXPath(`//*[@id="loginform"]/tbody/tr[4]/td[2]/table/tbody/tr/td[1]`).Click(); err != nil {
			log.Fatal("Sign in submit : ", err)
		}

		/* 

			post condo

		*/
		
		page.Navigate(`https://www.thaihometown.com/addnew`)

		if err := page.FindByID(`headtitle`).Fill("ขายคอนโดมิเนี่ยมสุดหรูอู้ฟู่"); err != nil {
			log.Fatal(`topic : `, err)
		}		

		if err := page.FindByID(`typepart`).Select("ประกาศให้เช่า"); err != nil {
			log.Fatal(` : `, err)
		}		

		if err := page.FindByID(`property_type`).Select("คอนโดมิเนียม"); err != nil {
			log.Fatal(` : `, err)
		}	

		if err := page.FindByID(`property_area`).Fill("23"); err != nil {
			log.Fatal(` : `, err)
		}	

		if err := page.FindByID(`property_sqm`).Select("ตร.ม"); err != nil {
			log.Fatal(` : `, err)
		}	

		if err := page.FindByID(`room1`).Select("3"); err != nil {
			log.Fatal(` bedroom : `, err)
		}	

		if err := page.FindByID(`room2`).Select("2"); err != nil {
			log.Fatal(` toilet : `, err)
		}	

		if err := page.FindByName(`conditioning`).Select("2 เครื่อง"); err != nil {
			log.Fatal(`  : `, err)
		}	

		if err := page.FindByID(`carpark`).Select("3 คัน"); err != nil {
			log.Fatal(`  : `, err)
		}	

		if err := page.FindByID(`property_city_bkk`).Select("ราษฎร์บูรณะ"); err != nil {
			log.Fatal(`  : `, err)
		}	
		
		if err := page.FindByID(`ad_title`).Fill("AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น AD นี้ต้องใส่ถึง 200 ตัวขึ้นไป นี่ที่พิมๆไปก็ไม่รู้ว่าถึงหรือยัง หากยังไม่ถึงเราจะบอกให้ว่าโพสนี้เป้นเพียงแค่ทดลองโพสไม่มีส่วนเกี่ยวข้องกับการขายจริงใดๆทั้งสิ้น"); err != nil {
			log.Fatal(`  : `, err)
		}	
		
		
		property := [10]int{1,3,8,21,7,40,23,15,18,30}
			for _, elementID := range property {
				prop_id := "check_ser" + strconv.Itoa(elementID)
				if err := page.FindByID(prop_id).Check(); err != nil {
					log.Fatal(`  : `, err)
				}
			}

		if err := page.FindByID(`property_bts`).Select("วุฒากาศ (S11)"); err != nil {
			log.Fatal(`  : `, err)
		}	

		if err := page.FindByID(`selling_price`).Fill(`0`); err != nil {
			log.Fatal(`  : `, err)
		}	

		if err := page.FindByID(`rent_price`).Fill(`25000`); err != nil {
			log.Fatal(`  : `, err)
		}	

		time.Sleep(2* time.Second)
		if err := page.FindByID(`type_forrent`).Select(`ต่อเดือน`); err != nil {
			log.Fatal(`  : `, err)
		}	

		if err := page.FindByID(`price_unit`).Fill(`1500`); err != nil {
			log.Fatal(`  : `, err)
		}	
		
		if err := page.FindByID(`typeunit1`).Check(); err != nil {
			log.Fatal(`  : `, err)
		}			


		info := [20]int{1,3,8,21,7,40,23,15,18,30,35,38,42,46,48,70,100,80,90,134}

		for _, elementID := range info {
			prop_id := "info" + strconv.Itoa(elementID)
			if err := page.FindByID(prop_id).Check(); err != nil {
				log.Fatal(`  : `, err)
			}
		}

		src, _ := page.FindByXPath(`//*[@id="inAddnew"]/table/tbody/tr[5]/td/table/tbody/tr/td/div[1]/img`).Attribute("src")

		number_code := src[(len(src)-4) : len(src)]

		log.Println(src)
		log.Println("number code : ",number_code)

		page.FindByID(`string1`).Fill(number_code)

		page.FindByXPath(`//*[@id="inAddnew"]/table/tbody/tr[5]/td/table/tbody/tr/td/img`).Click()
		


}
