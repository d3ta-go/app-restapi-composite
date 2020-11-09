package router

import (
	"github.com/d3ta-go/app-restapi-composite/interface/http-apps/restapi/echo/features/covid19"
	internalMiddleware "github.com/d3ta-go/system/interface/http-apps/restapi/echo/middleware"
	"github.com/labstack/echo/v4"
)

// SetCovid19 set Covid19 Router
func SetCovid19(eg *echo.Group, f *covid19.FCovid19) {

	gc := eg.Group("/covid19")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("/current/by-country", f.DisplayCurrentDataByCountry)
}
