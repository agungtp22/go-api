package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_restfull/model"
	"go_restfull/util/queries"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetNews(c *gin.Context) {

	var news []model.News

	//if err := model.DB.Find(&news).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Oops"})
	//	return
	//}

	result := model.DB.Raw(queries.QueryListNews()).Scan(&news)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "data": nil, "message": "Data Not Found"})
		return
	} else {
		c.JSON(200, gin.H{"status": true, "data": news, "message": "Success Get Data"})

	}
}

func GetDetailNews(c *gin.Context) {

	var news model.News

	id := c.Param("id")

	//if err := model.DB.First(&news, id).Error; err != nil {
	//	switch err {
	//	case gorm.ErrRecordNotFound:
	//		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "data": nil, "message": "Data Not Found"})
	//		return
	//	default:
	//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Internal Server Error"})
	//		return
	//
	//	}
	//}

	if err := model.DB.Where("id = ?", id).First(&news).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "data": nil, "message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Internal Server Error"})
			return

		}
	}

	c.JSON(200, gin.H{"status": true, "data": news, "message": "Success Get Data"})

}

func CreateNews(c *gin.Context) {
	var news model.News

	if err := c.ShouldBindBodyWithJSON(&news); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	result := model.DB.Create(&news)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": result.Error})
		return
	} else {
		c.JSON(200, gin.H{"status": true, "data": news, "message": "Success Create Data with ID "})

	}
}

func UpdateNews(c *gin.Context) {
	var news model.News

	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&news); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	//if model.DB.Model(&news).Where("id = ? ", id).Updates(&news).RowsAffected == 0 {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Failed Update Data"})
	//	return
	//}

	//if model.DB.Model(&news).Where("id = ? ", id).Update("title", &news.Title).RowsAffected == 0 {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Failed Update Data"})
	//	return
	//}

	//if model.DB.Raw(queries.QueryUpdateNews(id, news.Title)).RowsAffected == 0 {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Failed Update Data"})
	//	return
	//}

	result := model.DB.Exec(queries.QueryUpdateNews(id, news.Title))

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": result.Error})
		//	return
	} else {
		c.JSON(200, gin.H{"status": true, "data": news, "message": "Success Update Data"})
	}

}

func DeleteNews(c *gin.Context) {

	var news model.News

	var input model.Request

	//var inputMap = map[string]string{"id": "0", "title": "", "creator": ""}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input.Id, 10, 64)
	title := input.Title

	//titleMap, _ := inputMap["title"]

	fmt.Println(input)
	fmt.Println(title)
	fmt.Println(id)

	//if model.DB.Delete(&news, id).RowsAffected == 0 {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Failed Update Data"})
	//	return
	//}

	if model.DB.Where("title = ? ", title).Delete(&news).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "data": nil, "message": "Failed Update Data"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": nil, "message": "Success Delete Data"})

}
