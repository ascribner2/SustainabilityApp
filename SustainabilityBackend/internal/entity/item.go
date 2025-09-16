package entity

type ItemEntity interface {
	GetTitle() string
	GetOffset() float64
	GetDate() string
}

type Item struct {
	Title  string  `json:"title"`
	Offset float64 `json:"offset"`
	Date   string  `json:"date"`
}

func (i *Item) GetTitle() string {
	return i.Title
}

func (i *Item) GetOffset() float64 {
	return i.Offset
}

func (i *Item) GetDate() string {
	return i.Date
}
