package controllers

import (
	"medpointbe/internal/context"
	"medpointbe/internal/models"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}

type DoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/doctors" type:"custom"`
	Payload *DoctorRequest
	Model   models.Doctors
}

type DoctorControllerIds struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/doctors/{id}" type:"custom"`
	Payload *DoctorRequest
	Model   models.Doctors
}

func (c *DoctorController) Get(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	var doctors []models.Doctors

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
		Select([]string{"*"}).
		Get(&doctors)

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}
	response := map[string]any{
		"success": true,
		"data":    doctors,
	}

	return HttpResponse.SendJson(response, 200)
}
func (c *DoctorController) Post(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := uuid.New()
	payload := models.Doctors{Id: id, Name: c.Payload.Name, Specialization: c.Payload.Specialization}
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provid	ed!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
		Insert(&payload, nil)

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}

	response := map[string]interface{}{
		"success": true,
		"data":    payload,
	}

	return HttpResponse.SendJson(response, 201)
}

func (c *DoctorControllerIds) Patch(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(string)
	uuid, _ := uuid.Parse(id)

	payload := models.Doctors{Id: uuid, Name: c.Payload.Name, Specialization: c.Payload.Specialization}
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
		Eq("id", uuid).
		Update(payload, nil)

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}

	response := map[string]interface{}{
		"success": true,
		"data":    payload,
	}

	return HttpResponse.SendJson(response, 200)
}

func (c *DoctorControllerIds) Delete(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(string)
	uuid, _ := uuid.Parse(id)

	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
		Eq("id", uuid).
		Delete()

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}

	response := map[string]interface{}{
		"success": true,
	}

	return HttpResponse.SendJson(response, 200)
}
