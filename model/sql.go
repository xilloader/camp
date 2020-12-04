package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        int64     `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// github.com/jinzhu/grom
type JzModel struct {
	ID        int64      `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// gorm.io/gorm
type IoModel struct {
	ID        int64        `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
