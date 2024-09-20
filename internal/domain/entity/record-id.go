package entity

import "reflect"

type RecordId struct {
	value string
}

func NewRecordId(uuid string) *RecordId {
	recordId := new(RecordId)
	recordId.value = uuid
	return recordId
}

func (r *RecordId) Value() string {
	return r.value
}

func (r *RecordId) Equals(other *RecordId) bool {
	return reflect.DeepEqual(r.value, other.value)
}
