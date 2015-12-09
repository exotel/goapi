package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

//Resource defines te structure of a Resource
//an object of resource ios passed to the template (text/template) for generating handlegenerator
//from the template along with the functions
type Resource struct {
	Resource                     string
	Create, Read, Update, Delete bool
}

var (
	resource = flag.String("r", "", "The resource in singular form")
	methods  = flag.String("m", "", "The methods for which the client has to be written [ -m CRUD ]")
)

func generateclient(resource Resource) (err error) {
	var data []byte
	templatePath := fmt.Sprintf("%s//src/github.com/exotel/goapi/clientgenerator/%s", os.Getenv("GOPATH"), "client.template")
	data, err = ioutil.ReadFile(templatePath)
	if err != nil {
		return
	}
	var hanlder = string(data)
	// Create a new template and parse the letter into it.
	t := template.
		Must(template.
		New("client").
		Funcs(template.FuncMap{"lower": strings.ToLower, "reciever": func(s string) string { return "__receiver_" + string(([]rune(s))[0]) }}).
		Parse(hanlder))

	//craete the template into a bytes Buffer
	var bW bytes.Buffer
	w := bufio.NewWriter(&bW)
	err = t.Execute(w, resource)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}

	// format the bytes from bufffer
	var bFormated []byte
	bFormated, err = format.Source(bW.Bytes())
	if err != nil {
		return
	}

	//write to a file
	var f *os.File
	f, err = os.Create(strings.ToLower(resource.Resource) + "_client.go")
	if err != nil {
		return
	}

	_, err = bytes.NewBuffer(bFormated).WriteTo(f)
	if err != nil {
		return
	}
	return
}

func createResource() (res Resource, err error) {
	flag.Parse()
	res.Resource = *resource
	for _, m := range []rune(*methods) {
		switch m {
		case 'C':
			res.Create = true
		case 'R':
			res.Read = true
		case 'U':
			res.Update = true
		case 'D':
			res.Delete = true
		}
	}
	return
}

func main() {
	res, err := createResource()
	if err != nil {
		log.Fatal("Error happened creating the resource information", err.Error())
	}
	if err = generateclient(res); err != nil {
		log.Fatalln("Error happened generating the client ", err.Error())
	}
}
