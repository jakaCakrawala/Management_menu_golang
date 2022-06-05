package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type User struct {
	ID	string    `gorm:"primary_key;" json:"id"`
	NIK	string    `gorm:"size:255;not null;unique" json:"NIK"`
	FullName	string    `gorm:"size:255;" json:"full_name"`
	Email	string    `gorm:"size:100;unique" json:"email"`
	Address	string    `gorm:"size:255;" json:"address"`
	Phone	string    `gorm:"size:255;" json:"phone"`
	BirthDate	time.Time	`gorm:"size:255;" json:"birth_date"`
	BirthPlace	string    `gorm:"size:255;" json:"birth_place"`
	Image	string    `gorm:"size:200;" json:"image"`
	Password	string    `gorm:"size:255;not null;" json:"password"`
	PIN	string    `gorm:"size:255;" json:"pin"`
	IsActive	bool	`gorm:"default:true" json:"is_active"`
	IdUserLevel string `gorm:"size:255;" json:"id_user_level"`
	IdWorkUnit string `gorm:"size:255;" json:"id_work_unit"`
	CreatedAt	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.NIK = html.EscapeString(strings.TrimSpace(u.NIK))
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Address = html.EscapeString(strings.TrimSpace(u.Email))
	u.Phone = html.EscapeString(strings.TrimSpace(u.Email))
	u.BirthDate = time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	u.BirthPlace = html.EscapeString(strings.TrimSpace(u.BirthPlace))
	u.Image = html.EscapeString(strings.TrimSpace(u.Image))
	u.IsActive = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}


func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.NIK == "" {
			return errors.New("required NIK")
		}
		if u.FullName == "" {
			return errors.New("required FullName")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil

	default:
		if u.NIK == "" {
			return errors.New("required NIK")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error){
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(1000).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error){

	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"NIK":  u.NIK,
			"full_name":  u.FullName,
			"email":     u.Email,
			"address":  u.Address,
			"phone":  u.Phone,
			"birth_date":  u.BirthDate,
			"birth_place":  u.BirthPlace,
			"image":  u.Image,
			"pin":  u.PIN,
			"is_active":  u.IsActive,
			"id_user_level":  u.IdUserLevel,
			"id_work_unit":  u.IdWorkUnit,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}