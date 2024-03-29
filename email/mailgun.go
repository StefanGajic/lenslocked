// package email

// import (
// 	"fmt"

// 	mailgun "gopkg.in/mailgun/mailgun-go.v3"
// )

// const (
// 	welcomeSubject = "Welcome to LensGopher.com!"
// )

// const welcomeText = `Hi there!

// Welcome to LensGopher.com! We really hope you enjoy using
// our application!

// Best
// `

// const welcomeHTML = `Hi there!<br/>
// <br/>
// Welcome to
// <a href="https://www.lensgopher.com">LensLocked.com</a>! We really hope you enjoy using our application!<br/>
// <br/>
// Best,<br/>
// Jon
// `

// func WithMailgun(domain, apiKey, publicKey string) ClientConfig {
// 	return func(c *Client) {
// 		mg := mailgun.NewMailgun(domain, apiKey, publicKey)
// 		c.mg = mg
// 	}
// }

// func WithSender(name, email string) ClientConfig {
// 	return func(c *Client) {
// 		c.from = buildEmail(name, email)
// 	}
// }

// type ClientConfig func(*Client)

// func NewClient(opts ...ClientConfig) *Client {
// 	client := Client{
// 		// Set a default from email address...
// 		from: "golang@gmx.com",
// 	}
// 	for _, opt := range opts {
// 		opt(&client)
// 	}
// 	return &client
// }

// type Client struct {
// 	from string
// 	mg   mailgun.Mailgun
// }

// func (c *Client) Welcome(toName, toEmail string) error {
// 	message := mailgun.NewMessage(c.from, welcomeSubject, welcomeText, buildEmail(toName, toEmail))
// 	message.SetHtml(welcomeHTML)
// 	_, _, err := c.mg.Send(message)
// 	return err
// }

// func buildEmail(name, email string) string {
// 	if name == "" {
// 		return email
// 	}
// 	return fmt.Sprintf("%s <%s>", name, email)
// }
