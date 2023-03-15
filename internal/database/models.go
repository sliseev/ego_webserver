package database

import "time"

type Driver struct {
	ID        int    `gorm:"primaryKey"`
	FullName  string `gorm:"text; not null; uniqueIndex"`
	LicenseId string `gorm:"text; not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
