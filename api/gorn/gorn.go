package gorn

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofor-little/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorn struct {
	Data map[string]map[string]string
}

var DB *gorm.DB

func (gr *Gorn) Init(params ...any) {
	dir := ""
	if len(params) > 0 && params[0] != nil {
		dir = params[0].(string) + "/"
	}

	productionFlag := false
	if len(params) > 1 && params[1] != nil {
		productionFlag = params[1].(bool)
	}

	envFile := ".env"
	if productionFlag {
		envFile = ".env.production"
	}
	if err := env.Load(dir + envFile); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.Get("DATABASE_USER", "root"), env.Get("DATABASE_PASSWORD", ""), env.Get("DATABASE_HOST", "localhost"), env.Get("DATABASE_NAME", "gorn"))

	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   logger.Default.LogMode(logger.Info),
	})

	DB.Set("gorm:auto_preload", true)
}

func Flash(ctx *gin.Context, status uint8, msg any, data map[string]any) {
	ctx.JSON(http.StatusOK, gin.H{"status": status, "msg": msg, "data": data})
}

func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		formTag := field.Tag.Get("form")
		bindingTag := field.Tag.Get("binding")
		if formTag != "" {
			result[formTag] = bindingTag
		}
	}
	return result
}
func Atoi(str any) int {
	a, _ := strconv.Atoi(str.(string))
	return a
}

func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}

func InIntArray(num uint, arr []uint) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}

	return false
}

func RandSeq(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandNumber(n int) string {
	const letterBytes = "0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
