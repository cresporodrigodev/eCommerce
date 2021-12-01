package valueObject

import (
	"time"
)

type CreateAt struct {
	time time.Time
}

func NewCreateAt() CreateAt {
	return CreateAt{time: time.Now()}
}

func (c CreateAt) String() string {
	return c.time.String()
}

type UpdateAt struct {
	time time.Time
}

func NewUpdateAt() UpdateAt {
	return UpdateAt{time: time.Now()}
}

func (c UpdateAt) String() string {
	return c.time.String()
}
