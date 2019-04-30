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
		if err := page.Navigate(`https://www.renthub.in.th/apartments/new`); err != nil {
			log.Fatal("create apartment", err)
		}		

		// ชื่อที่พัก 

		page.Find(`#apartment_name`).Fill("ทดลองสร้างที่ apartment");

		page.Find(`#temp_eng_name`).Fill("Testing to create apartment");

		// service apartment

		page.Find(`#form-for-apartment > div:nth-child(2) > ul > li:nth-child(3) > div.indent.clearfix.service-aprtment > a`).Click();

		// contact

		page.Find(`#apartment_contact_person`).Fill("stopkung");

		page.Find(`#apartment_email`).Fill("stopkung@hotmail.com");

		page.Find(`#form-for-apartment > div:nth-child(2) > ul > li:nth-child(5) > div > div:nth-child(1) > input[type="text"]`).Fill("0876836191")

		page.Find(`#add_phone`).Click();

		page.Find(`#form-for-apartment > div:nth-child(2) > ul > li:nth-child(5) > div > div:nth-child(2) > input[type="text"]`).Fill("0876836199")

		page.Find(`#apartment_line_id`).Fill("stopline")

		// ที่อยู่

		page.Find(`#apartment_address`).Fill("622/208")

		page.Find(`#apartment_road`).Fill("บางขุนเทียน-ชายทะเล")

		page.Find(`#apartment_street`).Fill("เทียนทะเล 20")

		time.Sleep(2 * time.Second)

		page.Find(`#apartment_province_code`).Select("กรุงเทพมหานคร")

		time.Sleep(2 * time.Second)

		page.Find(`#apartment_district_code`).Select("บางขุนเทียน")

		time.Sleep(2 * time.Second)

		page.Find(`#apartment_subdistrict_code`).Select("ท่าข้าม")

		page.Find(`#apartment_postcode`).Fill("10140")

		page.Find(`#maplookup`).Click();

		// สิ่งอำนวยความสะดวก ภายในห้องพัก

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(2) > a:nth-child(2)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(2) > a:nth-child(4)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(2) > a:nth-child(6)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(2) > a:nth-child(8)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(2) > a:nth-child(11)`).Click();
			
		//	สิ่งอำนวยความสะดวก ภายในอาคารที่พัก

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(2)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(3)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(5)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(7)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(9)`).Click();

		page.Find(`#form-for-apartment > div:nth-child(3) > div > div:nth-child(3) > a:nth-child(10)`).Click();

		// ระบุประเภทห้องพัก

		page.Find(`#room-table > tbody > tr:nth-child(1) > td.name > input`).Fill("Superior")

		page.Find(`#room_type_0`).Select("3 ห้องนอน")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td.size > input[type="text"]`).Fill("18")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(5) > a`).Click()

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(5) > div > input[type="text"]:nth-child(2)`).Fill("10000")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(5) > div > input[type="text"]:nth-child(4)`).Fill("25000")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(6) > a`).Click()

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(6) > div > input[type="text"]:nth-child(2)`).Fill("300")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td:nth-child(6) > div > input[type="text"]:nth-child(4)`).Fill("800")

		page.Find(`#room-table > tbody > tr:nth-child(1) > td.available > a`).Click()

		// add another room

		page.Find(`#add_room`).Click()

		time.Sleep(2* time.Second)

		page.Find(`#room-table > tbody > tr:nth-child(2) > td.name > input[type="text"]`).Fill("Luxury")

		page.Find(`#room_type_1`).Select("สตูดิโอ")

		page.Find(`#room-table > tbody > tr:nth-child(2) > td.size > input[type="text"]`).Fill("70")

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(5) > a`).Click()

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(5) > div > input[type="text"]:nth-child(2)`).Fill("100000")

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(5) > div > input[type="text"]:nth-child(4)`).Fill("250000")

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(6) > a`).Click()

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(6) > div > input[type="text"]:nth-child(2)`).Fill("3000")

		page.Find(`#room-table > tbody > tr:nth-child(2) > td:nth-child(6) > div > input[type="text"]:nth-child(4)`).Fill("8000")

		//  ค่าใช้จ่าย
			// ค่าน้ำ
		page.Find(`#fee_water_price_type_2`).Click();

		page.Find(`#fee_water_price_monthly_per_person`).Fill("2000");

		page.Find(`#fee_water_price_monthly_per_person_remark`).Fill("80");

		page.Find(`#fee_water_price_per_person_exceed`).Fill("200");

		//  ค่าไฟฟ้า

		page.Find(`fee_electric_price_type_1`).Click()

		page.Find(`#fee_electric_price`).Fill("200");

		page.Find(`#fee_electric_price_minimum`).Fill("0");

		// ค่าบริการ

		page.Find(`#fee_service_fee_type_2`).Click();

		// ค่ามัดจำ / ประกัน

		page.Find(`#fee_deposit_type_0`).Click();

		// จ่ายล่วงหน้า

		page.Find(`#fee_advance_fee_month`).Select("1 เดือน");

		// ค่า โทรศักท์

		page.Find(`#fee_phone_price_type_3`).Click();

		// ค่า internet

		page.Find(`#fee_internet_price_bath`).Fill("699");

		// รายละเอียด thai
		detail_th := "นี่คือรายละเอียด Apartmemt ภาษาไทย";
		page.RunScript(`
				$("#form-for-apartment > div:nth-child(6) > ul > li:nth-child(1) > iframe").contents()[0].body.innerHTML = detail_th
			`,
			map[string]interface{}{ "detail_th": detail_th},
			nil);

		// รายละเอียด Eng

		detail_en := "detail in English";

		page.RunScript(`
			$("#form-for-apartment > div:nth-child(6) > ul > li:nth-child(2) > iframe").contents()[0].body.innerHTML = detail_en
			`,
			map[string]interface{}{ "detail_en": detail_en},
		nil);

		// upload pic

		page.Find(`#picture_upload_section > div.plupload.html5 > input[type="file"]`).UploadFile("../condo_1.jpg");

		page.Find(`#picture_upload_section > div.plupload.html5 > input[type="file"]`).UploadFile("../condo_2.jpg");

		page.Find(`#picture_upload_section > div.plupload.html5 > input[type="file"]`).UploadFile("../condo_3.jpg");

		// promotion

		page.Find(`#promotion > div > li:nth-child(4) > div > textarea`).Fill("รายละเอียด โปรโมชั่น")

		date_start := "2019-04-20";
		date_start_script := `$("#apartment_promotion_start").attr("value", date)`;

		page.RunScript(date_start_script, map[string]interface{}{"date": date_start},nil)

		date_end := "2019-04-20";
		date_end_script := `$("#apartment_promotion_end").attr("value", date)`;

		page.RunScript(date_end_script, map[string]interface{}{"date": date_end},nil)

		// wait for upload file

		time.Sleep(30 * time.Second)

		page.Find(`#form-for-apartment > div.actions > input`).Click()

}