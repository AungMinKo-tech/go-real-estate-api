package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func Paginate(c *gin.Context, db *gorm.DB, model interface{}, rows interface{}) (*Pagination, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var totalRows int64
	db.Model(model).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
	offset := (page - 1) * limit

	err := db.Limit(limit).Offset(offset).Find(rows).Error
	if err != nil {
		return nil, err
	}

	return &Pagination{
		Limit:      limit,
		Page:       page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Rows:       rows,
	}, nil
}
