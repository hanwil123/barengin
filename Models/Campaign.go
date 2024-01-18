package Models

type Campaigns struct {
	ID           uint64 `json:"id"`
	NameCampaign string `json:"nameCampaign"`
	Target       uint64 `json:"target"`
	Discount     uint64 `json:"discount"`
}
