package helper

import (
	"log"
	"net/smtp"
	"strings"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
	Code     string
	UserName string
}

func SendVerificationCode(params Params) error {
	// htmlFile, err := os.ReadFile("api/helper/format.html")
	// if err != nil {
	// 	log.Println("Cannot read html file", err.Error())
	// 	return err
	// }
	// temp, err := template.New("email").Parse(string(htmlFile))
	// if err != nil {
	// 	log.Println("Cannot parse html file")
	// 	return err
	// }

	var Builder strings.Builder
	// err = temp.Execute(&Builder, params)
	// if err != nil {
	// 	log.Println("Cannot execute HTML", err.Error())
	// 	return err
	// }

	message := "From: " + params.From + "\n" + "To: " + params.To + "\n" + "Message: " + "Subject: Email Verification\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		"<html>" +
		"<body>" +
		"<h1>Hi,</h1>" +
		"<p></p>" +
		"<p>Kod: <strong>" + params.Code + "</strong></p>" +
		"</body>" +
		"</html>\r\n" + "\n" + Builder.String()

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("smtp", params.From, params.Password, "smtp.gmail.com"),
		params.From, []string{params.To}, []byte(message),
	)

	if err != nil {
		log.Println("Could not send an email", err.Error())
		return err
	}

	return nil
}
