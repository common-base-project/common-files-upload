package email

/*
  @Author : zggong
*/

func SendEmail(textContent PlainTextContent) {
	emailClient := NewMailClient()
	emailClient.SendEmail(MailKindPlainText, textContent)
}
