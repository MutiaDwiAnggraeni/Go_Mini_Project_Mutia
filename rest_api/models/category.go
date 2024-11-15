package models

import "time"

type Category struct {
ID uint `gorm:"primaryKey;autoIncrement" json:"id_kategori"`
Name string `gorm:"type:varchar(255)" json:"name"`
CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}