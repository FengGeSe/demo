package http

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	errors "demo/errors"
)

const (
	PostMethod = "POST"
	GetMethod  = "GET"

	JsonContentTyp = "application/json"
	FormContentTyp = "application/x-www-form-urlencoded"
	XmlContentTyp  = "application/xml"
)

// server option
func ErrorServerOption() httptransport.ServerOption {
	return httptransport.ServerErrorEncoder(MyErrorEncoder)
}
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		terr := errors.ParseError(err.Error())
		if http.StatusText(terr.Code) != "" {
			// http错误码
			w.WriteHeader(terr.Code)
		}
		w.Header().Set("Content-Type", JsonContentTyp)
		w.Write([]byte(terr.String()))
	}
}

// check access
func checkAccess(r *http.Request, method, contentTyp string) error {
	var code int
	if r.Method != method {
		code = http.StatusMethodNotAllowed
		return errors.NewError(code, http.StatusText(code))
	}
	if contentTyp != "" && r.Header.Get("Content-Type") != contentTyp {
		code = http.StatusUnsupportedMediaType
		return errors.NewError(code, http.StatusText(code))
	}

	return nil
}

func ParamsCheckAccess(r *http.Request) error {
	return checkAccess(r, GetMethod, "")
}

func JsonCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, JsonContentTyp)
}

func FormCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, FormContentTyp)
}

func XmlCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, XmlContentTyp)
}
