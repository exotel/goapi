package api

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/exotel/goapi/assets"
	"github.com/exotel/goapi/helpers"
)

//IsValidAction Checks if there are enough details for the execution
func (c Client) IsValidAction() (ok bool, err error) {
	ok = c.action&c.validActions != 0x00
	if !ok {
		err = fmt.Errorf("The method %s is not supported for the resource", c.action)
	}
	return
}

//IsValidData checks if the data is valid or not
//It uses the tag `mandatory` to see if a field is mandatory or not
func (c Client) IsValidData() (ok bool, err error) {
	v := reflect.ValueOf(c.data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	//If not a struct its a problem now every data has to be a Struct
	if v.Kind() != reflect.Struct {
		err = fmt.Errorf(assets.String.PackageInconsistancy+" [only accepts structs; got (%+[1]v)%[1]T]", v)
		return
	}

	//Now pass through all the values and make sure all the mandatory fields are set
	tp := v.Type()
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		emptyValue := helpers.IsEmptyValue(v.Field(i))
		if emptyValue {
			if tagv := field.Tag.Get("mandatory"); tagv == "true" {
				err = fmt.Errorf(assets.String.MandatoryFieldUnavailable, field.Name)
			}
		}
	}
	ok = true
	return
}

//setURL sets the url for the request according to the data given
func (c *Client) setURL() (err error) {
	var parsedTemplate *template.Template
	parsedTemplate, err = template.New("url").Parse(c.url)
	if err != nil {
		return
	}
	//create the hash map for filling template from client and data
	var urlParamsMap map[string]interface{}
	urlParamsMap, err = helpers.StructToMap("url", c, c.data)
	if err != nil {
		return
	}
	//fill the template with proper values
	var b bytes.Buffer
	err = parsedTemplate.Execute(&b, urlParamsMap)
	if err != nil {
		return
	}
	c.url = string(b.Bytes())
	return
}

//Do does the actual job
func (c *Client) Do() (result interface{}, err error) {
	var ok bool
	//Check if the action requested is valid or not
	if ok, err = c.IsValidAction(); !ok {
		return
	}
	//Checking if the data is valid
	if ok, err = c.IsValidData(); !ok {
		return
	}
	//TODO: setup url using may be the text/template library
	//forming the url to make request
	fmt.Println("Setting the url")
	err = c.setURL()
	if err != nil {
		return
	}
	url := c.baseURL + c.url
	result, err = helpers.MakeHTTPRequest(url, c.Credentials, c.data, true)
	return
}
