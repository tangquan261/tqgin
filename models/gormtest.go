package models

import (
	"fmt"
)

type Animal struct {
	ID       int64  `gorm:"primary_key"`
	PlayerID int64  `gorm:"default:1;unique;not null"`
	Name     string `gorm:"default:'galeone';not null;unique"`
	Age      int64
}

func Gromtest() {
	return
	fmt.Println("gorm test")
	//DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Animal{})

	var animal Animal
	animal.ID = 11

	DB.Table("tq_animal").Where("id = (?)", animal.ID).Updates(map[string]interface{}{"player_id": 0})
	//DB.Model(&animal).UpdateColumn(Animal{PlayerID: 1})
	return
	//var animal Animal
	animal.Age = 1
	animal.Name = "11"
	//DB.Create(&animal)

	err := DB.First(&animal, "id = ?", animal.ID).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(animal)

	fmt.Println("has table", DB.HasTable(&UserInfo{}), DB.HasTable(&Account{}))
}
