package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Role      string    `gorm:"default:'warga'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminRequest struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Position  string    `json:"position"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	KtpURL    string    `json:"ktp_url"`
	Status    string    `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
