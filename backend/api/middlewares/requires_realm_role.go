package middlewares

import (
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/shared/enums"
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/shared/jwt"
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
)

// โค้ดที่ให้มาเป็น middleware ที่ใช้สำหรับตรวจสอบสิทธิ์การเข้าถึงและการอนุญาตในการเข้าถึงเส้นทาง (route) ในแอปพลิเคชันของคุณ โดยโค้ดนี้มีหน้าที่ตรวจสอบว่าผู้ใช้ที่เข้าถึงเส้นทางนั้นมีบทบาท (role) ที่กำหนดหรือไม่ และถ้าไม่มีบทบาทที่กำหนด โค้ดจะส่งคืนข้อความแสดงว่า "role authorization failed" และส่งสถานะการตอบสนอง 401 (Unauthorized) กลับไปยังผู้ใช้
func NewRequiresRealmRole(role string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		claims := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := jwt.NewJwtHelper(claims)
		if !jwtHelper.IsUserInRealmRole(role) {
			return c.Status(fiber.StatusUnauthorized).SendString("role authorization failed")
		}
		return c.Next()
	}
}
