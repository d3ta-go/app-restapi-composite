package covid19

import (
	"net/http"

	appCovid19 "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/application"
	appCovid19DTO "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/application/dto"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFCovid19 new FCovid19
func NewFCovid19(h *handler.Handler) (*FCovid19, error) {
	var err error

	f := new(FCovid19)
	f.SetHandler(h)

	if f.appCovid19, err = appCovid19.NewCovid19App(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FCovid19 represent FCovid19
type FCovid19 struct {
	features.BaseFeature
	appCovid19 *appCovid19.Covid19App
}

// DisplayCurrentDataByCountry display CurrentDataByCountry
func (f *FCovid19) DisplayCurrentDataByCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appCovid19DTO.DisplayCurrentDataByCountryReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appCovid19.CurrentSvc.DisplayCurrentDataByCountry(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}
