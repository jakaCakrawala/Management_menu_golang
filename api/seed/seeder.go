package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jakaCakrawala/smartoffice_api/api/models"
)


var userLevel = []models.HrUserlevel{
	{
		ID: "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		UserlevelName: "Admin",
	},
}

var users = []models.User{
	{
		NIK:         "181101",
		FullName:    "Jaka Cakrawala",
		Email:       "admin@eltran.co.id",
		Address:     "",
		Phone:       "",
		BirthPlace:  "",
		Image:       "",
		Password:    "password",
		PIN:         "",
		IsActive:    true,
		IdUserLevel: "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		IdWorkUnit:  "e841c308-b6a7-45fa-89fc-68c4872f7a1d",
	},
	{
		NIK:         "181102",
		FullName:    "Jaka Cakrawala",
		Email:       "admin2@eltran.co.id",
		Address:     "",
		Phone:       "",
		BirthPlace:  "",
		Image:       "",
		Password:    "password",
		PIN:         "",
		IsActive:    true,
		IdUserLevel: "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		IdWorkUnit:  "e841c308-b6a7-45fa-89fc-68c4872f7a1d",
	},
}

var menu = []models.Menu{
	{
		ID: "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",
		MenuName: "Settings",
		Link: "#",
		Icon: "fas fa-cog",
		IsActive: true,
		Order: 1,
	},
}

var submenu = []models.SubMenu{
	{
		ID: "1581590c-efb2-43c7-ad1d-701b19bc57aa",
		SubmenuName: "Menu",
		Icon: "fas fa-caret-right",
		Link: "settings/menu",
		Order: 1,
		IsActive: 	true,
		IDMenu: "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",

	},
	{
		ID: "acab7127-ab70-42bb-a1d9-e41e619748ee",
		SubmenuName: "Submenu",
		Icon: "fas fa-caret-right",
		Link: "settings/submenu",
		Order: 2,
		IsActive: 	true,
		IDMenu: "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",

	},
	{
		ID: "5ee1a893-d716-407d-a749-b58e893c6193",
		SubmenuName: "User Level",
		Icon: "fas fa-caret-right",
		Link: "settings/userlevel",
		Order: 3,
		IsActive: 	true,
		IDMenu: "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",

	},
	{
		ID: "5bef987b-1fe4-483b-b426-83a814e4281d",
		SubmenuName: "Users",
		Icon: "fas fa-caret-right",
		Link: "settings/users",
		Order: 4,
		IsActive: 	true,
		IDMenu: "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",

	},
}


var AccessMenu = []models.AccessMenu{
	{
		IDMenu : "f2d047dc-d9d0-4b13-841c-026c7d58a3ad",
		IdLevel : "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		View : true,
		Add : true,
		Edit : true,
		Delete : true,
		Print: true,
		Upload: true,
	},
}

var AccessSubmenu = []models.AccessSubmenu{
	{
		IdSubMenu : "1581590c-efb2-43c7-ad1d-701b19bc57aa",
		IdLevel : "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		View: true,
		Add: true,
		Edit: true,
		Delete: true,
		Print: true,
		Upload: true,		
	},
	{
		IdSubMenu : "acab7127-ab70-42bb-a1d9-e41e619748ee",
		IdLevel : "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		View: true,
		Add: true,
		Edit: true,
		Delete: true,
		Print: true,
		Upload: true,		
	},
	{
		IdSubMenu : "5ee1a893-d716-407d-a749-b58e893c6193",
		IdLevel : "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		View: true,
		Add: true,
		Edit: true,
		Delete: true,
		Print: true,
		Upload: true,		
	},
	{
		IdSubMenu : "5bef987b-1fe4-483b-b426-83a814e4281d",
		IdLevel : "7a8f9097-cc09-4d5f-bebb-07062a089e1e",
		View: true,
		Add: true,
		Edit: true,
		Delete: true,
		Print: true,
		Upload: true,		
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Menu{}, &models.SubMenu{}, &models.AccessMenu{}, &models.AccessSubmenu{}, &models.HrUserlevel{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Menu{}, &models.SubMenu{}, &models.AccessMenu{}, &models.AccessSubmenu{}, &models.HrUserlevel{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}


	for i := range userLevel {
		err = db.Debug().Model(&models.HrUserlevel{}).Create(&userLevel[i]).Error
		
		if err != nil {
			log.Fatalf("cannot seed userlevel table: %v", err)
		}
	}

	for j := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[j]).Error
		
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for k := range menu {
		err = db.Debug().Model(&models.Menu{}).Create(&menu[k]).Error
		
		if err != nil {
			log.Fatalf("cannot seed menus table: %v", err)
		}
	}

	for l := range submenu {
		err = db.Debug().Model(&models.SubMenu{}).Create(&submenu[l]).Error
		
		if err != nil {
			log.Fatalf("cannot seed submenus table: %v", err)
		}
	}

	for m := range AccessMenu {
		err = db.Debug().Model(&models.AccessMenu{}).Create(&AccessMenu[m]).Error
		
		if err != nil {
			log.Fatalf("cannot seed accessmenu table: %v", err)
		}
	}

	for n := range AccessSubmenu {
		err = db.Debug().Model(&models.AccessSubmenu{}).Create(&AccessSubmenu[n]).Error
		
		if err != nil {
			log.Fatalf("cannot seed accesssubmenu table: %v", err)
		}
	}

}
