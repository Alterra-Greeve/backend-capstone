package helper

import (
	"backendgreeve/config"
	"log"
	"net/smtp"
	"time"
)

type MailerInterface interface {
	Send(to string, code string) error
}

type Mailer struct {
	config config.SMTPConfig
}

func NewMailer(config config.SMTPConfig) MailerInterface {
	return &Mailer{
		config: config,
	}
}

func (m *Mailer) Send(to string, code string) error {
	from := m.config.Username
	pass := m.config.Password
	date := time.Now().Format("Monday, 2 January 2006")
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Capstone Project - Verification</title>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600&display=swap" rel="stylesheet" />
</head>
<body style="
      margin: 0;
      font-family: 'Poppins', sans-serif;
      background: #ffffff;
      font-size: 14px;
    ">
    <div style="
        max-width: 680px;
        margin: 0 auto;
        padding: 45px 30px 60px;
        background: #f4f7ff;
        background-image: url(https://archisketch-resources.s3.ap-northeast-2.amazonaws.com/vrstyler/1661497957196_595865/email-template-background-banner);
        background-repeat: no-repeat;
        background-size: 800px 452px;
        background-position: top center;
        font-size: 14px;
        color: #434343;
      ">
        <header>
            <table style="width: 100%;">
                <tbody>
                    <tr style="height: 0;">
                        <td>
                        </td>
                        <td style="text-align: right;">
                            <span style="font-size: 16px; line-height: 30px; color: #ffffff;">` + date + `</span>
                        </td>
                    </tr>
                </tbody>
            </table>
        </header>
        <main>
            <div style="
            margin: 0;
            margin-top: 70px;
            padding: 92px 30px 115px;
            background: #ffffff;
            border-radius: 30px;
            text-align: center;
          ">
                <div style="width: 100%; max-width: 489px; margin: 0 auto;">
                    <h1 style="
                margin: 0;
                font-size: 24px;
                font-weight: 500;
                color: #1f1f1f;
              ">
                        Kode OTP Anda
                    </h1>
                    <p style="
                margin: 0;
                margin-top: 17px;
                font-weight: 500;
                letter-spacing: 0.56px;
              ">
                        Terima kasih telah memilih Greeve. Gunakan OTP berikut
                        untuk menyelesaikan prosedur penggantian alamat email Anda. OTP ini
                        berlaku selama
                        <span style="font-weight: 600; color: #1f1f1f;">10 menit</span>.
                        Jangan bagikan kode ini kepada orang lain, termasuk karyawan Greeve.
                    </p>
                    <p style="
                margin: 0;
                margin-top: 60px;
                font-size: 40px;
                font-weight: 600;
                letter-spacing: 18px;
                color: #ba3d4f;
              ">
                        ` + code + `
                    </p>
                </div>
            </div>
            <p style="
            max-width: 400px;
            margin: 0 auto;
            margin-top: 90px;
            text-align: center;
            font-weight: 500;
            color: #8c8c8c;
          ">
                Need help? Ask at
                <a href="mailto:admin@greeve.store"
                    style="color: #499fb6; text-decoration: none;">admin@greeve.store</a>
            </p>
        </main>
        <footer style="
          width: 100%;
          max-width: 490px;
          margin: 20px auto 0;
          text-align: center;
          border-top: 1px solid #e6ebf1;
        ">
            <p style="
            margin: 0;
            margin-top: 40px;
            font-size: 16px;
            font-weight: 600;
            color: #434343;
          ">
                Greeve
            </p>
            <div style="margin: 0; margin-top: 16px;">
                <a href="#" style="display: inline-block;">
                    <img width="36px" alt="Facebook"
                        src="https://archisketch-resources.s3.ap-northeast-2.amazonaws.com/vrstyler/1661502815169_682499/email-template-icon-facebook" />
                </a>
                <a href="#" style="display: inline-block; margin-left: 8px;">
                    <img width="36px" alt="Instagram"
                        src="https://archisketch-resources.s3.ap-northeast-2.amazonaws.com/vrstyler/1661504218208_684135/email-template-icon-instagram" /></a>
                <a href="#" style="display: inline-block; margin-left: 8px;">
                    <img width="36px" alt="Twitter"
                        src="https://archisketch-resources.s3.ap-northeast-2.amazonaws.com/vrstyler/1661503043040_372004/email-template-icon-twitter" />
                </a>
                <a href="#" style="display: inline-block; margin-left: 8px;">
                    <img width="36px" alt="Youtube"
                        src="https://archisketch-resources.s3.ap-northeast-2.amazonaws.com/vrstyler/1661503195931_210869/email-template-icon-youtube" /></a>
            </div>
            <p style="margin: 0; margin-top: 16px; color: #434343;">
                Copyright © 2024 Greeve. All rights reserved.
            </p>
        </footer>
    </div>
</body>
</html>`

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = "Greeve Verification Code"
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	// Membangun header email
	header := ""
	for k, v := range headers {
		header += k + ": " + v + "\r\n"
	}

	msg := header + "\r\n" + body

	err := smtp.SendMail(m.config.Host+":"+m.config.Port,
		smtp.PlainAuth("", from, pass, m.config.Host),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	log.Println("Successfully sent to " + to)
	return nil
}
