package controllers

import (
	"medpointbe/internal/context"
	"medpointbe/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type ScheduleRequest struct {
	DoctorId  *uuid.UUID `json:"doctor_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	DayOfWeek *string    `json:"day_of_week"`
}

type ScheduleController struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/schedules" type:"custom"`
	Payload *ScheduleRequest
	Model   models.DoctorSchedules
}

type ScheduleControllerIds struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/schedules/{id}" type:"custom"`
	Payload *ScheduleRequest
	Model   models.DoctorSchedules
}

func (c *ScheduleController) Get(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	var schedules []models.DoctorSchedules

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
func (c *ScheduleController) Post(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	payload := models.DoctorSchedules{DoctorId: c.Payload.DoctorId, StartTime: c.Payload.StartTime, EndTime: c.Payload.EndTime}
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provid	ed!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.DoctorSchedules{}).
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

func (c *ScheduleControllerIds) Patch(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(int64)

	payload := models.DoctorSchedules{StartTime: c.Payload.StartTime, EndTime: c.Payload.EndTime, DayOfWeek: c.Payload.DayOfWeek}
	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
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

func (c *ScheduleControllerIds) Delete(ctx raiden.Context) error {
	HttpResponse := &context.RaidenContext{Context: ctx}

	id := ctx.RequestContext().UserValue("id").(int64)

	token := string(ctx.RequestContext().Request.Header.Peek("Authorization"))

	if token == "" {
		return HttpResponse.ThrowError("No token provided!", 401)
	}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
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
