package covid19

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

func TestFCovid19_DisplayCurrentDataByCountry(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.covid19.covid19.interface-layer.features.display-current-data-by-country.request")

	// variables
	reqDTO := `{"countryCode":"` + testData["country-code"] + `", "providers": [{"code":"` + testData["provider-code"] + `"}]}`
	// resDTO := `{"status":"OK","response":{"message":"Operation succeeded","data":{"status":"OK","data":null}},"serverInfo":{"serverTime":"2020-07-18T11:26:35.377625+07:00"}}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/covid19/current/by-country", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}
	if err := initialize.OpenAllCacheConnection(handler); err != nil {
		t.Errorf("initialize.OpenAllCacheConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	covid19, err := NewFCovid19(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, covid19.DisplayCurrentDataByCountry(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.covid19.covid19.interface-layer.features.display-current-data-by-country.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.covid19.DisplayCurrentDataByCountry: %s", res.Body.String())
	}
}
