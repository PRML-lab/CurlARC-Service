package entity

import "reflect"

type UserId struct {
	value string
}

func NewUserId(uuid string) *UserId {
	userId := new(UserId)
	userId.value = uuid
	return userId
}

func (r *UserId) Value() string {
	return r.value
}

func (r *UserId) Equals(other *UserId) bool {
	return reflect.DeepEqual(r.value, other.value)
}
