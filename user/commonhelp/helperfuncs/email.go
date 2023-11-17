package helperfuncs

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/muhammedarifp/user/config"
	"gopkg.in/gomail.v2"
)

var ()

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

func SendAccountBannedMail(to string, username string) bool {
	template := `
	<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Account Banned</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }

        .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #ffffff;
            border-radius: 5px;
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
        }

        .header {
            background-color: #e44d26;
            padding: 10px;
            text-align: center;
            color: #ffffff;
            border-radius: 5px 5px 0 0;
        }

        h1 {
            color: #e44d26;
        }

        p {
            color: #333333;
            line-height: 1.5;
        }

        .cta-button {
            display: inline-block;
            padding: 10px 20px;
            background-color: #e44d26;
            color: #ffffff;
            text-decoration: none;
            border-radius: 3px;
            margin-top: 20px;
        }

        .footer {
            margin-top: 20px;
            text-align: center;
            color: #777777;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="header">
            <h1>Account Banned</h1>
        </div>
        <p>Dear ` + username + `,</p>
        <p>We regret to inform you that your account has been banned due to a violation of our terms of service.</p>
        <p>If you believe this is an error or would like to appeal the decision, please contact our support team.</p>
        <a class="cta-button" href="mailto:support@example.com">Contact Support</a>
    </div>
    <div class="footer">
        <p>Thank you for using our service. © 2023 Your Company. All rights reserved.</p>
    </div>
</body>

</html>
	`

	cfg := config.GetConfig()
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Accout")

	m.SetBody("text/html", template)

	d := gomail.NewDialer("smtp.gmail.com", 587, cfg.EMAIL, cfg.EMAIL_PASSWORD)
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("errorrr : ", err.Error())
		return false
	} else {
		return true
	}

}

// Create random number
// Throw rand package
func RandomOtpGenarator() string {
	otp := rand.Intn(90000) + 10000
	return strconv.Itoa(otp)
}
