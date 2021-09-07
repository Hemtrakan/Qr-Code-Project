package constant

const (
	//URLCreate สร้าง Rich Menu
	URLCreate string = "https://api.line.me/v2/bot/richmenu"
	//UploadImage อัพรูปลงใน Rich Menu ต่อ Url ด้วย URLRichMenuID
	UploadImage string = "https://api-data.line.me/v2/bot/richmenu/"
	//SetDefaultRichMenu ตั้งต่าเริ่มต้น RichMenuID ต่อ URLRichMenuID
	SetDefaultRichMenu string = "https://api.line.me/v2/bot/user/all/richmenu/"
	//LinkRichMenuToUser "https://api.line.me/v2/bot/user/{userId}/richmenu/{richMenuId}"
	LinkRichMenuToUser string = "https://api.line.me/v2/bot/user/{userId}/richmenu/{richMenuId}"
	//GetDefaultRichMenuList ดึงค่า RichMenuList มาแสดง
	GetDefaultRichMenuList string = "https://api.line.me/v2/bot/richmenu/list"
	//GetDefaultRichMenuID ดึง ค่าเริ่มต้น RichMenuID ที่ใช้งานอยู่
	GetDefaultRichMenuID string = "https://api.line.me/v2/bot/user/all/richmenu"
	//CancelDefaultRichMenu ยกเลิกค่าเริ่มต้น RichMenu
	CancelDefaultRichMenu string = "https://api.line.me/v2/bot/user/all/richmenu"
	//DeleteRichMenu ลบ RichMenu by Id
	DeleteRichMenu string = "https://api.line.me/v2/bot/richmenu/"
)
