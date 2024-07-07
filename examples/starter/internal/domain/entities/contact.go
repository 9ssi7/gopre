package entities

type Contact struct {
	Base
	Email   string `json:"email" gorm:"type:varchar(255);"`
	Message string `json:"message"`
	IsSeen  bool   `json:"is_seen" gorm:"type:boolean;default:false"`
}

func NewContact(message string, email string) *Contact {
	return &Contact{
		Message: message,
		Email:   email,
	}
}

func (c *Contact) MarkSeen() {
	c.IsSeen = true
}
