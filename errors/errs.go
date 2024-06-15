package errs

import (
	"encoding/json"
	"net/http"
)

var MethodNotAllowed = map[string]string{
	"Method": "Not allowed",
}

var InvalidId = map[string]string{
	"error": "Invalid id",
}

var NotFound = map[string]string{
	"Error": "Not found",
}

var InvalidJson = map[string]string{
	"Error": "Invalid data",
}

var InternalServer = map[string]string{
	"Error": "Server error",
}


var Success = map[string]string{
    "message":"Success",
}

func ErrorHandle(w http.ResponseWriter, s int, m map[string]string) {
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(s)
}


