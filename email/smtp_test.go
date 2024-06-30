package email

import "testing"

func TestSend(t *testing.T) {
	smtpDialer := NewSMTPDialer("smtp.qiye.aliyun.com", 465, "name@example.com", "password")
	err := smtpDialer.Send(SendOption{
		FromAddr: "name@example.com",
		FromName: "From",
		ToAddr:   "to@example.com",
		ToName:   "To",
		CcAddr:   "cc@example.com",
		CcName:   "Cc",
		Subject:  "Hello",
		Body:     "This is a test email.",
		Attach:   "/tmp/1.txt",
	})
	if err != nil {
		t.Error("send failed", err)
	}
}
