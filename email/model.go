package email

import (
	"github.com/jordan-wright/email" // 导入发送邮件所需的库
)

// EType 表示邮件类型
type EType int

// AddrType 表示地址类型
type AddrType int

// QQ 邮件类型常量
const (
	QQ EType = iota // QQ 邮箱类型
)

// SendUtil 封装了发送邮件的相关信息和操作
type SendUtil struct {
	From    string       // 发件人邮箱地址
	To      string       // 收件人邮箱地址
	Subject string       // 邮件主题
	Text    string       // 邮件正文
	Attach  *Attach      // 附件信息
	MyInfo  *User        // 收件人信息
	ToCC    []string     // 抄送收件人信息
	ToBCC   []string     // 密送收件人信息
	addr    string       // SMTP 服务器地址
	host    string       // SMTP 服务器主机名
	se      *email.Email // 发送邮件实例
	cce     *email.Email // 抄送邮件实例
	bcce    *email.Email // 密送邮件实例
}

// Attach 表示邮件附件信息
type Attach struct {
	fileBytes []byte // 附件文件内容
	fileName  string // 附件文件名
	fileType  string // 附件文件类型
}

// NewAttach 创建新的附件信息
func NewAttach(fileBytes []byte, fileName string, fileType string) *Attach {
	return &Attach{fileBytes: fileBytes, fileName: fileName, fileType: fileType}
}

// User 表示收件人的信息
type User struct {
	Identity string // 收件人标识
	Username string // 收件人邮箱用户名
	Password string // 收件人邮箱密码
}

// NewUser 创建新的收件人信息
func NewUser(identity string, username string, password string) *User {
	return &User{Identity: identity, Username: username, Password: password}
}

// NewBCCSendUtil 创建新的密送邮件实例
func NewBCCSendUtil(from string, to string, subject string, text string, attach *Attach, myInfo *User, ToBCC []string, t EType) *SendUtil {
	e := &SendUtil{From: from, To: to, Subject: subject, Text: text, Attach: attach, MyInfo: myInfo, ToBCC: ToBCC}
	switchType(t, e)
	return e
}

// NewCCSendUtil 创建新的抄送邮件实例
func NewCCSendUtil(from string, to string, subject string, text string, attach *Attach, myInfo *User, ToCC []string, t EType) *SendUtil {
	e := &SendUtil{From: from, To: to, Subject: subject, Text: text, Attach: attach, MyInfo: myInfo, ToCC: ToCC}
	switchType(t, e)
	return e
}

// NewSimpleSendUtil 创建新的普通邮件实例
func NewSimpleSendUtil(from string, to string, subject string, text string, attach *Attach, myInfo *User, t EType) *SendUtil {
	e := &SendUtil{From: from, To: to, Subject: subject, Text: text, Attach: attach, MyInfo: myInfo}
	switchType(t, e)
	return e
}

// switchType 根据邮件类型切换 SMTP 服务器地址和主机名
func switchType(t EType, e *SendUtil) {
	switch t {
	case QQ:
		e.addr = "smtp.qq.com:25" // Q
		e.host = "smtp.qq.com"
		break
	}
}
