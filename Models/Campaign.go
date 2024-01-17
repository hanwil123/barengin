package Models

type Campaigns struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Target   uint64 `json:"target"`
	Discount uint64 `json:"discount"`
}
