package email

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v3"
)

func SendEmail(apiKey string, attatchmentFile []byte, emailAddress string) error {

    client := resend.NewClient(apiKey)

	f, err := os.ReadFile("report.txt")
	if err != nil {
		return err
	}

	attachment := &resend.Attachment{
		Content:  f,
		Filename: "report.txt",
	}

    params := &resend.SendEmailRequest{
        From:    "GitLog <gitlog@resend.dev>",
        To:      []string{emailAddress},
		Attachments: []*resend.Attachment{attachment},
		Html:    "<strong>Work report update</strong>",
        Subject: "Updated work report",
    }

    _, err = client.Emails.Send(params)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

	return nil
}