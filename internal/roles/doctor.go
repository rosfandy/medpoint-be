package roles

import (
	"github.com/sev-2/raiden"
)

type Doctor struct {
	raiden.RoleBase
}

func (r *Doctor) Name() string {
	return "doctor"
}

