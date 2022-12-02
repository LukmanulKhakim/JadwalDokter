package delivery

import (
	"jadwaldokter/features/poli/domain"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type poliHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := poliHandler{srv: srv}
	e.POST("/poli", handler.AddPoli())
	e.GET("/poli", handler.GetAllPoli())
	e.DELETE("/poli/:id", handler.DeletePoli())
}

func (ph *poliHandler) AddPoli() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		if strings.TrimSpace(input.Nama_poli) == "" {
			return c.JSON(http.StatusBadRequest, FailResponse("input empty"))
		}
		cnv := ToDomain(input)
		res, err := ph.srv.AddPoli(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("sucses add poli", ToResponse(res, "reg")))
	}
}

func (ph *poliHandler) GetAllPoli() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.GetAllPoli()
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

func (ph *poliHandler) DeletePoli() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id poli must integer"))
		}

		res, err := ph.srv.DeletePoli(uint(ID))
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
			return c.JSON(http.StatusAccepted, SuccessDeleteResponse("Success delete poli"))
		}

	}
}
