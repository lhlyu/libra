package common

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"html/template"
	"log"
	"time"
)

type yuEmail struct {
	email *gomail.Message
}

type MessageContent struct {
	FromUserName string
	TimeDate     string
	Content      string
}

func NewMessageContent(fromUserName, content string) *MessageContent {
	return &MessageContent{
		FromUserName: fromUserName,
		Content:      content,
		TimeDate:     time.Now().Format("2006-01-02 15:04:05"),
	}
}

func NewYuEmail(email *gomail.Message) *yuEmail {
	return &yuEmail{email}
}

// 邮箱模板
const _EMAIL_TEMPLATE = `
<html>
    <head>
        <meta charset="utf-8">
    </head>
    <body>
        <h4>有新的评论！</h4>
        </br>
        <div>
            <span>{{.FromUserName}}  |   {{.TimeDate}}</span></br>
            <p>{{.Content}}</p>
        </div>
    </body>
</html>
`

// 错误模板
const _EMAIL_ERROR_TEMPLATE = `
<html>
    <head>
        <meta charset="utf-8">
    </head>
    <body>
        <h4>{{.TimeDate}} | 异常报告</h4>
        </br>
        <div>
            {{range .Errs}}
            <code><pre>{{ . | printf "%+v"}}</pre></code><br>
            {{end}}
        </div>
    </body>
</html>
`

func (y *yuEmail) Send(msg *MessageContent) {
	if y.email == nil {
		log.Printf("有新评论:%+v\n", msg)
		return
	}
	t := template.Must(template.New("email_template.html").Parse(_EMAIL_TEMPLATE))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, msg)
	if err != nil {
		log.Println("send email err = ", err)
		return
	}
	y.email.SetBody("text/html", buf.String())
	d := gomail.NewDialer(Cfg.GetString("email.host"), Cfg.GetInt("email.port"), Cfg.GetString("email.user"), Cfg.GetString("email.password"))
	if err := d.DialAndSend(y.email); err != nil {
		log.Println("email send failure,err = ", err)
		return
	}
	log.Println("send success")
}
