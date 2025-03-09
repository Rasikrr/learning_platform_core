package converters

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertToTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ConvertFromTimestamp(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	out := t.AsTime()
	return &out
}
