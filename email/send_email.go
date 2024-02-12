package email

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// IsCCOrBCCSend 用于发送包含抄送或密送收件人的邮件
func (s *SendUtil) IsCCOrBCCSend(isText bool, isCC bool) error {
	if s.From == "" ||
		s.To == "" ||
		s.Subject == "" ||
		s.Text == "" ||
		s.addr == "" ||
		s.host == "" ||
		s.MyInfo == nil ||
		s.MyInfo.Username == "" ||
		s.MyInfo.Password == "" {
		return fmt.Errorf("发送邮件失败, 信息不全")
	}

	var e *email.Email
	if isCC {
		if s.cce == nil {
			s.cce = email.NewEmail()
		}
		e = s.cce
	} else {
		if s.bcce == nil {
			s.bcce = email.NewEmail()
		}
		e = s.bcce
	}

	// 添加抄送或密送收件人
	e.From = s.From
	e.To = []string{s.To}
	if isCC {
		e.Cc = s.ToCC // 设置抄送收件人
	} else {
		e.Bcc = s.ToBCC // 设置密送收件人
	}
	e.Subject = s.Subject

	// 设置邮件内容（文本或 HTML）
	if isText {
		e.Text = []byte(s.Text)
	} else {
		e.HTML = []byte(s.Text)
	}

	// 添加附件（如果有）
	if s.Attach != nil {
		_, err := e.Attach(bytes.NewReader(s.Attach.fileBytes), s.Attach.fileName, s.Attach.fileType)
		if err != nil {
			return fmt.Errorf("发送邮件失败: %v", err)
		}
	}

	// 发送邮件
	err := e.Send(s.addr, smtp.PlainAuth(s.MyInfo.Identity, s.MyInfo.Username, s.MyInfo.Password, s.host))
	if err != nil {
		return fmt.Errorf("发送邮件失败: %v", err)
	}

	return nil
}

// SimpleSend 用于发送简单的邮件（不含抄送或密送收件人）
func (s *SendUtil) SimpleSend(isText bool) error {

	if s.From == "" ||
		s.To == "" ||
		s.Subject == "" ||
		s.Text == "" ||
		s.addr == "" ||
		s.host == "" ||
		s.MyInfo == nil ||
		s.MyInfo.Username == "" ||
		s.MyInfo.Password == "" {
		return fmt.Errorf("发送邮件失败, 信息不全")
	}

	var e *email.Email
	if s.se == nil {
		s.se = email.NewEmail()
	}

	e = s.se
	// 设置发件人、收件人和主题
	e.From = s.From
	e.To = []string{s.To}
	e.Subject = s.Subject

	// 设置邮件内容（文本或 HTML）
	if isText {
		e.Text = []byte(s.Text)
	} else {
		e.HTML = []byte(s.Text)
	}

	// 添加附件（如果有）
	if s.Attach != nil {
		_, err := e.Attach(bytes.NewReader(s.Attach.fileBytes), s.Attach.fileName, s.Attach.fileType)
		if err != nil {
			return fmt.Errorf("发送邮件失败: %v", err)
		}
	}

	// 发送邮件
	err := e.Send(s.addr, smtp.PlainAuth(s.MyInfo.Identity, s.MyInfo.Username, s.MyInfo.Password, s.host))
	if err != nil {
		return fmt.Errorf("发送邮件失败: %v", err)
	}
	return nil
}
