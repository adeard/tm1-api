package tm1

import (
	"fmt"
	"net/http"
	"time"
	"tm1-api/domain"

	"github.com/gin-gonic/gin"
)

type tm1Handler struct {
	tm1Service Service
}

func NewTm1Handler(v1 *gin.RouterGroup, tm1Service Service) {

	handler := &tm1Handler{tm1Service}

	v1.GET("map", handler.GetMap)
	v1.GET(":uri1/:uri2", handler.GetTm)
	v1.POST("post", handler.SendTm)
	v1.POST("post/ratest", handler.PostRaTest)
}

// @Summary Send Tm1 Data
// @Description Send Tm1 Data
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Response{}
// @Router /api/v1/send [post]
// @Tags TM1
func (h *tm1Handler) SendTm(c *gin.Context) {
	start := time.Now()
	input := domain.Tm1RequestData{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.tm1Service.SendTm(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	result := domain.Response{
		Data:        res,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get Map Html
// @Description Get Map Html
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Response{}
// @Router /api/v1/map [get]
// @Tags TM1
func (h *tm1Handler) GetMap(c *gin.Context) {
	input := domain.MapRequestData{}

	c.ShouldBind(&input)
	c.HTML(http.StatusOK, "map.tmpl", gin.H{
		"title": input.Title,
		"lat":   input.Lat,
		"lng":   input.Lng,
	})
}

// @Summary Send Tm1 Data
// @Description Send Tm1 Data
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Response{}
// @Router /api/v1/post/ratest [post]
// @Tags TM1
func (h *tm1Handler) PostRaTest(c *gin.Context) {
	start := time.Now()
	input := domain.Tm1RequestDynamicData{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.tm1Service.SendRaTest(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	result := domain.Response{
		Data:        res,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Send Tm1 Data
// @Description Send Tm1 Data
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Response{}
// @Router /api/v1/post/:uri1/:uri2 [post]
// @Tags TM1
func (h *tm1Handler) GetTm(c *gin.Context) {
	start := time.Now()
	uri1 := c.Param("uri1")
	uri2 := c.Param("uri2")
	queryParams := c.Request.URL.Query()
	queryString := queryParams.Encode()

	res, err := h.tm1Service.GetTm(uri1, uri2, queryString)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	result := domain.Response{
		Data:        res,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	}

	c.JSON(http.StatusOK, result)
}
