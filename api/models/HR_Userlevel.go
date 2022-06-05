package models

import(
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type HrUserlevel struct {
	ID string `gorm:"primary_key;size:255;" json:"id"`
	UserlevelName string `gorm:"size:255;not null;unique" json:"userlevel_name"`
	IsActive bool `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *HrUserlevel) BeforeSave() error {
	generateID := uuid.New().String()
	if u.ID == "" {
		u.ID = generateID
	}
	return nil
}

func (u *HrUserlevel) Prepare() {
	u.UserlevelName = html.EscapeString(strings.TrimSpace(u.UserlevelName))
}

func (u *HrUserlevel) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.UserlevelName == "" {
			return errors.New("required user level Name")
		}
		return nil
	case "create":
		if u.UserlevelName == "" {
			return errors.New("required user level Name")
		}
		return nil
	default:
		if u.UserlevelName == "" {
			return errors.New("required user level Name")
		}
		return nil
	}
}

func (u *HrUserlevel) SaveHrUserlevel(db *gorm.DB) (*HrUserlevel, error) {
	err = db.Debug().Model(&HrUserlevel{}).Create(&u).Error
	if err != nil {
		return &HrUserlevel{}, err
	}
	return u, nil
}

func (u *HrUserlevel) FindAllHrUserlevel(db *gorm.DB) (*[]HrUserlevel, error) {
	userlevel := []HrUserlevel{}
	err = db.Debug().Model(&HrUserlevel{}).Limit(100).Find(&userlevel).Error
	if err != nil {
		return &[]HrUserlevel{}, err
	}
	return &userlevel, err
}

func (u *HrUserlevel) FindHrUserlevelByID(db *gorm.DB, pid string) (*HrUserlevel, error) {
	err = db.Debug().Model(&HrUserlevel{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return &HrUserlevel{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &HrUserlevel{}, errors.New("User level not found")
	}
	return u, err
}

func (u *HrUserlevel) UpdateAUserlevel(db *gorm.DB, pid string) (*HrUserlevel, error) {


	db = db.Debug().Model(&HrUserlevel{}).Where("id = ?", pid).Take(&HrUserlevel{}).UpdateColumns(
		map[string]interface{}{
			"userlevel_name": u.UserlevelName,
			"is_active": u.IsActive,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		log.Fatal(err)
		return &HrUserlevel{}, db.Error
	}

	err = db.Debug().Model(&HrUserlevel{}).Where("id = ?", pid).Take(&u).Error

	if err != nil {
		return &HrUserlevel{}, err
	}
	return u, nil

}

func (u *HrUserlevel) DeleteAUserlevel(db *gorm.DB, pid string) (int64, error) {

	db = db.Debug().Model(&HrUserlevel{}).Where("id = ?", pid).Take(&HrUserlevel{}).Delete(&HrUserlevel{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
