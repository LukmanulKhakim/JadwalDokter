package delivery

import (
	"jadwaldokter/features/dokter/domain"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type dokterHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := dokterHandler{srv: srv}
	e.POST("/dokter", handler.AddDokter())
	e.GET("/dokter", handler.GetDokter())
	e.DELETE("/dokter/:id", handler.DeleteDokter())

}

func (dh *dokterHandler) AddDokter() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		//IdPoli := strconv.Itoa(int(input.Poli_ID))
		if strings.TrimSpace(input.Nama_dokter) == "" && input.Poli_ID == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("input empty"))
		}
		cnv := ToDomain(input)
		res, err := dh.srv.AddDokter(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("sucses add dokter", ToResponse(res, "reg")))
	}
}

func (dh *dokterHandler) GetDokter() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := dh.srv.GetDokter()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			return c.JSON(http.StatusOK, SuccessResponse("success get all poli", ToResponse(res, "all")))
		}
		return nil
	}
}

func (dh *dokterHandler) DeleteDokter() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id poli must integer"))
		}

		res, err := dh.srv.DeleteDokter(uint(ID))
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
			return c.JSON(http.StatusAccepted, SuccessDeleteResponse("Success delete dokter"))
		}

	}
}
