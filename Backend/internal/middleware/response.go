package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ResponseHandler custom response handler to add "err" field
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  error
		code gcode.Code = gcode.CodeOK
		res             = r.GetHandlerResponse()
	)
	if err = r.GetError(); err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != 200 {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case 404:
				code = gcode.CodeNotFound
			case 403:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeInternalError
			}
		}
	}

	// Custom response structure
	type Response struct {
		Code    int         `json:"code"    dc:"Error code"`
		Message string      `json:"message" dc:"Error message"`
		Data    interface{} `json:"data"    dc:"Result data for certain request according to model definition"`
		Err     string      `json:"err"     dc:"Error details"`
	}

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	r.Response.WriteJson(Response{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Err:     errStr,
	})
}
