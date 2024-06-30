package email

import "gopkg.in/gomail.v2"

type SMTPDialer struct {
	dialer *gomail.Dialer
}

type SendOption struct {
	FromAddr string
	FromName string
	ToAddr   string
	ToName   string
	CcAddr   string
	CcName   string
	Subject  string
	Body     string
	Attach   string
}

func NewSMTPDialer(host string, port int, username string, password string) *SMTPDialer {
	return &SMTPDialer{
		dialer: gomail.NewDialer(host, port, username, password),
	}
}

func (s *SMTPDialer) Send(opt SendOption) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", opt.FromAddr, opt.FromName)
	m.SetAddressHeader("To", opt.ToAddr, opt.ToName)
	if opt.CcAddr != "" {
		m.SetAddressHeader("Cc", opt.CcAddr, opt.CcName)
	}
	m.SetHeader("Subject", opt.Subject)
	m.SetBody("text/html", opt.Body)
	if opt.Attach != "" {
		m.Attach(opt.Attach)
	}

	return s.dialer.DialAndSend(m)
}
