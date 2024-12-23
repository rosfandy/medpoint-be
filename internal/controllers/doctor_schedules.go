package controllers

import (
	"medpointbe/internal/context"
	"medpointbe/internal/models"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type SchedulesRequest struct {
	DoctorId  string `json:"doctor_id"`
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type SchedulesController struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/schedules" type:"custom"`
	Payload *SchedulesRequest
}

type SchedulesControllerIds struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/schedules/{id}" type:"custom"`
	Payload *SchedulesRequest
}

func (c *SchedulesController) Get(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	var schedules []models.DoctorSchedules
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Select([]string{"*"}).
		Get(&schedules)

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}

	response := map[string]any{
		"success": true,
		"data":    schedules,
	}

	return HttpResponse.SendJson(response, 200)
}

func (c *SchedulesController) Post(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	payload := models.DoctorSchedules{
		DoctorId:  c.Payload.DoctorId,
		DayOfWeek: c.Payload.DayOfWeek,
		StartTime: c.Payload.StartTime,
		EndTime:   c.Payload.EndTime,
	}

	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
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

func (c *SchedulesControllerIds) Patch(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(string)
	payload := models.DoctorSchedules{
		DoctorId:  c.Payload.DoctorId,
		DayOfWeek: c.Payload.DayOfWeek,
		StartTime: c.Payload.StartTime,
		EndTime:   c.Payload.EndTime,
	}

	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Eq("id", id).
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

func (c *SchedulesControllerIds) Delete(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(string)
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Eq("id", id).
		Delete()

	if err != nil {
		return HttpResponse.ThrowError(err.Error(), 500)
	}

	response := map[string]interface{}{
		"success": true,
	}

	return HttpResponse.SendJson(response, 200)
}
