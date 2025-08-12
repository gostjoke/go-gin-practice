package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 響應結構
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 分頁響應結構
type PaginatedResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// 成功響應
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// 錯誤響應
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// 分頁響應
func PaginatedSuccessResponse(c *gin.Context, data interface{}, currentPage, perPage int, total int64) {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	c.JSON(200, PaginatedResponse{
		Code:    200,
		Message: "success",
		Data:    data,
		Meta: PaginationMeta{
			CurrentPage: currentPage,
			PerPage:     perPage,
			Total:       total,
			TotalPages:  totalPages,
		},
	})
}

// 獲取分頁參數
func GetPaginationParams(c *gin.Context) (int, int) {
	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}

	return page, limit
}

// 計算偏移量
func GetOffset(page, limit int) int {
	return (page - 1) * limit
}
