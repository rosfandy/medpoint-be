package context

import (
	"encoding/json"

	"github.com/sev-2/raiden"
)

type RaidenContext struct {
	raiden.Context
}

func (ctx *RaidenContext) SendJson(data interface{}, statusCode int) error {
	ctx.Context.RequestContext().Response.Header.SetContentType("application/json")
	byteData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	ctx.Context.RequestContext().Response.SetStatusCode(statusCode)
	ctx.Context.RequestContext().Response.SetBody(byteData)

	return nil
}

func (ctx *RaidenContext) ThrowError(err string, statusCode int) error {
	response := map[string]interface{}{
		"success": false,
		"error":   err,
	}

	ctx.SendJson(response, statusCode)
	return nil
}
