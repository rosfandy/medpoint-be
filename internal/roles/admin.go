package roles

import (
	"github.com/sev-2/raiden"
)

type Admin struct {
	raiden.RoleBase
}

func (r *Admin) Name() string {
	return "admin"
}

