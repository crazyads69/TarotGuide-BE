package schemas

// Card là mô hình cho bảng "cards"
type Card struct {
	CardID   uint   `json:"card_id" gorm:"column:card_id;primaryKey;autoIncrement" binding:"required"`
	CardName string `json:"card_name" gorm:"column:card_name;"`
}
