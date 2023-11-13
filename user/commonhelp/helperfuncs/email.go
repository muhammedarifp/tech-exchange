package helperfuncs

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/db"
	"gopkg.in/gomail.v2"
)

func SendVerificationMail(useremail, otp, unique string) bool {
	cfg := config.GetConfig()
	username := cfg.EMAIL
	pass := cfg.EMAIL_PASSWORD
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EMAIL)
	m.SetHeader("To", useremail)
	m.SetHeader("Subject", "Email verification")

	html_template := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Email Verification</title>
	<style>
		body {
			font-family: 'Roboto', Arial, sans-serif;
			margin: 0;
			padding: 0;
			background-color: #f9f9f9;
		}
		.container {
			width: 80%;
			max-width: 600px;
			margin: 20px auto;
			background-color: #fff;
			padding: 30px;
			border-radius: 8px;
			border: 2px solid #ccc;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			text-align: center;
		}
		.header h1 {
			color: #333;
			font-size: 28px;
			margin: 0;
		}
		.content {
			padding: 20px;
		}
		.content p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		}
		.otp-code {
			color: #007bff;
			padding: 15px;
			font-size: 24px;
			font-family: 'monospace', monospace;
			font-weight: bold;
			border-radius: 5px;
			display: inline-block;
			margin-bottom: 20px;
		}
		.unique-code {
			color: black;
			padding: 15px;
			font-size: 10px;
			font-family: 'monospace', monospace;
		}
		.btn {
			display: inline-block;
			padding: 8px 20px;
			text-decoration: none;
			color: #007bff;
			border: 2px solid #007bff;
			border-radius: 5px;
			font-size: 16px;
			transition: background-color 0.3s;
		}
		.btn:hover {
			background-color: #007bff;
			color: #fff;
		}
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>Email Verification</h1>
		</div>
		<div class="content">
			<p>Thank you for signing up. To complete your registration, please use the following OTP code:</p>
			<div class="otp-code">` + otp + `</div>
			<div class="unique-code">` + unique + `</div>
			<p>Enter this code on the registration page to verify your email.</p>
			<a href="#" class="btn">Verify Email</a>
		</div>
	</div>
</body>
</html>
`
	m.SetBody("text/html", html_template)

	d := gomail.NewDialer("smtp.gmail.com", 587, username, pass)
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}

}

// Verify user enter otp
func VerifyUserOtp(otp, userid string) (bool, error) {
	rdb := db.CreateRedisConnection(1)
	val := rdb.Get(context.Background(), userid)
	db_stored_otp, dberr := val.Result()
	if dberr != nil {
		return false, dberr
	}

	if db_stored_otp != otp {
		return false, errors.New("incorrect otp found")
	}

	return true, nil
}

// Create random number
// Throw rand package
func RandomOtpGenarator() string {
	otp := rand.Intn(90000) + 10000
	return strconv.Itoa(otp)
}
