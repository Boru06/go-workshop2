package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"workshop2/database"
	m "workshop2/models"

	"github.com/gofiber/fiber/v2"
)

func Factorial(c *fiber.Ctx) error {

	// แปลง string เป็น int
	num, err := strconv.Atoi(c.Params("fact"))
	// เช็คว่าเป็นค่าว่าง หรือ ว่ามีค่าน้อยว่า0หรือไม่
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
	// ดึงค่ามาจาก query param
	taxID := c.Query("tax_id")
	// ถ้าQuery param ที่รับเข้ามามีค่าว่าง ก็จะแจ้งerror กลับไป
	if taxID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "tax_id is required",
		})
	}

	// เก็บค่า ascii
	var ascii []int
	// ลูป แปลงตัวอักษรแต่ละตัวเป็นเลขascii
	for _, v := range taxID {
		ascii = append(ascii, int(v))
	}
	// return ค่ากลับไปเป็น json
	return c.JSON(fiber.Map{
		"ascii": ascii,
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

	// ตรวจอีเมล
	email := `^[\w\.-]+@[\w\.-]+\.\w+$`
	if match, _ := regexp.MatchString(email, req.Email); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "อีเมลไม่ถูกต้อง",
		})
	}

	// ตรวจชื่อผู้ใช้
	user := `^[a-zA-Z0-9_]+$`
	if match, _ := regexp.MatchString(user, req.Username); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "ชื่อผู้ใช้ไม่ถูกต้อง",
		})
	}

	// ตรวจรหัสผ่าน
	pass := `^.{6,20}$`
	if match, _ := regexp.MatchString(pass, req.Password); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "รหัสผ่านต้องยาว 6-20 ตัวอักษร",
		})
	}

	// ตรวจ Line ID
	line := `^[a-zA-Z0-9._-]+$`
	if match, _ := regexp.MatchString(line, req.LineID); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "Line ID ไม่ถูกต้อง",
		})
	}

	// ตรวจเบอร์โทรศัพท์
	tel := `^[0-9]{9,10}$`
	if match, _ := regexp.MatchString(tel, req.Tel); !match {
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
	website := `^[a-z0-9-]{2,30}$`
	if match, _ := regexp.MatchString(website, req.Website); !match {
		return c.Status(400).JSON(fiber.Map{
			"error": "ชื่อเว็บไซต์ไม่ถูกต้อง",
		})
	}

	//ข้อมูลถูกต้องทั้งหมด
	return c.JSON(fiber.Map{
		"message": "สมัครสมาชิกสำเร็จ",
	})
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("hello, World")
}

// Dog

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetRemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	if err := db.Unscoped().
		Where("deleted_at is not null").
		Find(&dogs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(dogs)
}

func GetDogsRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	min := c.QueryInt("min")
	max := c.QueryInt("max")

	if min == 0 || max == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "please provide min and max query parameters",
		})
	}

	if err := db.
		Where("dog_id > ? AND dog_id < ?", min, max).
		Find(&dogs).Error; err != nil {

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// ❌ ไม่พบข้อมูลในฐานข้อมูล
	if len(dogs) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "no dogs found in this range",
		})
	}

	return c.Status(200).JSON(dogs)
}
func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes

	// ตัวแปรนับผลรวมแต่ละสี
	redCount := 0
	greenCount := 0
	pinkCount := 0
	noColorCount := 0

	for _, v := range dogs {

		typeStr := ""

		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			redCount++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			greenCount++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			pinkCount++
		} else {
			typeStr = "no color"
			noColorCount++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}

		dataResults = append(dataResults, d)
	}

	// response
	r := m.ResultData{
		Count:   len(dataResults), // ผลรวมทั้งหมด
		Name:    "golang-test",
		Data:    dataResults,
		Red:     redCount,
		Green:   greenCount,
		Pink:    pinkCount,
		NoColor: noColorCount,
	}

	return c.Status(200).JSON(r)
}

// company function
func GetcpnAll(c *fiber.Ctx) error {
	db := database.DBConn
	var cpn []m.Company

	db.Find(&cpn) //delelete = null
	return c.Status(200).JSON(cpn)
}

func Getcpn(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Company

	result := db.Find(&dog, "id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func Addcpn(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var cpn m.Company

	if err := c.BodyParser(&cpn); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&cpn)
	return c.Status(201).JSON(cpn)
}

func Updatecpn(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Company
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func Removecpn(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Company

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// }
