package retirement

import (
	"fmt"
	"retirement53/model"
)

type Notification struct{ Recipient, Subject, Body string }

func ConfirmationNotification(c model.RetirementCase) Notification {
	return Notification{Recipient: c.UserID, Subject: "退休确认", Body: fmt.Sprintf("已确认 %d/%d", c.Confirmations, c.Expected)}
}
func Pending(c model.RetirementCase) bool { return c.Status == "open" && c.Confirmations < c.Expected }
func Progress(c model.RetirementCase) float64 {
	if c.Expected == 0 {
		return 0
	}
	return float64(c.Confirmations) / float64(c.Expected)
}
func StatusLabel(c model.RetirementCase) string {
	switch c.Status {
	case "open":
		return "待审核"
	case "confirmed":
		return "已确认"
	case "closed":
		return "已归档"
	default:
		return "未知"
	}
}
