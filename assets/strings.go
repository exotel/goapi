package assets

type strings struct {
	InvalidAction             string
	MandatoryFieldUnavailable string
	UnsupportedMethod         string
	PackageInconsistancy      string
	HTTPRequestError          string
	MethodNotSpecified        string
}

//String has all the constant strings defined
var String strings

//Initialize all the strings here
func init() {
	String.InvalidAction = "The action is not chosen or it is wrong"
	String.MandatoryFieldUnavailable = "Mandatory Field %s is empty"
	String.PackageInconsistancy = "PACKAGE-INCONSISTANCY, REPORT DEVELOPER (sarath@exotel.in) "
	String.HTTPRequestError = "Error occured while making http request : %s"
	String.UnsupportedMethod = "The method %s is not supported for the resource"
	String.MethodNotSpecified = "Method is not specified"
}

//Methods is a contant struct instance storing HTTP methods
var Methods = struct {
	Get, Post, Delete, Put string
}{"GET", "POST", "DELETE", "PUT"}
