package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Profiles struct {
	db.ModelBase
	Id         uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	RoleId     int64      `json:"role_id,omitempty" column:"name:role_id;type:bigint;nullable:false;default:'3'"`
	Name       *string    `json:"name,omitempty" column:"name:name;type:varchar;nullable"`
	CreatedAt  time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false;default:now()"`
	SupabaseId *uuid.UUID `json:"supabase_id,omitempty" column:"name:supabase_id;type:uuid;nullable"`
	Email      *string    `json:"email,omitempty" column:"name:email;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"profiles" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorProfiles []*Doctors `json:"doctor_profiles,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasMany;primaryKey:id;foreignKey:profile_id"`
	Role           *Roles     `json:"role,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:role_id"`
}
