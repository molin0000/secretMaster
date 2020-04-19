package interact

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

type Interact struct{}

func (i *Interact) IsGroupAdmin(group, qq uint64) bool {
	member := cqp.GetGroupMemberInfo(int64(group), int64(qq), false)
	return member.Auth > 1
}

func (i *Interact) IsGroupMaster(group, qq uint64) bool {
	member := cqp.GetGroupMemberInfo(int64(group), int64(qq), false)
	return member.Auth > 2
}
