package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type DoctorSchedules struct {
	db.ModelBase
	Id          int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	DoctorId    *uuid.UUID `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable;default:gen_random_uuid()"`
	DayOfWeek   *string    `json:"day_of_week,omitempty" column:"name:day_of_week;type:varchar;nullable"`
	StartTime   *time.Time `json:"start_time,omitempty" column:"name:start_time;type:timez;nullable"`
	EndTime     *time.Time `json:"end_time,omitempty" column:"name:end_time;type:timez;nullable"`
	IsAvailable bool       `json:"is_available,omitempty" column:"name:is_available;type:boolean;nullable:false;default:false"`
	CreatedAt   time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor_schedules" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor *Doctors `json:"doctor,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
}
