package roles

import (
	"github.com/sev-2/raiden"
)

type User struct {
	raiden.RoleBase
}

func (r *User) Name() string {
	return "user"
}

