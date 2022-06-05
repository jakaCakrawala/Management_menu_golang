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

type AccessSubmenu struct {
	
	ID	string    `gorm:"primary_key" json:"id"`
	IdSubMenu string	`gorm:"size:255;not null;unique" json:"id_sub_menu"`
	View bool `gorm:"default:true" json:"view"`
	Add bool `gorm:"default:true" json:"add"`
	Edit bool `gorm:"default:true" json:"edit"`
	Delete bool `gorm:"default:true" json:"delete"`
	Print bool `gorm:"default:true" json:"print"`
	Upload bool `gorm:"default:true" json:"upload"`
	IdLevel string `gorm:"size:255;" json:"id_level"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}

type ResultAccessSubMenu struct {
	IDAccess	string // Table Access_Menus
	IDSubMenu string // Table SubMenus
	SubmenuName string // Table SubMenus
	Link string // Table Menus
	Icon string // Table Menus
	Order int // Table Menus
	View bool // Table Access_Menus
	Add bool // Table Access_Menus
	Edit bool // Table Access_Menus
	Delete bool // Table Access_Menus
	Print bool // Table Access_Menus
	Upload bool // Table Access_Menus
	IsActive bool // Table SubMenus
	IDMenu string // Table Menus
	MenuName string // Table Menus
	IdLevel string 	// Table hr_userlevels
	UserlevelName string // Table hr_userlevels
}	

func (u *AccessSubmenu) BeforeSave() error {
	generateID := uuid.New().String()
	if u.ID == "" {
		u.ID = generateID
	}
	return nil
}

func (u *AccessSubmenu) Prepare() {
	u.IdSubMenu = html.EscapeString(strings.TrimSpace(u.IdSubMenu))
	u.IdLevel = html.EscapeString(strings.TrimSpace(u.IdLevel))
}

func (u *AccessSubmenu) Validate(action string)	error{	
	switch strings.ToLower(action) {
	case "update":
		if u.IdSubMenu == "" {
			return errors.New("required Sub Menu Name")
		}
		if u.IdLevel == "" {
			return errors.New("required Level")
		}
		return nil
	case "create":
		if u.IdSubMenu == "" {
			return errors.New("required Sub Menu Name")
		}
		if u.IdLevel == "" {
			return errors.New("required Level")
		}
		return nil
	default:
		if u.IdSubMenu == "" {
			return errors.New("required Sub Menu Name")
		}
		if u.IdLevel == "" {
			return errors.New("required Level")
		}
		return nil
	}
}

func (u *AccessSubmenu) SaveAccessMenu(db *gorm.DB) (*AccessSubmenu, error) {
	err = db.Debug().Model(&AccessSubmenu{}).Create(&u).Error
	if err != nil {
		return &AccessSubmenu{}, err
	}
	return u, nil
}

func (u *AccessSubmenu) FindAllAccessSubmenu(db *gorm.DB) (*[]AccessSubmenu, error) {
	accessMenu := []AccessSubmenu{}
	err = db.Debug().Model(&AccessSubmenu{}).Limit(100).Find(&accessMenu).Error
	if err != nil {
		return &[]AccessSubmenu{}, err
	}
	return &accessMenu, err
}

func (u *AccessSubmenu) FindAccessSubmenuById(db *gorm.DB, pid string) (*AccessSubmenu, error) {
	err = db.Debug().Model(&AccessSubmenu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &AccessSubmenu{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &AccessSubmenu{}, errors.New("access menu not found")
	}
	return u, err
}


func (u *AccessSubmenu) ViewAccessSubMenu(db *gorm.DB) (*[]ResultAccessSubMenu, error) {
	resultAccessSubMenus := []ResultAccessSubMenu{}
	err = db.Debug().Model(&AccessSubmenu{}).Select("access_submenus.id as id_access, access_submenus.id_sub_menu , sub_menus.submenu_name, sub_menus.link,sub_menus.icon, sub_menus.order, access_submenus.view, access_submenus.view, access_submenus.add, access_submenus.edit, access_submenus.delete, access_submenus.print, access_submenus.upload, sub_menus.is_active, menus.id as id_menu, menus.menu_name, hr_userlevels.id as id_level, hr_userlevels.userlevel_name").
	Joins("JOIN sub_menus ON access_submenus.id_sub_menu = sub_menus.id").
	Joins("JOIN hr_userlevels ON access_submenus.id_level = hr_userlevels.id").
	Joins("JOIN menus ON sub_menus.id_menu = menus.id").Scan(&resultAccessSubMenus).Error
	if err != nil {
		return &[]ResultAccessSubMenu{}, err
	}
	return &resultAccessSubMenus, err
}


func (u *AccessSubmenu) UpdateAccessSubmenu(db *gorm.DB, pid string) (*AccessSubmenu, error) {

	db = db.Debug().Model(&AccessSubmenu{}).Where("id = ?", pid).Take(&AccessSubmenu{}).UpdateColumns(
		map[string]interface{}{
			"id_sub_menu": u.IdSubMenu,
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
		return &AccessSubmenu{}, db.Error
	}
	
	err = db.Debug().Model(&AccessSubmenu{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &AccessSubmenu{}, err
	}
	return u, nil
}

func (u *AccessSubmenu) DeleteAccessSubmenu(db *gorm.DB, pid string) (int64, error) {
	db = db.Debug().Model(&AccessSubmenu{}).Where("id = ?", pid).Take(&AccessSubmenu{}).Delete(&AccessSubmenu{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}