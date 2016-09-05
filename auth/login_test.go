package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

const (
	loginUrl    = "/auth/login"
	registerUrl = "/auth/register"
)

var u User
var respRecorder *httptest.ResponseRecorder

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a valid user$`, iHaveAValidUser)
	s.Step(`^I request "([^"]*)" "([^"]*)" with request body$`, iRequestWithRequestBody)
	s.Step(`^I get a (\d+) response$`, iGetAResponse)
	s.Step(`^I get a response body with a valid JWT$`, iGetAResponseBodyWithAValidJWT)
	s.Step(`^I request "([^"]*)" "([^"]*)"$`, iRequest)
}

func iHaveAValidUser() error {

	u.Username = uuid.NewV4().String()
	u.Password = "password"
	uJson, err := json.Marshal(u)
	if err != nil {
		return err
	}
	body := string(uJson)
	if err := requestApi(http.MethodPost, registerUrl, &body, registerHandler); err != nil {
		return err
	}

	return iGetAResponse(201)
}

func iRequestWithRequestBody(httpMethod, url string, body *gherkin.DocString) error {
	reqBody := body.Content
	reqBody = strings.Replace(reqBody, "{username}", u.Username, -1)
	reqBody = strings.Replace(reqBody, "{password}", u.Password, -1)

	return requestApi(httpMethod, url, &reqBody, LoginHandler)
}

func iRequest(httpMethod, url string) error {
	return requestApi(httpMethod, url, nil, LoginHandler)
}

func requestApi(httpMethod, resourcePattern string, body *string, handler func(http.ResponseWriter, *http.Request)) error {

	url := fmt.Sprintf("http://localhost:%d%s", port, resourcePattern)

	var b io.Reader
	if body != nil {
		b = bytes.NewReader([]byte(*body))
	}

	req := httptest.NewRequest(httpMethod, url, b)
	respRecorder = httptest.NewRecorder()
	handler(respRecorder, req)

	return nil
}

func iGetAResponse(eStatus int) error {
	aStatus := respRecorder.Result().StatusCode
	if eStatus != aStatus {
		b, _ := ioutil.ReadAll(respRecorder.Result().Body)
		aBody := string(b)
		return fmt.Errorf("\tExpected %d \n \t Found %d \n\n \t%d \n \t%s", eStatus, aStatus, respRecorder.Result().StatusCode, aBody)
	}
	return nil
}

func iGetAResponseBodyWithAValidJWT() error {
	b, _ := ioutil.ReadAll(respRecorder.Result().Body)
	aBody := string(b)

	aBearer := gjson.Get(aBody, "tokenType").String()
	aToken := gjson.Get(aBody, "accessToken").String()

	if "Bearer" != aBearer {
		return fmt.Errorf("\tExpected %d \n \t Found %d \n\n \t%d \n \t%s", "Bearer", aBearer, respRecorder.Result().StatusCode, aBody)
	}
	if err := verifyClaims(aToken); err != nil {
		return fmt.Errorf("invalid jwt %s %s", aToken, err.Error())
	}

	return nil
}
