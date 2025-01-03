package models

import (
	"time"

	"github.com/sev-2/raiden/pkg/db"
)

type Roles struct {
	db.ModelBase
	Id        int64     `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	RoleName  string    `json:"role_name,omitempty" column:"name:role_name;type:varchar;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"roles" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ProfileRoles []*Profiles `json:"profile_roles,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:role_id"`
	// UsersThroughProfilesRole []*Users    `json:"users_through_profiles_role,omitempty" join:"joinType:manyToMany;through:profiles;sourcePrimaryKey:id;sourceForeignKey:role_id;targetPrimaryKey:id;targetForeign:role_id"`
}
