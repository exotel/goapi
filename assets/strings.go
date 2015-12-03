package assets

type strings struct {
	InvalidAction             string
	MandatoryFieldUnavailable string
	PackageInconsistancy      string
	HTTPRequestError          string
}

//Strings has all the constant strings defined
var String strings

//Initialize all the strings here
func init() {
	String.InvalidAction = "The action is not chosen or it is wrong"
	String.MandatoryFieldUnavailable = "Mandatory Field %s is empty"
	String.PackageInconsistancy = "PACKAGE-INCONSISTANCY, REPORT DEVELOPER (sarath@exotel.in) "
	String.HTTPRequestError = "Error occured while making http request : %s"
}
