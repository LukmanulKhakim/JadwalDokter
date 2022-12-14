package delivery

import (
	"jadwaldokter/features/jadwal/domain"
	"jadwaldokter/utils/helper"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type jadwalHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := jadwalHandler{srv: srv}
	e.POST("/jadwal", handler.AddJadwal())
	e.GET("/jadwal", handler.GetJadwal())
	e.DELETE("/jadwal/:id", handler.DeleteJadwal())
}

func (jh *jadwalHandler) AddJadwal() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		if strings.TrimSpace(input.Hari) == "" && strings.TrimSpace(input.Jam) == "" && input.Dokter_ID == 0 && input.Poli_ID == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("input empty"))
		}
		cnv := ToDomain(input)
		res, err := jh.srv.AddJadwal(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("sucses add jadwal", ToResponse(res, "reg")))
	}
}

func (jh *jadwalHandler) GetJadwal() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.QueryParam("key")
		if strings.ContainsAny(key, "11-02-2022") {
			t, err := time.Parse("01-02-2006", key)
			if err != nil {
				log.Println(err.Error())
			}
			keyCon := t.Weekday().String()
			key = helper.Convert(keyCon)
			log.Println(key)
		} else {
			key = c.QueryParam("key")
			log.Println(key)
		}
		res, err := jh.srv.GetJadwal(key)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			log.Println(key)
			return c.JSON(http.StatusOK, SuccessResponse("success get jadwal", ToResponse(res, "all")))
		}
		return nil
	}
}

func (jh *jadwalHandler) DeleteJadwal() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id poli must integer"))
		}

		res, err := jh.srv.DeleteJadwal(uint(ID))
		log.Println("res data :", res)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusBadRequest, FailResponse("not found"))
			} else if strings.Contains(err.Error(), "database") {
				return c.JSON(http.StatusBadRequest, FailResponse("not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			return c.JSON(http.StatusAccepted, SuccessDeleteResponse("Success delete jadwal"))
		}

	}
}
