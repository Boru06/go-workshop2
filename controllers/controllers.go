package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"workshop2/database"
	 m"workshop2/models"

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

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	
	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

func Getcpn(c *fiber.Ctx) error {
	db := database.DBConn
	var cpn []m.Company

	db.Find(&cpn) //delelete = null
	return c.Status(200).JSON(cpn)
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

// }
