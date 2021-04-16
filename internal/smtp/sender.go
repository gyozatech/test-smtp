package smtp

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"strings"

	"catapulta.hoplo.com/smtp/internal/email"
)

type sender struct {
	Hostname string
}

func (s *sender) Send(from string, to []string, msg string) error {
	for _, addr := range to {
		_, domain, err := email.Split(addr)
		if err != nil {
			return err
		}

		mxs, err := net.LookupMX(domain)
		if err != nil {
			return err
		}
		if len(mxs) == 0 {
			mxs = []*net.MX{{Host: domain}}
		}

		for _, mx := range mxs {
			if err := s.sendMailToHost(from, to, addr, mx.Host, msg); err != nil {
				fmt.Printf("err: %v\n", err)
				continue
			}
			fmt.Printf("Mail sent to %v\n", addr)
			break
		}
	}

	return nil
}

func (s *sender) sendMailToHost(from string, to []string, addr string, host string, msg string) error {
	c, err := smtp.Dial(host + ":25")
	if err != nil {
		return err
	}

	if err := c.Hello(s.Hostname); err != nil {
		return err
	}

	if ok, _ := c.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{ServerName: host}
		if err := c.StartTLS(tlsConfig); err != nil {
			return err
		}
	}

	if err := c.Mail(from); err != nil {
		return err
	}

	if err := c.Rcpt(addr); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}

	r := strings.NewReader(msg)
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	if err := c.Quit(); err != nil {
		return err
	}
	return nil
}
