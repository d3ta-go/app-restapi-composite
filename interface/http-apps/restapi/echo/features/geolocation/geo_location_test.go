package geolocation

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ht "github.com/d3ta-go/app-restapi-composite/interface/http-apps/restapi/echo/features/helper_test"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFGeoLocation_ListAllCountry(t *testing.T) {
	// handler
	h := ht.NewHandler()

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/geolocation/countries/list-all", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.ListAllCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		t.Logf("RESPONSE.geoLoc.ListAllCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_RefreshCountryIndexer(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.refresh-indexer.request")

	// variables
	reqDTO := `{
		"processType":"` + testData["processing-type"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/countries/indexer/refresh", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.RefreshCountryIndexer(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.refresh-indexer.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.RefreshCountryIndexer: %s", res.Body.String())
	}
}

func TestFGeoLocation_SearchCountryIndexer(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.search-indexer.request")

	// variables
	reqDTO := `{
		"name":"` + testData["name"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/countries/indexer/search", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.SearchCountryIndexer(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.search-indexer.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.SearchCountryIndexer: %s", res.Body.String())
	}
}

func TestFGeoLocation_AddCountry(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.add.request")

	// variables
	reqDTO := `{
		"code":"` + testData["code"] + `", 
		"name": "` + testData["name"] + `", 
		"ISO2Code": "` + testData["iso2-code"] + `", 
		"ISO3Code": "` + testData["iso3-code"] + `", 
		"WHORegion": "` + testData["who-region"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/geolocation/country", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.AddCountry(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.add.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.AddCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_GetCountry(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.get-detail.request")

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/geolocation/country/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues(testData["code"])

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.GetCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.get-detail.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.GetCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_UpdateCountry(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.update.request")

	// variables
	reqDTO := `{
		"name": "` + testData["name"] + `",
		"ISO2Code": "` + testData["iso2-code"] + `",
		"ISO3Code": "` + testData["iso3-code"] + `",
		"WHORegion": "` + testData["who-region"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/geolocation/country/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues(testData["code"])

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.UpdateCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.update.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.UpdateCountry: %s", res.Body.String())
	}
}

func TestFGeoLocation_DeleteCountry(t *testing.T) {
	// handler
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.interface-layer.features.delete.request")

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/geolocation/country/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	// c.SetPath("/api/v1/geolocation/country/:code")
	c.SetParamNames("code")
	c.SetParamValues(testData["code"])

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		t.Errorf("initialize.OpenAllIndexerConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	geoLoc, err := NewFGeoLocation(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, geoLoc.DeleteCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.interface-layer.features.delete.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.geoLoc.DeleteCountry: %s", res.Body.String())
	}
}
