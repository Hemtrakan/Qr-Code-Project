package constant

type URL string

const (
	//URLCreate สร้าง Rich Menu
	URLCreate URL = "https://api.line.me/v2/bot/richmenu"
	//UploadImage อัพรูปลงใน Rich Menu ต่อ Url ด้วย URLRichMenuID
	UploadImage URL = "https://api-data.line.me/v2/bot/richmenu/"
	//SetDefaultRichMenu ตั้งต่าเริ่มต้น RichMenuID ต่อ URLRichMenuID
	SetDefaultRichMenu URL = "https://api.line.me/v2/bot/user/all/richmenu/"
	//GetDefaultRichMenuList ดึงค่า RichMenuList มาแสดง
	GetDefaultRichMenuList URL = "https://api.line.me/v2/bot/richmenu/list"
	//GetDefaultRichMenuID ดึง ค่าเริ่มต้น RichMenuID ที่ใช้งานอยู่
	GetDefaultRichMenuID URL = "https://api.line.me/v2/bot/user/all/richmenu"
	//CancelDefaultRichMenu ยกเลิกค่าเริ่มต้น RichMenu
	CancelDefaultRichMenu URL = "https://api.line.me/v2/bot/user/all/richmenu"
	//DeleteRichMenu ลบ RichMenu by Id
	DeleteRichMenu URL = "https://api.line.me/v2/bot/richmenu/"

)

