package Models

type Products struct {
	ID          uint64  `json:"id"`
	ProductID   string  `json:"productID"`
	NameProduct string  `json:"nameProduct"`
	Price       float64 `json:"price"`
	CampaignID  uint64  `json:"campaign_id"`
}
