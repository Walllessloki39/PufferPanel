package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Server struct {
	ID   uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"-"`
	Name string    `gorm:"UNIQUE_INDEX;size:20;NOT NULL" json:"-"`
	UUID uuid.UUID `gorm:"UNIQUE_INDEX;NOT NULL" json:"-"`

	NodeID uint `gorm:"NOT NULL" json:"-"`
	Node   Node `gorm:"association_autoupdate:false" json:"-"`

	//CreatedAt time.Time `json:"-"`
	//UpdatedAt time.Time `json:"-"`
}

type Servers []*Server

func MigrateServerModel(db *gorm.DB) (err error) {
	err = db.Model(&Server{}).AddForeignKey("node_id", "nodes(id)", "RESTRICT", "RESTRICT").Error
	return
}