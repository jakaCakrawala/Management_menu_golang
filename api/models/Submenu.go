package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type SubMenu struct {
	ID string `gorm:"primary_key;size:255;" json:"id"`
	SubmenuName string `gorm:"size:255;not null;unique" json:"submenu_name"`
	Link string `gorm:"size:255;" json:"link"`
	IsActive bool `gorm:"default:true" json:"is_active"`
	Icon string `gorm:"size:255;" json:"icon"`
	Order int `gorm:"default:0" json:"order"`
	IDMenu string `gorm:"size:255;" json:"id_menu"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *SubMenu) BeforeSave() error {
	generateID := uuid.New().String()
	if u.ID == "" {
		u.ID = generateID
	}
	return nil
}

func (u *SubMenu) Prepare() {
	u.SubmenuName = html.EscapeString(strings.TrimSpace(u.SubmenuName))
	u.Link = html.EscapeString(strings.TrimSpace(u.Link))
	u.Icon = html.EscapeString(strings.TrimSpace(u.Icon))
	u.IDMenu = html.EscapeString(strings.TrimSpace(u.IDMenu))
}

func (u * SubMenu) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.SubmenuName == "" {
			return errors.New("required Menu Name")
		}
		if u.Link == "" {
			return errors.New("required Link")
		}
		return nil
	case "create":
		if u.SubmenuName == "" {
			return errors.New("required Menu Name")
		}
		if u.Link == "" {
			return errors.New("required Link")
		}
		return nil
	default:
		if u.SubmenuName == "" {
			return errors.New("required Menu Name")
		}
		if u.Link == "" {
			return errors.New("required Link")
		}
		return nil
	
	}
	
}

func (u *SubMenu) SaveSubMenu(db *gorm.DB) (*SubMenu, error) {
	err = db.Debug().Model(&SubMenu{}).Create(&u).Error
	if err != nil {
		return &SubMenu{}, err
	}
	return u, nil

}

func (u *SubMenu) FindAllSubMenu(db * gorm.DB) (*[]SubMenu, error) {
	submenu := []SubMenu{}
	err = db.Debug().Model(&SubMenu{}).Limit(1000).Find(&submenu).Error
	if err != nil {
		return &[]SubMenu{}, err
	}
	return &submenu, nil

}

func (u *SubMenu) FindSubMenuByID(db *gorm.DB, pid string) (*SubMenu, error) {
	err = db.Debug().Model(&SubMenu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &SubMenu{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SubMenu{}, errors.New("SubMenu not found")
	}
	return u, err
}

func (u *SubMenu) UpdateSubMenu(db *gorm.DB, pid string) (*SubMenu, error) {
	db = db.Debug().Model(&SubMenu{}).Where("id = ?", pid).Take(&SubMenu{}).UpdateColumns(
		map[string]interface{}{
			"submenu_name": u.SubmenuName,
			"link": u.Link,
			"is_active": u.IsActive,
			"icon": u.Icon,
			"order": u.Order,
			"id_menu": u.IDMenu,
			"updated_at": time.Now(),
		},
	)
	
	if db.Error != nil {
		log.Fatal(err)
		return &SubMenu{}, db.Error
	}

	err = db.Debug().Model(&SubMenu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &SubMenu{}, err
	}
	return u, nil
}

func (u *SubMenu) DeleteSubMenu(db *gorm.DB, pid string) (int64, error) {

	db = db.Debug().Model(&SubMenu{}).Where("id = ?", pid).Take(&SubMenu{}).Delete(&SubMenu{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}