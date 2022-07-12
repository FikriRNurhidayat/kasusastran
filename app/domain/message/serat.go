package message

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
)

type Serat Message[*entity.Serat]
type Serats Message[[]*entity.Serat]
