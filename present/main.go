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
	"github.com/golang-jwt/jwt"
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
	app := fiber.New(fiber.Config{})

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
	api.Post("admin", admin) // todo สำหรับ สมัคร admin เท่านั้น
	api.Get("testqrcode",Test)

	qr := app.Group("/qr")
	qr.Post("/:id", getDataQrCode)                      //  Id >>> QrCodeUUId
	qr.Post("getDataQrCodeJson/:id", getDataQrCodeJson) //  Id >>> QrCodeUUId

	// -- Todo Owner
	owner := app.Group("/owner")
	owner.Post("login", login)
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
	// -- API Owner Account
	owner.Get("getAccount", getAccount)
	owner.Post("register_operator", registerOperatorOwner)

	owner.Get("getOperator", getOperator) // todo ดูข้อมูลทั่งหมดของ Operator ById Owner
	owner.Get("getOperatorById/:id", getOperatorById)
	owner.Put("updateProfile/:id", updateProfile)
	owner.Put("changePasswordOwner", ChangePasswordOwner)
	owner.Put("changePasswordOperator", ChangePasswordOperator)
	owner.Delete("deleteAccount/:id", deleteAccountOperator)

	owner.Get("getQrCode", getQrCodeOwnerById) // Id >>> OwnerId

	// -- Todo Admin
	admin := app.Group("/admin")
	admin.Post("login", LoginAdmin)
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
	admin.Get("getOwnerByIdOps/:id", getOwnerByIdOps) // todo ดูข้อมูล Owner ById Ops ยังต้องแก้ SQL ยังไม่ได้ join
	admin.Get("getAccountById/:id", getAccountById)
	admin.Put("updateProfile/:id", updateProfile)
	admin.Put("changePassword/:id", changePassword)
	admin.Delete("deleteAccount/:id", deleteAccount)

	// -- createQrCode
	admin.Get("getDateQrCodeById/:id", getDataQrCode)
	admin.Post("createQrCode", createQrCode)
	admin.Post("genQrCodeToFileZipByTemplateName", genQrCodeToFileZipByTemplateName)
	admin.Post("genQrCodeToFileZipByQrCodeId", genQrCodeToFileZipByQrCodeId)
	admin.Get("getAllQrCodeByOwnerId/:id", getQrCodeById) // Id >>> OwnerId
	admin.Get("getAllQrCode",getAllQrCode)
	admin.Post("insertDataQrCode",insertDataQrCode)
	admin.Delete("delQrCode", deleteQrCode) // todo ลบ QrCode
	//admin.Get("getQrCodeFile/:name", genQrCodeByName)

	// -- TeamPage
	admin.Get("getTemplate", getTemplate)

	// -- TeamPage Qr Computer
	com := admin.Group("/computer")
	com.Post("/:id", Insert) // สำหรับเพิ่มครั้งแรก
	com.Put("/:id", UpData)  // สำหรับแก้ไขข้อมูล

	_ = app.Listen(":8000")

}

func ValidateStruct(dataStruct interface{}) error {
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
