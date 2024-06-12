package operation

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	Username    string         `gorm:"type:varchar(100);unique;not null"`
	Password    string         `gorm:"type:varchar(20);not null"`
	Bio         sql.NullString `gorm:"type:text;default:null"` // sql.NullString 可以表示空字符串
	IsPremium   bool           `gorm:"default:false"`
	LastLogin   *time.Time     `gorm:"default:null"` // 空指针；可以用 updatedAt 代替；但是为了演示空指针，选择保留
	PhoneNumber *string        `gorm:"default:null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//UserUUID    string         `gorm:"type:varchar(36);unique;not null"`
//Email       string         `gorm:"type:varchar(100);unique;not null"`

func UserOperation(db *gorm.DB) {
	// 为了保证数据一致性，当修改了 struct 后，不会删除列，以保证数据完整性，但是会添加新列
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("auto migrate error: ", err)
		return
	}

	// create
	currentTime := time.Now()
	users := []User{
		{
			Username:  "user1",
			Password:  "password1",
			IsPremium: true,
		},
		{
			Username:    "user2",
			Password:    "password2",
			IsPremium:   false,
			Bio:         sql.NullString{String: "", Valid: false}, // bio 空
			LastLogin:   nil,
			PhoneNumber: nil,
		},
		{
			Username:    "user3",
			Password:    "password3",
			IsPremium:   true,
			Bio:         sql.NullString{String: "user3 bio3", Valid: true},
			LastLogin:   &currentTime,
			PhoneNumber: func() *string { s := "122-9876-3456"; return &s }(),
		},
	}

	err = db.Create(&users).Error
	if err != nil {
		log.Fatal("create failed")
	}

	//fmt.Println(users)

	// query
	var user User
	db.First(&user, "username = ?", "user3")
	fmt.Printf("username: %s, password: %s, bio: %v, lastLogin: %v\n", user.Username, user.Password, user.Bio, user.LastLogin)
	var noPremiumUsers []User
	db.Find(&noPremiumUsers, "is_premium = ?", "false")
	fmt.Println("no premium users: ", noPremiumUsers)

	// delete all
	//db.Where("1 = 1").Delete(&User{})
	db.Exec("DELETE FROM users")

}
