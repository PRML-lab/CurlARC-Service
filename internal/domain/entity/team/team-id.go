package entity

import "reflect"

type TeamId struct {
	value string
}

func NewTeamId(uuid string) *TeamId {
	teamId := new(TeamId)
	teamId.value = uuid
	return teamId
}

func (r *TeamId) Value() string {
	return r.value
}

func (r *TeamId) Equals(other *TeamId) bool {
	return reflect.DeepEqual(r.value, other.value)
}
