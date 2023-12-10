package constant

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	SUCESS               = 200
	BADREQUEST           = 400
	UNAUTHORIZED         = 401
	FORBIDDEN            = 403
	NOTFOUND             = 404
	METHODNOTALLOWED     = 405
	INTERNALSERVERERROR  = 500
	NOTIMPLEMENTED       = 501
	BADGATEWAY           = 502
	SERVICEUNAVAILABLE   = 503
	GATEWAYTIMEOUT       = 504
	UNSUPPORTEDMEDIATYPE = 415
	UNPROCESSABLEENTITY  = 422
)

/************************************Errors***************************/
var (
	INTERNALSERVERERRORMESSAGE = "Internal server error"
	PROMPTQUESTIONERROR        = "Error while finding prompt questions:"
)

/*************************************Redis keys***********************/
var (
	METOAPI    = "https://api.open-meteo.com/v1/forecast"
	QUERYNAME1 = "longitude"
	QUERYNAME2 = "latitude"
)
var Test bool

/*********************************GENERIC FUNCTION FOR JSON FORMATTING**************************************/
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Httpmethod(req interface{}, url string, method string, resp interface{}, username, password string) error {

	client := &http.Client{}
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if username != "" && password != "" {
		request.SetBasicAuth(username, password)
	}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &resp)
	return nil
}
