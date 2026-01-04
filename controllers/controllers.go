package controllers

import (
	// "log"
	"regexp"
	"strconv"
	m "workshop2/models"

	// "strings"

	// "myapi/database"
	// m "myapi/models"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Factorial(c *fiber.Ctx) error {
	num, err := strconv.Atoi(c.Params("fact"))
	if err != nil || num < 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid number",
		})
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i

	}

	return c.JSON(result)
}

func TaxID(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")
	if taxID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "tax_id is required",
		})
	}

	var ascii []int
	for _, ch := range taxID {
		ascii = append(ascii, int(ch))
	}

	return c.JSON(fiber.Map{
		"tax_id": taxID,
		"ascii":  ascii,
	})
}

func Register(c *fiber.Ctx) error {
	var req m.Register

	// แปลง JSON จาก body → struct
	if err := c.BodyParser(&req); err != nil {
		// ถ้า body ไม่ใช่ JSON หรือรูปแบบผิด
		return c.Status(400).JSON(fiber.Map{
			"error": "รูปแบบข้อมูลไม่ถูกต้อง",
		})
	}

	emailRegex := `^[\w\.-]+@[\w\.-]+\.\w+$` // ตรวจรูปแบบอีเมล
	userRegex := `^[a-zA-Z0-9_]+$`           // ชื่อผู้ใช้ a-z A-Z 0-9 _
	passRegex := `^.{6,20}$`                 // รหัสผ่าน 6-20 ตัวอักษร
	lineRegex := `^[a-zA-Z0-9._-]+$`         // Line ID
	telRegex := `^[0-9]{9,10}$`              // เบอร์โทร 9-10 หลัก
	websiteRegex := `^[a-z0-9-]{2,30}$`      // ชื่อเว็บไซต์

	// =========================
	// ตรวจสอบข้อมูลทีละช่อง
	// =========================

	// ตรวจอีเมล
	if match, _ := regexp.MatchString(emailRegex, req.Email); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "อีเมลไม่ถูกต้อง",
		})
	}

	// ตรวจชื่อผู้ใช้
	if match, _ := regexp.MatchString(userRegex, req.Username); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "ชื่อผู้ใช้ไม่ถูกต้อง",
		})
	}

	// ตรวจรหัสผ่าน
	if match, _ := regexp.MatchString(passRegex, req.Password); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "รหัสผ่านต้องยาว 6-20 ตัวอักษร",
		})
	}

	// ตรวจ Line ID
	if match, _ := regexp.MatchString(lineRegex, req.LineID); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "Line ID ไม่ถูกต้อง",
		})
	}

	// ตรวจเบอร์โทรศัพท์
	if match, _ := regexp.MatchString(telRegex, req.Tel); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "เบอร์โทรศัพท์ไม่ถูกต้อง",
		})
	}

	// ตรวจว่ามีการเลือกประเภทธุรกิจหรือไม่
	if req.Business == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "กรุณาเลือกประเภทธุรกิจ",
		})
	}

	// ตรวจชื่อเว็บไซต์
	if match, _ := regexp.MatchString(websiteRegex, req.Website); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "ชื่อเว็บไซต์ไม่ถูกต้อง",
		})
	}

	// =========================
	// กรณีข้อมูลถูกต้องทั้งหมด
	// =========================
	return c.JSON(fiber.Map{
		"message": "สมัครสมาชิกสำเร็จ",
	})
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("hello, World")
}
