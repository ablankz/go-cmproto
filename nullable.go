package cmproto

import nullable "buf.build/gen/go/cresplanex/types/protocolbuffers/go/cresplanex/nullable/v1"

func FromNullableString(nullableString *nullable.NullableString) *string {
	if nullableString.GetHasValue() {
		return &nullableString.Value
	}
	return nil
}

func FromNullableInt32(nullableInt32 *nullable.NullableInt32) *int32 {
	if nullableInt32.GetHasValue() {
		return &nullableInt32.Value
	}
	return nil
}

func FromNullableInt64(nullableInt64 *nullable.NullableInt64) *int64 {
	if nullableInt64.GetHasValue() {
		return &nullableInt64.Value
	}
	return nil
}

func FromNullableFloat(nullableFloat *nullable.NullableFloat) *float32 {
	if nullableFloat.GetHasValue() {
		return &nullableFloat.Value
	}
	return nil
}

func FromNullableDouble(nullableDouble *nullable.NullableDouble) *float64 {
	if nullableDouble.GetHasValue() {
		return &nullableDouble.Value
	}
	return nil
}

func FromNullableBool(nullableBool *nullable.NullableBool) *bool {
	if nullableBool.GetHasValue() {
		return &nullableBool.Value
	}
	return nil
}

func FromNullableBytes(nullableBytes *nullable.NullableBytes) *[]byte {
	if nullableBytes.GetHasValue() {
		return &nullableBytes.Value
	}
	return nil
}

func ToNullableString(val string) *nullable.NullableString {
	return &nullable.NullableString{HasValue: true, Value: val}
}

func ToNullableInt32(val int32) *nullable.NullableInt32 {
	return &nullable.NullableInt32{HasValue: true, Value: val}
}

func ToNullableInt64(val int64) *nullable.NullableInt64 {
	return &nullable.NullableInt64{HasValue: true, Value: val}
}

func ToNullableFloat(val float32) *nullable.NullableFloat {
	return &nullable.NullableFloat{HasValue: true, Value: val}
}

func ToNullableDouble(val float64) *nullable.NullableDouble {
	return &nullable.NullableDouble{HasValue: true, Value: val}
}

func ToNullableBool(val bool) *nullable.NullableBool {
	return &nullable.NullableBool{HasValue: true, Value: val}
}

func ToNullableBytes(val []byte) *nullable.NullableBytes {
	return &nullable.NullableBytes{HasValue: true, Value: val}
}

func ToNullableStringPtr(val *string) *nullable.NullableString {
	if val == nil {
		return &nullable.NullableString{HasValue: false}
	}
	return &nullable.NullableString{HasValue: true, Value: *val}
}

func ToNullableInt32Ptr(val *int32) *nullable.NullableInt32 {
	if val == nil {
		return &nullable.NullableInt32{HasValue: false}
	}
	return &nullable.NullableInt32{HasValue: true, Value: *val}
}

func ToNullableInt64Ptr(val *int64) *nullable.NullableInt64 {
	if val == nil {
		return &nullable.NullableInt64{HasValue: false}
	}
	return &nullable.NullableInt64{HasValue: true, Value: *val}
}

func ToNullableFloatPtr(val *float32) *nullable.NullableFloat {
	if val == nil {
		return &nullable.NullableFloat{HasValue: false}
	}
	return &nullable.NullableFloat{HasValue: true, Value: *val}
}

func ToNullableDoublePtr(val *float64) *nullable.NullableDouble {
	if val == nil {
		return &nullable.NullableDouble{HasValue: false}
	}
	return &nullable.NullableDouble{HasValue: true, Value: *val}
}

func ToNullableBoolPtr(val *bool) *nullable.NullableBool {
	if val == nil {
		return &nullable.NullableBool{HasValue: false}
	}
	return &nullable.NullableBool{HasValue: true, Value: *val}
}

func ToNullableBytesPtr(val *[]byte) *nullable.NullableBytes {
	if val == nil {
		return &nullable.NullableBytes{HasValue: false}
	}
	return &nullable.NullableBytes{HasValue: true, Value: *val}
}
