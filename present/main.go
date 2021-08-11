package present

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html"
	jwt "github.com/golang-jwt/jwt"
	"os"
	"qrcode/access/constant"
	"qrcode/control"
	"time"
)

type ContextApi struct {
	fiber.Ctx
	apiControl *control.APIControl
}

func APICreate(ctrl *control.APIControl) {
	engine := html.New("./views", ".html")
	engine.Layout("embed")
	engine.Reload(true)
	engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stderr,
	}))
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                              // todo รอคีย์ domain จะมาจาก env
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH", //
	}))

	app.Use(func(context *fiber.Ctx) error {
		context.Locals(constant.LocalsKeyControl, ctrl)
		return context.Next()
	})

	api := app.Group("/api")
	api.Post("login", login)
	api.Post("admin", admin)

	//qr := app.Group("/qr")
	//qr.Get(":id", getByIdTeamPage)
	//qr.Get("getAllLogTeamPage/:id", getAllLogTeamPage)
	//qr.Get("test", proxy.Forward("http://localhost:12000"))

	// -- Todo Owner
	owner := app.Group("/owner")
	owner.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(constant.SecretKey),
		SuccessHandler: func(context *fiber.Ctx) error {
			user := context.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			var userRole = claims["role"]
			if userRole == string(constant.Owner) {
				return context.Next()
			} else {
				return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}
			return context.Next()
		},
		ErrorHandler: AuthError,
		AuthScheme:   "Bearer",
	}))
	// -- API Owner
	owner.Get("getAccount", getAccount)
	owner.Post("register_operator", registerOperator)
	//owner.Get("getAllTeamPage", getAllTeamPage)

	// -- Todo Admin
	admin := app.Group("/admin")
	admin.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(constant.SecretKey),
		SuccessHandler: func(context *fiber.Ctx) error {
			user := context.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			var userRole = claims["role"]
			if userRole == string(constant.Admin) {
				return context.Next()
			} else {
				return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}
			return context.Next()
		},
		ErrorHandler: AuthError,
		AuthScheme:   "Bearer",
	}))

	// -- Account
	admin.Post("register_owner", registerOwner)
	admin.Post("register_operator", registerOperator)
	admin.Get("getAccount", getAccount)
	admin.Get("getAllAccountOwner", getAllAccountOwner)
	admin.Get("getSubOwner/:id", getSubOwner) // todo ดูข้อมูลทั่งหมดของ Operator ById Owner
	admin.Get("getAllAccountOperator", getAllAccountOperator)
	admin.Get("getOwnerByIdOps/:id",getOwnerByIdOps) // todo ดูข้อมูล Owner ById Ops ยังต้องแก้ SQL ยังไม่ได้ join
	admin.Get("getAccountById/:id", getAccountById)
	admin.Put("updateProfile/:id", updateProfile)
	admin.Put("changePassword/:id", changePassword)
	admin.Delete("deleteAccount/:id", deleteAccount)

	// -- TeamPage
	admin.Get("getTemplate", getTemplate)

	// -- createQrCode
	admin.Post("createQrCode", createQrCode)
	admin.Post("genQrCodeToFileZipByTemplateName", genQrCodeToFileZipByTemplateName)
	admin.Post("genQrCodeToFileZipByQrCodeId", genQrCodeToFileZipByQrCodeId)
	admin.Get("getQrCode/:id", getQrCodeById) // Id >>> OwnerId
	admin.Get("getQrCodeFile/:name", genQrCodeByName)
	admin.Delete("delQrCode", deleteQrCode) // todo ลบ QrCode

	//admin.Get("getAllTeamPage", getAllTeamPage) // todo ส่วนของ Owner Admin สามารถสร้างเพื่อทดสอบได้
	//admin.Get("getAllTeamPageById/:id", getAllTeamPageById)
	//admin.Post("insertTeamPage", insertTeamPage)
	//admin.Put("updateTeamPage/:id", updateTeamPage)
	//admin.Delete("deleteTeamPage/:id", deleteTeamPage)

	// -- Type TeamPage
	//admin.Get("teamPageType",)
	//admin.Post("teamPageType",)
	//admin.Put("teamPageType",)
	//admin.Delete("teamPageType",)

	// -- QrCode

	// -- File
	admin.Post("upload_file", uploadFile)
	admin.Get("get_url_file", getUrlFile)

	_ = app.Listen(":8000")

}

func validateStruct(dataStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			return errors.New(fmt.Sprintf("%s: %s", err.StructField(), err.Tag()))
		}
	} else {
		return nil
	}
	return err
}
