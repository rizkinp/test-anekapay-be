package entity

type Animal struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	Legs      int    `json:"legs"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
