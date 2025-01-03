package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id             uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid();unique"`
	ProfileId      *uuid.UUID `json:"profile_id,omitempty" column:"name:profile_id;type:uuid;nullable;unique"`
	Name           string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	Specialization string     `json:"specialization,omitempty" column:"name:specialization;type:varchar;nullable:false"`
	CreatedAt      time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorScheduleDoctors []*DoctorSchedules `json:"doctor_schedule_doctors,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	Profile               *Profiles          `json:"profile,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasOne;primaryKey:id;foreignKey:profile_id"`
}
