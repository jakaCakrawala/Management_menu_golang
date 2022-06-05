package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)


type Menu struct {

	ID	string    `gorm:"primary_key;" json:"id"`
	MenuName string    `gorm:"size:255;not null;unique" json:"menu_name"`
	Link string   `gorm:"size:255;" json:"link"`
	IsActive bool	`gorm:"default:true" json:"is_active"`
	Parent bool `gorm:"default:true" json:"parent"`
	Icon string `gorm:"size:255;" json:"icon"`
	Order int `gorm:"default:0" json:"order"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Menu) BeforeSave() error {
	generateID := uuid.New().String()
	if u.ID == "" {
		u.ID = generateID
	}
	return nil
}

func (u * Menu) Prepare(){
	u.MenuName = html.EscapeString(strings.TrimSpace(u.MenuName))
	u.Link = html.EscapeString(strings.TrimSpace(u.Link))
	u.Icon = html.EscapeString(strings.TrimSpace(u.Icon))
}

func (u * Menu) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.MenuName == "" {
			return errors.New("required menu name")
		}
		if u.Link == "" {
			return errors.New("required link")
		}
		return nil
	case "create":
		if u.MenuName == "" {
			return errors.New("required menu name")
		}
		if u.Link == "" {
			return errors.New("required link")
		}
		return nil
	default:
		if u.MenuName == "" {
			return errors.New("required menu name")
		}
		if u.Link == "" {
			return errors.New("required link")
		}
		return nil
	
	}
	
}

func (u *Menu) SaveMenu(db *gorm.DB) (*Menu, error) {

	err = db.Debug().Model(&Menu{}).Create(&u).Error
	if err != nil {
		return &Menu{}, err
	}
	return u, nil

}

func (u *Menu) FindAllMenu(db * gorm.DB) (*[]Menu, error) {

	menus := []Menu{}
	err = db.Debug().Model(&Menu{}).Limit(100).Find(&menus).Error
	if err != nil {
		return &[]Menu{}, err
	}
	return &menus, err
}

func (u *Menu) FindMenuByID(db *gorm.DB, pid string) (*Menu, error) {

	err = db.Debug().Model(&Menu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &Menu{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Menu{}, errors.New("Menu not found")
	}
	return u, err
}

func (u *Menu) UpdateMenu(db *gorm.DB, pid string) (*Menu, error) {

	db = db.Debug().Model(&Menu{}).Where("id = ?", pid).Take(&Menu{}).UpdateColumns(
		map[string]interface{}{
			"menu_name": u.MenuName,
			"link": u.Link,
			"is_active": u.IsActive,
			"parent": u.Parent,
			"icon": u.Icon,
			"order": u.Order,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		log.Fatal(err)
		return &Menu{}, db.Error
	}

	err = db.Debug().Model(&Menu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &Menu{}, err
	}
	return u, nil
}

func (u *Menu) DeleteMenu(db *gorm.DB, pid string) (int64, error) {

	db = db.Debug().Model(&Menu{}).Where("id = ?", pid).Take(&Menu{}).Delete(&Menu{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Menu not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

