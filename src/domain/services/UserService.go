package services

import (
	"AlexSilverson/lab2/src/domain/entity"
	"AlexSilverson/lab2/src/utility"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var SecretKey = "it_is_my_password"

type userService struct {
	db *gorm.DB
}
type UserService interface {
	AddUser(user entity.User) error
	UpdRole(id uint) error
	Login(user entity.User) (string, error)
	GetUserById(id uint) (*entity.User, error)
}

func (r userService) GetUserById(id uint) (*entity.User, error) {
	var user entity.User
	erdb := r.db.First(&user, id)
	if erdb.Error != nil {
		return &user, erdb.Error
	}
	return &user, nil
}

func (r userService) AddUser(user entity.User) error {
	user.HashPas, _ = utility.GetHesh(user.HashPas)
	erdb := r.db.Create(&user)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r userService) UpdRole(id uint) error {
	var curUser entity.User
	erdb := r.db.First(&curUser, id)
	if erdb.Error != nil {
		return erdb.Error
	}
	curUser.Role = 1
	erdb = r.db.Save(&curUser)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r userService) Login(user entity.User) (string, error) {
	var dbUser entity.User
	erdb := r.db.Where("login = ?", user.Login).First(&dbUser)
	fmt.Println(dbUser, user)
	if erdb.Error != nil {
		return "", erdb.Error
	}
	fmt.Println("here in login before pass")
	if !utility.Compare(user.HashPas, dbUser.HashPas) {
		fmt.Print("here in login bad pass")
		return "", errors.New("invalid password")
	}

	exTime := time.Now().Add(5 * time.Minute)
	fmt.Println("here in Login token 1")
	claims := &entity.Claims{
		Role: dbUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   dbUser.Login,
			ExpiresAt: exTime.Unix(),
		},
	}
	fmt.Println("here in Login token 2")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("here in Login token 3")
	tokenString, err := token.SignedString([]byte(SecretKey))
	fmt.Println("here in Login error", err)
	if err != nil {
		return "", errors.New("token couldnt be generated")
	}
	return tokenString, nil
}

func NewUserSevice(db *gorm.DB) UserService {
	return &userService{db: db}
}
