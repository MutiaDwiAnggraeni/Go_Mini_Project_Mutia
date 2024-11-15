package models

import (
"time"
)

type Transaction struct {
ID uint `gorm:"primaryKey;autoIncrement" json:"id_transaction"`
UserID uint `gorm:"not null" json:"id_user"`
User User `gorm:"foreignKey:UserID"`
ProductID uint `gorm:"not null" json:"id_product"`
Product Product `gorm:"foreignKey:ProductID"`
Status string `gorm:"type:varchar(50)" json:"status"`
CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}