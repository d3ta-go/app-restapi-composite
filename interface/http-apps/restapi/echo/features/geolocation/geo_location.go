package geolocation

import (
	"net/http"

	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
	appGeoLoc "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/application"
	appGeoLocDTOCountry "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/application/dto/country"
)

// NewFGeoLocation new FGeoLocation
func NewFGeoLocation(h *handler.Handler) (*FGeoLocation, error) {
	var err error

	f := new(FGeoLocation)
	f.SetHandler(h)

	if f.appGeoLocation, err = appGeoLoc.NewGeoLocationApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FGeoLocation represent FGeoLocation
type FGeoLocation struct {
	features.BaseFeature
	appGeoLocation *appGeoLoc.GeoLocationApp
}

// ListAllCountry list AllCountry
func (f *FGeoLocation) ListAllCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	resp, err := f.appGeoLocation.CountrySvc.ListAll(i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// RefreshCountryIndexer refresh Country Indexer
func (f *FGeoLocation) RefreshCountryIndexer(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTOCountry.RefreshCountryIndexerReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.RefreshIndexer(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// SearchCountryIndexer search Country Indexer
func (f *FGeoLocation) SearchCountryIndexer(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTOCountry.SearchCountryIndexerReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.SearchIndexer(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// GetCountry get Country
func (f *FGeoLocation) GetCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appGeoLocDTOCountry.GetCountryReqDTO)
	req.Code = code

	resp, err := f.appGeoLocation.CountrySvc.GetDetail(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// AddCountry add Country
func (f *FGeoLocation) AddCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appGeoLocDTOCountry.AddCountryReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appGeoLocation.CountrySvc.Add(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// UpdateCountry update Country
func (f *FGeoLocation) UpdateCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	reqKeys := new(appGeoLocDTOCountry.UpdateCountryKeysDTO)
	reqKeys.Code = code

	reqData := new(appGeoLocDTOCountry.UpdateCountryDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	req := new(appGeoLocDTOCountry.UpdateCountryReqDTO)
	req.Keys = reqKeys
	req.Data = reqData

	resp, err := f.appGeoLocation.CountrySvc.Update(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// DeleteCountry delete Country
func (f *FGeoLocation) DeleteCountry(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appGeoLocDTOCountry.DeleteCountryReqDTO)
	req.Code = code

	resp, err := f.appGeoLocation.CountrySvc.Delete(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}
