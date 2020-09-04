package models

import(
	"strings"
	"errors"

	"github.com/badoux/checkmail"
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)
type User struct {
	gorm.model
	Email string `gorm:"type:varchar(50);unique_index" json:"email"`
	Username string `gorm: "type:varchar(20);unique_index" json:"username"`
	Password string `gorm:"size:100;not null"json:"password"`
	Role string `gorm:"size:100;not null"json:"password"`
}