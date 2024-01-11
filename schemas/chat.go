package schemas

type Chat struct {
	ChatID      uint   `gorm:"column:chat_id;primaryKey;autoIncrement" json:"chat_id"`
	Message     string `gorm:"column:message" json:"message"`
	RandomCards string `gorm:"column:randoms_cards" json:"randoms_cards"`
	Feedback    bool   `gorm:"column:feedback" json:"feedback"`
	InputID     uint   `gorm:"column:input_id" json:"input_id"`
	Block       bool   `gorm:"column:block" json:"block"`
}

type ChatInput struct {
	InputID uint   `gorm:"column:input_id;primaryKey;autoIncrement" json:"input_id"`
	Message string `gorm:"column:message" json:"message"`
	Block   bool   `gorm:"column:block" json:"block"`
}
