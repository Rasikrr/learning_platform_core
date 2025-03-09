package converters

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertToTimestampPb(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ConvertToTimePtr(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	out := t.AsTime()
	return &out
}

func ConvertToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	out := t.AsTime()
	return out
}
