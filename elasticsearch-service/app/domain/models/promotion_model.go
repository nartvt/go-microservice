package models

type PromotionModel struct {
	Id            int
	Code          string
	CreatedAt     string
	UpdatedAt     string
	MaxActiveTime int
}

func (*PromotionModel) Index() string {
	return "promotion_index"
}
