package model

import (
	"general/config"

	"gorm.io/gorm"
)

type PromoStruct struct {
	ID     int    `json:"id_promo"`
	Banner string `json:"img"`
	Status int    `json:"status"`
}

type PromoCountStruct struct {
}

type PromoDataStruct struct {
	CurrentPage int           `json:"current_page"`
	LastPage    float64       `json:"last_page"`
	PerPage     int           `json:"per_page"`
	Total       int64         `json:"total"`
	Detail      []interface{} `json:"detail"`
}

func PromoAll(db *gorm.DB, page int, count int64) (results []PromoStruct, err error) {
	var offset int
	if page > 1 {
		offset = (page - 1) * 10
	} else {
		offset = 0
	}
	rows, err := db.Raw("SELECT * From promo Where status = 1 Limit 10 Offset ?", offset).Rows()
	defer rows.Close()
	if err != nil {
		config.Throw(err.Error())
	}
	for rows.Next() {
		db.ScanRows(rows, &results)
	}
	return results, nil
}
