package main

import (
	"fmt"
)

func main() {

	// emailBuilder := new(EmailBuilder)
	emailBuilder := NewEmailBuilder()

	mail := emailBuilder.From("foo@bar.com").
		To("bar@baz.com").
		Subject("Meeting").
		Body("Hello, do you want to meet?").
		Build()

	fmt.Println(mail)
}

type Email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email Email
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{}
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func (b *EmailBuilder) Build() Email {
	return b.email
}
