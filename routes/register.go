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

	// Capture form values for first name, last name, email, mobile, username, and password
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Hash the password using bcrypt (to store a hashed password, not plain text)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password.")
	}

	// Prepare the user data for insertion, including the hashed password
	membersData := bson.M{
		"firstname": firstname,
		"lastname":  lastname,
		"email":     email,
		"mobile":    mobile,
		"username":  username,
		"password":  hashedPassword, // Storing the hashed password
	}

	// Insert the user data into MongoDB
	_, err = members.InsertOne(context.TODO(), membersData)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user.")
	}

	// Return success response
	return c.SendString("User registered successfully!")
}

