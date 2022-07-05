package seeds

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/db"
	"github.com/RIDCHA-DATA/golang-rest-api/internal/pkg/models"
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	_ "k8s.io/client-go/kubernetes/typed/core/v1/fake"
)

func randomInt(min, max int) int {

	return rand.Intn(max-min) + min
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func seedUsers(db *gorm.DB) {

	var action []models.Action
	db.Model(&action).Related(&action, "Action")
	actionCount := len(action)
	actionToSeed := 20
	actionToSeed -= actionCount
	if actionToSeed > 0 {
		for i := 0; i < actionToSeed; i++ {
			action := models.Action{Action: fake.FirstName()}
			// No need to add the role as we did for seedAdmin, it is added by the BeforeSave hook
			db.Set("gorm:association_autoupdate", false).Create(&action)
		}
	}
}

func Seed() {
	db := db.SetupDB()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(db)
	seedUsers(db)
}
