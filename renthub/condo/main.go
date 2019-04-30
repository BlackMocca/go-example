package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
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
		
		// post condominium
			if err := page.Navigate(`https://www.renthub.in.th/condo_listings/new`); err != nil {
				log.Fatal("Link to post condominium")
			}
			// fill input
			if err := page.FindByID(`condo_listing_title`).Fill("[Test]ขายห้อง คอนโดมิเนี่ยนสาทร"); err != nil {
				log.Fatal("หัวข้อประกาศ[Thai] : ",err)
			}
			if err := page.FindByID(`english_title`).Fill("[Test]Selling Condominium at Sathorn"); err != nil {
				log.Fatal("หัวข้อประกาศ[English] : ",err)
			}
			// search โครงการ
			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[3]/div/div/div/div[1]/input`).Fill("Kulapapruek Grand Park Udon Thani");err != nil {
				log.Fatal("search : ",err)
			} 

			// wait for Ajax data โครงการ
			time.Sleep(2 * time.Second)
			
			// select parent element from ajax data
			elements := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[3]/div/div/div/div[2]/div`).All("div"); 
			log.Println("dropdown โครงการ",elements)
				// select element
			if err := elements.At(0).Click(); err != nil {
				log.Fatal("Mouse to element")
			}

			//detail of room
			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[4]/div/span[1]/div/div[1]/input`).Fill("อาคารบ้านรักษา"); err != nil {
				log.Fatal("ระบุอาคาร : ",err)
			}

			
			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[4]/div/span[2]/div/div[1]/input`).Fill("3 ห้องนอน"); err != nil{
				log.Fatal("input ประเภทห้อง : ", err)
			}
			rooms := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[4]/div/span[2]/div/div[2]/div`).All("div")
				log.Println("Drop down ห้อง : ",rooms)
				if err := rooms.At(0).Click(); err != nil {
					log.Fatal("Mouse to element")
				}

			if err := page.FindByID(`room_information_room_home_address`).Fill("6/367"); err != nil {
				log.Fatal("บ้านเลขที่ ", err)
			}

			if err := page.FindByID(`room_information_room_no`).Fill("2301"); err != nil {
				log.Fatal("ห้องที่ ", err)
			}

			if err := page.FindByID(`room_information_on_floor`).Fill("3"); err != nil {
				log.Fatal("ชั้นที่ ", err)
			}

			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[5]/div/span[4]/div/div[1]/input`).Fill("ทิศตะวันออก"); err != nil {
				log.Fatal("ทิศ ", err)
			}
			directions := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[5]/div/span[4]/div/div[2]/div`).All("div");
				log.Println("dropdown direction : ",directions)
				if err := directions.At(0).Click(); err != nil {
					log.Fatal("Mouse to Element")
				}

			if err := page.FindByID(`room_information_no_of_bed`).Fill("3"); err != nil {
				log.Fatal("ห้องนอน ", err)
			}

			if err := page.FindByID(`room_information_no_of_bath`).Fill("2"); err != nil {
				log.Fatal("ห้องน้ำ ", err)
			}

			if err := page.FindByID(`room_information_room_area`).Fill("17"); err != nil {
				log.Fatal("พื้นที่ห้อง ", err)
			}
			
			if err := page.FindByID(`room_information_remark`).Fill("บันทึกช่วยจำ: ทดสอบประกาศ"); err != nil {
				log.Fatal("บันทึกช่วยจำ ", err)
			}

			if err := page.FindByID(`room_information_remark`).Fill("บันทึกช่วยจำ: ทดสอบประกาศ"); err != nil {
				log.Fatal("บันทึกช่วยจำ ", err)
			}

			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[2]/ul/li[5]/div/span[4]/div/div[1]/input`).Fill("ทิศตะวันออก"); err != nil {
				log.Fatal("ทิศ ", err)
			}
			
			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[3]/ul/li[1]/div/div/div[1]/input`).Fill("เช่า"); err != nil {
				log.Fatal("input ประเภทประกาศ ", err)
			}
			category_post := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[3]/ul/li[1]/div/div/div[2]/div`).All("div");
				log.Println("dropdown category_post : ",category_post)
				if err := category_post.At(0).Click(); err != nil {
					log.Fatal("Mouse to Element")
				}

			// ข้อมูลให้เช่า
			rental_section := page.FindByID(`rental_price`).AllByClass("section"); 
			
			rental_rental_section := rental_section.At(0).AllByClass("dashed_row") 
				// if err := rental_rental_section.At(0).Find("input[type=radio]").Click(); err != nil{
				// 	log.Fatal(err)
				// }
				if err := rental_rental_section.At(0).Find("input[type=text]").Fill("12500"); err != nil{
					log.Fatal("radio",err)
				}

			rental_daily_rental_price := rental_section.At(1).AllByClass("dashed_row") 
				if err := rental_daily_rental_price.At(1).Find("input[type=radio]").Click(); err != nil {
					log.Fatal(err)
				}

			rental_deposit := rental_section.At(2).AllByClass("dashed_row") 
				if err := rental_deposit.At(0).Find("input[type=radio]").Click(); err != nil {
					log.Fatal(err)
				}
				if err := rental_deposit.At(0).FindByID("rental_deposit_month").Select("3 เดือน"); err != nil{
					log.Fatal("radio",err)
				}
			
			rental_advance_fee := rental_section.At(3).AllByClass("dashed_row") 
				if err := rental_advance_fee.At(0).Find("input[type=radio]").Click(); err != nil {
					log.Fatal(err)
				}
				if err := rental_advance_fee.At(0).FindByID("rental_advance_fee_month").Select("6 เดือน"); err != nil{
					log.Fatal("radio",err)
				}

				amenity := make(map[int]string)
				amenity[1] = "เครื่องปรับอากาศ"
				amenity[3] = "ทีวี"
				amenity[9] = "เตาปรุงอาหาร"
				amenity[11] = "เครื่องซักผ้า"
				
				//property etc ferniture
				amenity_list := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[4]/div`).AllByClass("amenity_list")
				log.Println(amenity_list.Count())
				for index, _ := range amenity {
					log.Println(amenity_list.At(index))
					if index < 6 {
						if err := amenity_list.At(0).All("a").At(index).Click(); err != nil{
							log.Fatal(err)
						}
					} else {
						index = index - 6
						if err := amenity_list.At(1).All("a").At(index).Click(); err != nil{
							log.Fatal(err)
						}
					}
				}
			
			// detail post
			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[5]/ul[1]/li/trix-editor`).Fill("รายละเอียดของประกาศ ภาษาไทย"); err != nil {
				log.Fatal("รายละเอียดของประกาศ[Thai] ", err)
			}

			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[5]/ul[2]/li/trix-editor`).Fill("Detail post"); err != nil {
				log.Fatal("รายละเอียดของประกาศ[Eng]  ", err)
			}

			// upload image
			picture_upload_section := page.FindByID("picture_upload_section").AllByClass("plupload")
			if err := picture_upload_section.Find("input[type=file]").UploadFile("condo_1.jpg"); err != nil {
				log.Fatal("upload image condo1  ", err)
			}
			if err := picture_upload_section.Find("input[type=file]").UploadFile("condo_2.jpg"); err != nil {
				log.Fatal("upload image condo2  ", err)
			}
			if err := picture_upload_section.Find("input[type=file]").UploadFile("condo_3.jpg"); err != nil {
				log.Fatal("upload image condo3  ", err)
			}


			// contact
			if err := page.FindByID(`condo_listing_contact_person`).Fill("Stopkung"); err != nil {
				log.Fatal("ื่อผู้ประกาศ ", err)
			}
			
			if err := page.FindByName(`condo_listing[phone[0]]`).Fill("0861321523"); err != nil {
				log.Fatal("เบอร์ติดต่อผู้ประกาศ1 ", err)
			}

			if err := page.FindByID(`add_phone`).Click(); err != nil {
				log.Fatal("เพิ่ม เบอร์ติดต่อ ", err)
			}

			if err := page.FindByName(`condo_listing[phone[1]]`).Fill("0841235213"); err != nil {
				log.Fatal("เบอร์ติดต่อผู้ประกาศ2 ", err)
			}

			if err := page.FindByID(`condo_listing_email`).Fill("l3lackmocca@gmail.com"); err != nil {
				log.Fatal("email ติดต่อผู้ประกาศ ", err)
			}
		
			
			// ลงประกาศ
			  // รอรูปอัพโหลด
			time.Sleep(30 * time.Second)

			if err := page.FindByXPath(`//*[@id="form-for-condo-listing"]/div[9]/input`).Click(); err != nil {
				log.Fatal("ลงประกาศไม่สำเร็จ : ", err)
			}
			log.Println("Boost post Renthub Success")

	//   if err := driver.Stop(); err != nil {
		// log.Fatal("Failed to close pages and stop WebDriver:", err)
	//   }
}