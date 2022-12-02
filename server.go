package main

import (
	"jadwaldokter/config"
	dDokter "jadwaldokter/features/dokter/delivery"
	rDokter "jadwaldokter/features/dokter/repository"
	sDokter "jadwaldokter/features/dokter/services"
	dJadwal "jadwaldokter/features/jadwal/delivery"
	rJadwal "jadwaldokter/features/jadwal/repository"
	sJadwal "jadwaldokter/features/jadwal/services"
	dPoli "jadwaldokter/features/poli/delivery"
	rPoli "jadwaldokter/features/poli/repository"
	sPoli "jadwaldokter/features/poli/services"
	"jadwaldokter/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdlPoli := rPoli.New(db)
	serPoli := sPoli.New(mdlPoli)
	dPoli.New(e, serPoli)

	mdlDokter := rDokter.New(db)
	serDokter := sDokter.New(mdlDokter)
	dDokter.New(e, serDokter)

	mdlJadwal := rJadwal.New(db)
	serJadwal := sJadwal.New(mdlJadwal)
	dJadwal.New(e, serJadwal)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
