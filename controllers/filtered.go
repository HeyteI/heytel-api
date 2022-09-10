package controllers

import (
	"Heytel/database"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"Heytel/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		
		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}

func uuidToStringFunc() mapstructure.DecodeHookFunc {
    return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
        if f.Kind() != reflect.String {
            return data, nil
        }
        if t != reflect.TypeOf(uuid.UUID{}) {
            return data, nil
        }

        return uuid.Parse(data.(string))
    }
}

func Decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(toTimeHookFunc(), uuidToStringFunc()),
		Result:     result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

func FilteredUpdate(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var responseUser models.User
	var responseRoom models.Room
	var responseInvoice models.Invoice

	var data map[string]interface{}

	jsonData, err := ctx.GetRawData()
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	log.Print(data)

	master, ok := data["master"].(map[string]interface{})
	if !ok {
		panic(ok)
	}

	user, ok := master["user"].(map[string]interface{})
	if !ok {
		panic(ok)
	}
	room, ok := master["room"].(map[string]interface{})
	if !ok {
		panic(ok)
	}

	invoice, ok := master["invoice"].(map[string]interface{})
	if !ok {
		panic(ok)
	}

	log.Print(invoice)

	Decode(user, &responseUser)
	Decode(room, &responseRoom)
	Decode(invoice, &responseInvoice)

	log.Print(responseUser)
	log.Print(responseRoom)
	log.Print(responseInvoice)

	// here im getting data to update TODO TRANSACTION

	tx := db.Begin()
	tx.SavePoint("sp")

	var dbUser models.User

	if result := tx.Where("id = ?", responseUser.ID).First(&dbUser); result.Error != nil {
		log.Print("not found responseUser")
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		tx.RollbackTo("sp")
		return
	}

	tx.Model(&dbUser).Updates(responseUser)

	var dbRoom models.Room

	if result := tx.Where("id = ?", responseRoom.ID).First(&dbRoom); result.Error != nil {
		log.Print("not found responseRoom")
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		tx.RollbackTo("sp")
		return
	}

	tx.Model(&dbRoom).Updates(responseRoom)

	var dbInvoice models.Invoice

	if result := tx.Where("id = ?", responseInvoice.ID).First(&dbInvoice); result.Error != nil {
		log.Print("not found responseInvoice")
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		tx.RollbackTo("sp")
		return
	}

	tx.Model(&dbInvoice).Updates(responseInvoice)

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": jsonData})
}
