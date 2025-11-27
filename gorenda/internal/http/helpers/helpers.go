package helpers

import (
	"errors"
	"fmt"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func WriteBadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": http.StatusBadRequest, "message": msg}})
}

func WriteNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": http.StatusNotFound, "message": msg}})
}

func WriteInternalServerError(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": http.StatusInternalServerError, "message": msg}})
}

func WriteCreated(c *gin.Context, payload any) {
	c.JSON(http.StatusCreated, gin.H{"response": payload})
}

func WriteOK(c *gin.Context, payload any) { c.JSON(http.StatusOK, gin.H{"response": payload}) }

func GetIdFromQuery(c *gin.Context) (int64, error) {
	queryId := c.Param("id")

	if queryId == "" {
		return 0, fmt.Errorf("%w: id is required", car_service.ErrInvalidInput)
	}

	id, err := strconv.Atoi(queryId)
	if err != nil {
		return 0, fmt.Errorf("%w: id must be a number", car_service.ErrInvalidInput)
	}
	return int64(id), nil
}

func HandleError(c *gin.Context, err error) {
	for target, status := range ErrorCodeMap {
		if errors.Is(err, target) {
			c.JSON(status, gin.H{"error": err.Error(), "code": status})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops.. Something went wrong", "code": http.StatusBadRequest})
}
