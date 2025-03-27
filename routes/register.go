package routes

import (
	"context"
	"log"

	"github.com/danfelab/youthcongressnepal/connect"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

    // Render register.html from ./public
    return c.Render("register", fiber.Map{})    
}
func RegisterForm(c *fiber.Ctx) error {
	
	// Get the MongoDB client and members collection from the connect package
	_, members := connect.DB()

	// Capture form values for first name, last name, email, mobile, username, password, and address fields
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	// Permanent address fields
	pDistrict := c.FormValue("p_district")
	pPalika := c.FormValue("p_palika")
	pWada := c.FormValue("p_wada")
	pTole := c.FormValue("p_tole")

	// Temporary address fields
	tDistrict := c.FormValue("t_district")
	tPalika := c.FormValue("t_palika")
	tWada := c.FormValue("t_wada")
	tTole := c.FormValue("t_tole")

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password.")
	}

	// Prepare the user data for insertion, including the hashed password and address fields
	membersData := bson.M{
		"firstname":  firstname,
		"lastname":   lastname,
		"email":      email,
		"mobile":     mobile,
		"username":   username,
		"password":   hashedPassword, // Storing the hashed password
		"permanent_address": bson.M{
			"district": pDistrict,
			"palika":  pPalika,
			"wada":    pWada,
			"tole":    pTole,
		},
		"temporary_address": bson.M{
			"district": tDistrict,
			"palika":  tPalika,
			"wada":    tWada,
			"tole":    tTole,
		},
	}

	// Insert the user data into MongoDB
	_, err = members.InsertOne(context.TODO(), membersData)
	if err != nil {
		log.Printf("error inserting user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user.")
	}
	
	return c.Status(fiber.StatusOK).SendString("User registered successfully.")
}

