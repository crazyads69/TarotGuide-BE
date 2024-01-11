package schemas

type Prompt struct {
	PromptID string `gorm:"column:prompt_id;primaryKey" json:"prompt_id"`
	Prompt   string `gorm:"column:prompt" json:"prompt"`
}
