package helpers

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/exotel/goapi/assets"
	"github.com/exotel/goapi/assets/types"
	"github.com/parnurzeal/gorequest"
)

//Exorequester makes actual requests to the p.URL
type Exorequester struct {
	LastError error
	mode      types.Mode
	requester *gorequest.SuperAgent
}

//NewExorequester creates a new instance of Exorequester
func NewExorequester() *Exorequester {
	return &Exorequester{requester: gorequest.New()}
}

//Debug set the client in debug mode
func (p *Exorequester) Debug(debug bool) *Exorequester {
	if debug {
		p.mode = types.DEBUG
		return p
	}
	p.mode = types.PRODUCTION
	return p
}

//SetAuth sets basic auth
func (p *Exorequester) SetAuth(username string, password string) *Exorequester {
	p.requester.SetBasicAuth(username, password)
	return p
}

//SetHeaders Sets needed headers
func (p *Exorequester) SetHeaders() *Exorequester {
	p.requester.
		Set("'User-Agent", "goapi-v1").
		Set("Accept-language", "es").
		Set("accept", "application/json")
	return p
}

//POST Makes post request to the end point
func (p *Exorequester) POST(url string) *Exorequester {
	p.requester.Post(url)
	return p
}

//GET Makes get request to the end point
func (p *Exorequester) GET(url string) *Exorequester {
	p.requester.Get(url)
	return p
}

//DELETE Makes Delete request to the end point
func (p *Exorequester) DELETE(url string) *Exorequester {
	p.requester.Delete(url)
	return p
}

//PUT Makes put request to the end point
func (p *Exorequester) PUT(url string) *Exorequester {
	p.requester.Put(url)
	return p
}

//Log writes log for the Client
func (p *Exorequester) Log(format string, vals ...interface{}) {
	if p.mode == types.DEBUG {
		fmt.Printf(format+"\n", vals...)
	}
	return
}

//Params sets the parameters for requestss
func (p *Exorequester) Params(requestData interface{}) *Exorequester {
	if p.LastError != nil {
		return p
	}
	data, err := StructToMap("queryparam", requestData)
	if err != nil {
		p.LastError = err
		return p
	}
	//var values ur.Values
	for key, value := range data {
		p.requester.QueryData.Set(key, reflect.ValueOf(value).String())
	}
	return p
}

//Do does the execution of request
func (p *Exorequester) Do() (status int, response []byte, err error) {
	if p.LastError != nil {
		err = p.LastError
		return
	}
	p.Log("\nusing the url : %s", p.requester.Url)
	p.Log("using the headers : %v", p.requester.Header)
	p.Log("using autherisation : %v", p.requester.BasicAuth)
	p.Log("using the query parameters : %v", p.requester.QueryData)

	resp, body, errs := p.requester.End()

	if len(errs) > 0 {
		err = fmt.Errorf(assets.String.HTTPRequestError, errs[0])
		return
	}
	if resp == nil {
		err = fmt.Errorf(assets.String.HTTPRequestError, "Unable to Connect")
		return
	}

	if resp != nil {
		status = resp.StatusCode
	}

	p.Log("received data %s", body)
	response = make([]byte, len(body))
	response = []byte(body)
	if err != nil {
		err = errors.New("Error occured decoding the response :" + err.Error())
	}
	return
}

//MakeHTTPRequest makes a request to exotel platform api for desire operation
func MakeHTTPRequest(url string, credentials types.Credentials, method types.Action, data interface{}, debug bool) (status int, resp []byte, err error) {
	requester := NewExorequester()
	switch method {
	case types.READ, types.BULKREAD:
		requester = requester.GET(url)
	case types.CREATE:
		requester = requester.POST(url)
	case types.DELETE:
		requester = requester.DELETE(url)
	case types.UPDATE:
		requester = requester.PUT(url)
	}

	status, resp, err = requester.
		SetAuth(credentials.UserName, credentials.AccessToken).
		SetHeaders().
		Debug(debug).
		Params(data).
		Do()
	return
}
