package cmproto

import (
	flex "buf.build/gen/go/cresplanex/types/protocolbuffers/go/cresplanex/flex/v1"
	nullable "buf.build/gen/go/cresplanex/types/protocolbuffers/go/cresplanex/nullable/v1"
)

func FromFlex(f *flex.Flex) (any, error) {
	if f == nil {
		return nil, nil
	}

	switch f.WhichFlex() {
	case flex.Flex_StringValue_case:
		return f.GetStringValue(), nil
	case flex.Flex_IntValue_case:
		return f.GetIntValue(), nil
	case flex.Flex_LongValue_case:
		return f.GetLongValue(), nil
	case flex.Flex_FloatValue_case:
		return f.GetFloatValue(), nil
	case flex.Flex_DoubleValue_case:
		return f.GetDoubleValue(), nil
	case flex.Flex_BoolValue_case:
		return f.GetBoolValue(), nil
	case flex.Flex_BytesValue_case:
		return f.GetBytesValue(), nil
	case flex.Flex_FlexValue_case:
		return FromFlex(f.GetFlexValue())
	case flex.Flex_ArrayValue_case:
		var result []any
		for _, f := range f.GetArrayValue().Flex {
			item, err := FromFlex(f)
			if err != nil {
				return nil, err
			}
			result = append(result, item)
		}
		return result, nil
	case flex.Flex_MapValue_case:
		result := make(map[string]any)
		for k, v := range f.GetMapValue().Flex {
			value, err := FromFlex(v)
			if err != nil {
				return nil, err
			}
			result[k] = value
		}
		return result, nil
	case flex.Flex_NullableStringValue_case:
		if f.GetNullableStringValue().GetHasValue() {
			return f.GetNullableStringValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableIntValue_case:
		if f.GetNullableIntValue().GetHasValue() {
			return f.GetNullableIntValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableLongValue_case:
		if f.GetNullableLongValue().GetHasValue() {
			return f.GetNullableLongValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableFloatValue_case:
		if f.GetNullableFloatValue().GetHasValue() {
			return f.GetNullableFloatValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableDoubleValue_case:
		if f.GetNullableDoubleValue().GetHasValue() {
			return f.GetNullableDoubleValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableBoolValue_case:
		if f.GetNullableBoolValue().GetHasValue() {
			return f.GetNullableBoolValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableBytesValue_case:
		if f.GetNullableBytesValue().GetHasValue() {
			return f.GetNullableBytesValue().GetValue(), nil
		}
		return nil, nil
	case flex.Flex_NullableFlexValue_case:
		if f.GetNullableFlexValue().GetHasValue() {
			return FromFlex(f.GetNullableFlexValue().Get_Value())
		}
		return nil, nil
	case flex.Flex_NullableArrayValue_case:
		if f.GetNullableArrayValue().GetHasValue() {
			var result []any
			for _, f := range f.GetNullableArrayValue().Get_Value().Flex {
				item, err := FromFlex(f)
				if err != nil {
					return nil, err
				}
				result = append(result, item)
			}
			return result, nil
		}
		return nil, nil
	case flex.Flex_NullableMapValue_case:
		if f.GetNullableMapValue().GetHasValue() {
			result := make(map[string]any)
			for k, v := range f.GetNullableMapValue().Get_Value().Flex {
				value, err := FromFlex(v)
				if err != nil {
					return nil, err
				}
				result[k] = value
			}
			return result, nil
		}
		return nil, nil
	default:
		return nil, nil
	}
}

func FromFlexArray(f *flex.FlexArray) ([]any, error) {
	var result []any
	for _, f := range f.Flex {
		item, err := FromFlex(f)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func FromFlexMap(f *flex.FlexMap) (map[string]any, error) {
	result := make(map[string]any)
	for k, v := range f.Flex {
		value, err := FromFlex(v)
		if err != nil {
			return nil, err
		}
		result[k] = value
	}
	return result, nil
}

func FromNullableFlex(f *flex.NullableFlex) (any, error) {
	if f.GetHasValue() {
		return FromFlex(f.Get_Value())
	}
	return nil, nil
}

func FromNullableFlexArray(f *flex.NullableFlexArray) ([]any, error) {
	if f.GetHasValue() {
		return FromFlexArray(f.Get_Value())
	}
	return nil, nil
}

func FromNullableFlexMap(f *flex.NullableFlexMap) (map[string]any, error) {
	if f.GetHasValue() {
		return FromFlexMap(f.Get_Value())
	}
	return nil, nil
}

func NewFlex(obj any) (*flex.Flex, error) {
	flexBuilder := &flex.Flex{}
	switch v := obj.(type) {
	case nil:
		flexBuilder.SetNullableFlexValue(&flex.NullableFlex{HasValue: false})
	case string:
		flexBuilder.SetNullableStringValue(&nullable.NullableString{HasValue: true, Value: v})
	case int32:
		flexBuilder.SetNullableIntValue(&nullable.NullableInt32{HasValue: true, Value: v})
	case int64:
		flexBuilder.SetNullableLongValue(&nullable.NullableInt64{HasValue: true, Value: v})
	case float32:
		flexBuilder.SetNullableFloatValue(&nullable.NullableFloat{HasValue: true, Value: v})
	case float64:
		flexBuilder.SetNullableDoubleValue(&nullable.NullableDouble{HasValue: true, Value: v})
	case bool:
		flexBuilder.SetNullableBoolValue(&nullable.NullableBool{HasValue: true, Value: v})
	case []byte:
		flexBuilder.SetNullableBytesValue(&nullable.NullableBytes{HasValue: true, Value: v})
	case *flex.Flex:
		f, err := NewNullableFlex(v)
		if err != nil {
			return nil, err
		}
		flexBuilder.SetNullableFlexValue(f)
	case []any:
		arrayBuilder := &flex.FlexArray{}
		for _, val := range v {
			f, err := NewFlex(val)
			if err != nil {
				return nil, err
			}
			arrayBuilder.Flex = append(arrayBuilder.Flex, f)
		}
		flexBuilder.SetNullableArrayValue(&flex.NullableFlexArray{HasValue: true, Value: arrayBuilder})
	case map[string]any:
		mapBuilder := &flex.FlexMap{
			Flex: make(map[string]*flex.Flex),
		}
		for k, val := range v {
			f, err := NewFlex(val)
			if err != nil {
				return nil, err
			}
			mapBuilder.Flex[k] = f
		}
		flexBuilder.SetNullableMapValue(&flex.NullableFlexMap{HasValue: true, Value: mapBuilder})
	default:
		f, err := NewNullableFlex(obj)
		if err != nil {
			return nil, err
		}
		flexBuilder.SetNullableFlexValue(f)
	}
	return flexBuilder, nil
}

func NewNullableFlex(obj any) (*flex.NullableFlex, error) {
	if obj == nil {
		return &flex.NullableFlex{HasValue: false}, nil
	}
	f, err := NewFlex(obj)
	if err != nil {
		return nil, err
	}
	return &flex.NullableFlex{HasValue: true, Value: f}, nil
}

func NewFlexArray(obj []any) (*flex.FlexArray, error) {
	arrayBuilder := &flex.FlexArray{}
	for _, val := range obj {
		f, err := NewFlex(val)
		if err != nil {
			return nil, err
		}
		arrayBuilder.Flex = append(arrayBuilder.Flex, f)
	}
	return arrayBuilder, nil
}

func NewFlexMap(obj map[string]any) (*flex.FlexMap, error) {
	mapBuilder := &flex.FlexMap{
		Flex: make(map[string]*flex.Flex),
	}
	for k, val := range obj {
		f, err := NewFlex(val)
		if err != nil {
			return nil, err
		}
		mapBuilder.Flex[k] = f
	}
	return mapBuilder, nil
}

func NewNullableFlexArray(obj []any) (*flex.NullableFlexArray, error) {
	arrayBuilder, err := NewFlexArray(obj)
	if err != nil {
		return nil, err
	}
	return &flex.NullableFlexArray{HasValue: true, Value: arrayBuilder}, nil
}

func NewNullableFlexMap(obj map[string]any) (*flex.NullableFlexMap, error) {
	mapBuilder, err := NewFlexMap(obj)
	if err != nil {
		return nil, err
	}
	return &flex.NullableFlexMap{HasValue: true, Value: mapBuilder}, nil
}
