package router

import (
	"github.com/d3ta-go/app-restapi-composite/interface/http-apps/restapi/echo/features/geolocation"
	internalMiddleware "github.com/d3ta-go/system/interface/http-apps/restapi/echo/middleware"
	"github.com/labstack/echo/v4"
)

// SetGeoLocation set GeoLocation Router
func SetGeoLocation(eg *echo.Group, f *geolocation.FGeoLocation) {

	gc := eg.Group("/geolocation")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.GET("/countries/list-all", f.ListAllCountry)
	gc.POST("/countries/indexer/refresh", f.RefreshCountryIndexer)
	gc.POST("/countries/indexer/search", f.SearchCountryIndexer)
	gc.GET("/country/:code", f.GetCountry)
	gc.POST("/country", f.AddCountry)
	gc.PUT("/country/:code", f.UpdateCountry)
	gc.DELETE("/country/:code", f.DeleteCountry)
}
