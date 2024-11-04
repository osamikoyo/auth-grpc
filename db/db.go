package db

import (
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"auth-grpc/db/models"
)

func Register(u models.User) error {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		loger.Error(err.Error())
		return err
	}
	if err := db.Create(&u).Error; err != nil {
		loger.Error(err.Error())
		return err
	}
	return nil
}

const Key = "klpdfsdlfsdldklfvjsffcasdvocadvkalisjfiosadvhnsjdkvbnasjcbahsgfdfvfhldsgdfsdghhdghg"

func GenerateToken(u models.User) string {
	payload := jwt.MapClaims{
		"sub": u.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := tkn.SignedString(Key)
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
	}
	return t
}
func Get_Prof(user models.User) models.User {
	var u models.User

	dsn := "host=localhost user=osami password= dbname=users port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
	}

	if err := db.Where(&models.User{Email: user.Email, Token: user.Token}).Find(&u); err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error.Error())
	}

	return u
}
func Login_User(u models.User) error {
	dsn := "host=localhost user=osami password= dbname=users port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	//инициализация логера
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err != nil {
		loger.Error(err.Error())
		return err
	}
	tkn := GenerateToken(u)
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		loger.Error(err.Error())
		return err
	}
	if errs := db.Where("Email = ?", u.Email).Where("Password = ?", string(pass)).Update("Token", tkn).Error; errs != nil {
		loger.Error(errs.Error())
		return errs
	}
	return nil
}
