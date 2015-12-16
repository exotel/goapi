package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"text/template"

	"github.com/exotel/goapi/assets"
	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/helpers"
)

//IsValidAction Checks if there are enough details for the execution
//It checks for if an action is set and if set is it a valid action on that resources
//For example if a resource does not accept POST then action being POST would return error
func (c Client) IsValidAction() (ok bool, err error) {
	ok = c.action&c.validActions != 0x00
	if !ok {
		if c.action == types.Action(0) {
			err = fmt.Errorf(assets.String.MethodNotSpecified)
		}
		err = fmt.Errorf(assets.String.UnsupportedMethod, c.action)
	}
	return
}

//IsValidCredentials Checks if the available credentials are valid or not
//checks if credentials are provded ,it does not check if the provided credentials are valid or not
func (c Client) IsValidCredentials() (ok bool, err error) {
	ok = len(c.Credentials.UserName) != 0 && len(c.Credentials.AccessToken) != 0
	if !ok {
		err = errors.New("ACCESS_TOKEN and 	USERNAME are mandatory")
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
		fmt.Println(field.Name, " => ", v.Field(i).Interface())
		if emptyValue {
			fmt.Println("EMPTY VALUE", field.Name, " => ", v.Field(i).Interface())
			if tagv := field.Tag.Get("mandatory"); tagv == "true" {
				err = fmt.Errorf(assets.String.MandatoryFieldUnavailable, field.Name)
				return
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
func (c *Client) Do() (status int, result map[string]interface{}, err error) {
	var ok bool
	//Check if the credentials are provded
	if ok, err = c.IsValidCredentials(); !ok {
		return
	}

	//Check if the action requested is valid or not
	if ok, err = c.IsValidAction(); !ok {
		return
	}
	//Checking if the data is valid
	if ok, err = c.IsValidData(); !ok {
		return
	}

	//forming the url to make request
	err = c.setURL()
	if err != nil {
		return
	}
	url := c.baseURL + c.url
	var bresult []byte
	status, bresult, err = helpers.MakeHTTPRequest(url, c.Credentials, c.action, c.data, c.mode == types.DEBUG)
	if err != nil {
		return
	}
	result = make(map[string]interface{})
	err = json.Unmarshal(bresult, &result)
	return
}
