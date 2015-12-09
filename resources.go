package api

//go:generate  clientgenerator -r Account -m CRU
//go:generate  clientgenerator -r IncomingPhoneNumber  -m CRUD
//go:generate  clientgenerator -r OutgoingCallerID  -m CRD
//go:generate  clientgenerator -r AvailablePhoneNumber -m R
//go:generate  clientgenerator -r Call  -m CR
//go:generate  clientgenerator -r Recording -m R
