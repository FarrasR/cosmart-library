package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryHelper struct {
	c *gin.Context
}

func NewQueryHelper(c *gin.Context) *QueryHelper {
	return &QueryHelper{c}
}

func (h *QueryHelper) GetInt(key string, defValue int) int {

	value, ok := h.c.GetQuery(key)

	if !ok {
		return defValue
	}

	if value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}

	return defValue
}

func (h *QueryHelper) GetString(key string, defValue string) string {
	value, ok := h.c.GetQuery(key)

	if !ok {
		return defValue
	}

	return value
}
