package controller

import (
	"general/config"
	"general/model"

	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Promo struct {
	gorm.Model
	Title string
}

func GetPromo(c *gin.Context) {
	db := config.DB.Begin()
	config.Block{
		Try: func() {
			var count int64
			var pagelast float64
			page, _ := c.GetQuery("page")
			pageint, _ := strconv.Atoi(page)
			db.Table("promo").Where("status = 1").Count(&count)
			pagelast = float64(count) / 10
			results, err := model.PromoAll(db, pageint, count)
			promoDataStruct := model.PromoDataStruct{}
			promocountStruct := count
			if err != nil {
				config.Throw(err.Error())
			}
			promoDataStruct.CurrentPage = pageint
			promoDataStruct.LastPage = math.Ceil(pagelast)
			promoDataStruct.PerPage = 10
			promoDataStruct.Total = promocountStruct
			for _, result := range results {
				temp := struct {
					ID     int
					Banner string
					Status int
				}{
					ID:     result.ID,
					Banner: result.Banner,
					Status: result.Status,
				}
				promoDataStruct.Detail = append(promoDataStruct.Detail, temp)
			}

			db.Commit()
			config.Response(c, "Sukses menampilkan data", promoDataStruct)
		},
		Catch: func(e config.Exception) {
			db.Rollback()
			config.ErrorResponse(c, e)
		},
	}.Do()
}
