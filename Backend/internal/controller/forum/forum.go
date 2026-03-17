package forum

import (
	"TrangleAgent/api/forum"
)

type ControllerV1 struct{}

func NewV1() forum.IForumV1 {
	return &ControllerV1{}
}
