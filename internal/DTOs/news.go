package DTOs

type News struct {
	Id         int64   `json:"Id" validate:"omitempty"`
	Title      string  `json:"Title" validate:"omitempty"`
	Content    string  `json:"Content" validate:"omitempty"`
	Categories []int64 `json:"Categories" validate:"omitempty,dive,min=1"`
}
