package Models

import "github.com/golang/protobuf/ptypes/timestamp"

type Order struct {
	ID          uint64      `gorm:"primaryKey" json:"id"`
	UserID      uint64      `json:"userid"`
	Address     string      `json:"address"`
	PhoneNumber string      `json:"phoneNumber"`
	Status      string      `json:"status"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt   timestamp.Timestamp
	UpdatedAt   timestamp.Timestamp
}

type OrderItem struct {
	ID         uint64   `json:"id"`
	OrderID    uint64   `json:"orderID"`
	Orders     Order    `gorm:"foreignKey:OrderID"`
	ProductID  uint64   `json:"productID"`
	Products   Products `gorm:"foreignKey:ProductID"`
	Quantity   uint64   `json:"quantity"`
	Price      uint64   `json:"price"`
	TotalPrice uint64   `json:"totalPrice"`
}

type CampaignOrder struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	CampaignID uint64    `json:"campaignID"`
	Campaigns  Campaigns `gorm:"foreignKey:CampaignID"`
	OrderID    uint64    `json:"orderID"`
	Orders     Order     `gorm:"foreignKey:OrderID"`
	Status     string    `json:"status"`
}
