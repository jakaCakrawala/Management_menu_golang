package models

import(
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type AccessMenu struct {

	ID	string    `gorm:"primary_key" json:"id"`
	IDMenu string	`gorm:"size:255;not null;unique" json:"id_menu"`
	View bool	`gorm:"default:true" json:"view"`
	Add bool `gorm:"default:true" json:"add"`
	Edit bool `gorm:"default:true" json:"edit"`
	Delete bool `gorm:"default:true" json:"delete"`
	Print bool `gorm:"default:true" json:"print"`
	Upload bool `gorm:"default:true" json:"upload"`
	IdLevel string `gorm:"size:255;" json:"id_level"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}

type ResultAccessMenu struct {
	IDAccess	string // Table Access_Menus
	IDMenu string // Table Menus
	MenuName string // Table Menus
	Link string // Table Menus
	Icon string // Table Menus
	Order int // Table Menus
	View bool // Table Access_Menus
	Add bool // Table Access_Menus
	Edit bool // Table Access_Menus
	Delete bool // Table Access_Menus
	Print bool // Table Access_Menus
	Upload bool // Table Access_Menus
	Parent bool // Table Menus
	IsActive bool // Table Menus
	IdLevel string 	// Table hr_userlevels
	UserlevelName string // Table hr_userlevels
}	

var err error

func (u *AccessMenu) BeforeSave() error {
	generateID := uuid.New().String()
	if u.ID == "" {
		u.ID = generateID
	}
	return nil
}

func (u *AccessMenu) Prepare() {
	u.IDMenu = html.EscapeString(strings.TrimSpace(u.IDMenu))
	u.IdLevel = html.EscapeString(strings.TrimSpace(u.IdLevel))
}

func (u *AccessMenu) Validate(action string) error{
	switch strings.ToLower(action) {
	case "update":
		if u.IDMenu == "" {
			return errors.New("required menu name")
		}
		if u.IdLevel == "" {
			return errors.New("required level")
		}
		return nil
	case "create":
		if u.IDMenu == "" {
			return errors.New("required Menu Name")
		}
		if u.IdLevel == "" {
			return errors.New("required Level")
		}
		return nil
	default:
		if u.IDMenu == "" {
			return errors.New("required Menu Name")
		}
		if u.IdLevel == "" {
			return errors.New("required Level")
		}
		return nil
	}
}

func (u *AccessMenu) SaveAccessMenu(db *gorm.DB) (*AccessMenu, error) {
	err = db.Debug().Model(&AccessMenu{}).Create(&u).Error
	if err != nil {
		return &AccessMenu{}, err
	}
	return u, nil
}

func (u *AccessMenu) FindAllAccessMenu(db *gorm.DB) (*[]AccessMenu, error) {
	accessmenu := []AccessMenu{}
	err = db.Debug().Model(&AccessMenu{}).Limit(100).Find(&accessmenu).Error
	if err != nil {
		return &[]AccessMenu{}, err
	}
	return &accessmenu, err
}

func (u *AccessMenu) FindAllAccessMenuByID(db *gorm.DB, id string) (*[]AccessMenu, error) {
	accessmenu := []AccessMenu{}
	err = db.Debug().Model(&AccessMenu{}).Where("id_menu = ?", id).Find(&accessmenu).Error
	if err != nil {
		return &[]AccessMenu{}, err
	}
	return &accessmenu, err
}

func (u *AccessMenu) ViewAccessMenu(db *gorm.DB) (*[]ResultAccessMenu, error) {
	resultAccessMenus := []ResultAccessMenu{}
	err = db.Debug().Model(&AccessMenu{}).Select("access_menus.id as id_access, access_menus.id_menu, menus.menu_name, menus.link, menus.icon, menus.order, access_menus.view, access_menus.view, access_menus.add, access_menus.edit, access_menus.delete, access_menus.print, access_menus.upload, menus.parent, menus.is_active, hr_userlevels.id as id_level, hr_userlevels.userlevel_name").
	Joins("JOIN menus ON access_menus.id_menu = menus.id").
	Joins("JOIN hr_userlevels ON access_menus.id_level = hr_userlevels.id").Scan(&resultAccessMenus).Error
	if err != nil {
		return &[]ResultAccessMenu{}, err
	}
	return &resultAccessMenus, err
}

func (u *AccessMenu) UpdateAccessMenu(db *gorm.DB, pid string) (*AccessMenu, error) {

	db = db.Debug().Model(&AccessMenu{}).Where("id = ?", pid).Take(&AccessMenu{}).UpdateColumns(
		map[string]interface{}{
			"id_menu": u.IDMenu,
			"view": u.View,
			"add": u.Add,
			"edit": u.Edit,
			"delete": u.Delete,
			"print": u.Print,
			"upload": u.Upload,
			"id_level": u.IdLevel,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		log.Fatal(err)
		return &AccessMenu{}, db.Error
	}

	err = db.Debug().Model(&AccessMenu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &AccessMenu{}, err
	}
	return u, nil
	
}

func (u *AccessMenu) DeleteAccessMenu(db *gorm.DB, pid string) (int64, error) {

	db = db.Debug().Model(&AccessMenu{}).Where("id = ?", pid).Take(&AccessMenu{}).Delete(&AccessMenu{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
