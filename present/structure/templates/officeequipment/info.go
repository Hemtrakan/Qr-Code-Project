package officeequipment

import "time"

type Info struct {
	//ProductCode รหัสสินค้า
	ProductCode string

	//ProductName ชื่อสินค้า
	ProductName string

	//EmployeeID รหัสพนักงาน
	EmployeeID string

	//ProductType ประเภทสินค้า
	ProductType string

	//Department แผนก
	Department string

	//ProductUser ชื่อผู้ใช้สินค้า
	ProductUser string

	//ProductDetails รายละเอียดสินค้า
	ProductDetails string

	//SerialNumber หมายเลขซีเรียล
	SerialNumber string

	//Note หมายเหตุ
	Note string

	//ProductInsurance ประกันสินค้า
	ProductInsurance time.Time

	//PurchaseDate วันที่ซื้อ
	PurchaseDate time.Time

	//StartUsingTheProduct เริ่มใช้ผลิตภัณฑ์
	StartUsingTheProduct time.Time

	//EndUsingTheProduct สิ้นสุดการใช้ผลิตภัณฑ์
	EndUsingTheProduct time.Time
}