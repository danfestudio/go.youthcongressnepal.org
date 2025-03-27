package routes

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/danfelab/youthcongressnepal/connect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var store = session.New()

func Register(c *fiber.Ctx) error {
    submitted := c.Query("submitted") == "true"
    return c.Render("register", fiber.Map{
        "otpForm": submitted,
    })
}

// GenerateOTP generates a random 5-digit OTP as a string
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	otp := rand.Intn(90000) + 10000  // Generate a number between 10000 and 99999
	return fmt.Sprintf("%d", otp)    // Convert the OTP to a string
}


func RegisterForm(c *fiber.Ctx) error {

    // Create a session object
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

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

    // Check if this is Form 1 submission (based on presence of registration fields)
    if firstname != "" || lastname != "" || email != "" || mobile != "" || username != "" || password != "" {
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
            "password":   hashedPassword,
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
        
        // Generate a random 5-digit OTP
        otp := GenerateOTP()

        // Store OTP in session
        sess.Set("otp", otp)
        err = sess.Save()
        if err != nil {
            log.Println("Error saving session:", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error saving OTP to session.")
        }
            
        if mobile != "" {
            apiKey := "8EB4212649769CADB2CE340DD3FB2026"
            campaign := "youthcongressnepal"
            routeID := "SI_Alert"
            msg := fmt.Sprintf("Thank you for registering with Youth Congress Nepal. Your OTP is: %s", otp)

            data := url.Values{
                "key":      {apiKey},
                "campaign": {campaign},
                "routeid":  {routeID},
                "contacts": {mobile},
                "msg":      {msg},
            }

            resp, err := http.PostForm("https://user.birasms.com/api/smsapi", data)
            if err != nil {
                log.Printf("Error sending SMS: %v", err)
            } else {
                defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Printf("Error reading SMS response: %v", err)
                } else {
                    log.Printf("SMS response: %s", string(body))
                }
            }
        }

        // Redirect to /register with submitted=true to show Form 2 (CHANGED)
        return c.Redirect("/register?submitted=true")
    }

        // Handle Form 2 submission (Update OTP)
	// Handle Form 2 submission (Update OTP)
	if c.FormValue("otp") != "" {
	    enteredOTP := c.FormValue("otp")

	    // Retrieve OTP from session
	    sessionOTP := sess.Get("otp").(string)

	    // Compare the entered OTP with the OTP stored in session (both are strings)
	    if enteredOTP == sessionOTP {
	        log.Printf("OTP match successful. Proceeding with registration.")
	        return c.SendString("OTP verified successfully. Proceeding with registration.")
	    } else {
	        log.Printf("OTP mismatch. Please try again.")
	        return c.SendString("Invalid OTP. Please try again.")
	    }
	}

	// If OTP is not provided, handle it accordingly
	return c.SendString("Please enter the OTP to proceed.")       
   
}