package models

import (
"time"
)

type Product struct {
ID uint `gorm:"primaryKey;autoIncrement" json:"id_product"`
Name string `gorm:"type:varchar(255)" json:"name"`
Description string `gorm:"type:text" json:"description"`
Price float64 `gorm:"type:decimal(10,2)" json:"price"`
CategoryID uint `gorm:"not null" json:"id_category"`
Category Category `gorm:"foreignKey:CategoryID"`
CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
