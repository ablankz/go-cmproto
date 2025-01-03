package cmproto_test

import (
	"testing"

	flex "buf.build/gen/go/cresplanex/types/protocolbuffers/go/cresplanex/flex/v1"
	nullable "buf.build/gen/go/cresplanex/types/protocolbuffers/go/cresplanex/nullable/v1"

	"github.com/ablankz/go-cmproto"
	"github.com/stretchr/testify/assert"
)

func TestFromFlex(t *testing.T) {
	tests := []struct {
		name     string
		input    *flex.Flex
		expected any
	}{
		{
			name:     "Nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "StringValue",
			input:    &flex.Flex{Flex: &flex.Flex_StringValue{StringValue: "test"}},
			expected: "test",
		},
		{
			name:     "IntValue",
			input:    &flex.Flex{Flex: &flex.Flex_IntValue{IntValue: 42}},
			expected: int32(42),
		},
		{
			name:     "ArrayValue",
			input:    &flex.Flex{Flex: &flex.Flex_ArrayValue{ArrayValue: &flex.FlexArray{Flex: []*flex.Flex{{Flex: &flex.Flex_StringValue{StringValue: "item1"}}}}}},
			expected: []any{"item1"},
		},
		{
			name:     "NullableStringValue with value",
			input:    &flex.Flex{Flex: &flex.Flex_NullableStringValue{NullableStringValue: &nullable.NullableString{HasValue: true, Value: "test"}}},
			expected: "test",
		},
		{
			name:     "NullableStringValue without value",
			input:    &flex.Flex{Flex: &flex.Flex_NullableStringValue{NullableStringValue: &nullable.NullableString{HasValue: false}}},
			expected: nil,
		},
		{
			name: "Complex",
			input: &flex.Flex{Flex: &flex.Flex_MapValue{MapValue: &flex.FlexMap{
				Flex: map[string]*flex.Flex{
					"key1": {Flex: &flex.Flex_StringValue{StringValue: "value1"}},
					"key2": {Flex: &flex.Flex_ArrayValue{ArrayValue: &flex.FlexArray{Flex: []*flex.Flex{{Flex: &flex.Flex_StringValue{StringValue: "value2"}}}}}},
				},
			}}},
			expected: map[string]any{
				"key1": "value1",
				"key2": []any{"value2"},
			},
		},
		{
			name: "Complex2",
			input: &flex.Flex{Flex: &flex.Flex_MapValue{MapValue: &flex.FlexMap{
				Flex: map[string]*flex.Flex{
					"MetricsBreakTime":     {Flex: &flex.Flex_StringValue{StringValue: "10m"}},
					"MetricsInterval":      {Flex: &flex.Flex_StringValue{StringValue: "5s"}},
					"RequestPerSlaveCount": {Flex: &flex.Flex_IntValue{IntValue: 3000}},
					"SlaveCount":           {Flex: &flex.Flex_IntValue{IntValue: 2}},
					"ThreadPerSlaveCount":  {Flex: &flex.Flex_IntValue{IntValue: 50}},
					"slaveLists": {Flex: &flex.Flex_ArrayValue{ArrayValue: &flex.FlexArray{Flex: []*flex.Flex{
						{Flex: &flex.Flex_MapValue{MapValue: &flex.FlexMap{
							Flex: map[string]*flex.Flex{
								"address": {Flex: &flex.Flex_StringValue{StringValue: "127.0.0.1"}},
								"id":      {Flex: &flex.Flex_StringValue{StringValue: "slave1"}},
								"port":    {Flex: &flex.Flex_IntValue{IntValue: 50051}},
								"uuid":    {Flex: &flex.Flex_StringValue{StringValue: "b4a9297f-1651-4b37-9c2e-6d6223cd72ee"}},
							},
						}}},
						{Flex: &flex.Flex_MapValue{MapValue: &flex.FlexMap{
							Flex: map[string]*flex.Flex{
								"address": {Flex: &flex.Flex_StringValue{StringValue: "127.0.0.2"}},
								"id":      {Flex: &flex.Flex_StringValue{StringValue: "slave2"}},
								"port":    {Flex: &flex.Flex_IntValue{IntValue: 50051}},
								"uuid":    {Flex: &flex.Flex_StringValue{StringValue: "69a82f87-99ab-49d6-a9d6-f4948fd10d09"}},
							},
						}}},
					}}}},
				},
			}}},
			expected: map[string]any{
				"MetricsBreakTime":     "10m",
				"MetricsInterval":      "5s",
				"RequestPerSlaveCount": int32(3000),
				"SlaveCount":           int32(2),
				"ThreadPerSlaveCount":  int32(50),
				"slaveLists": []any{
					map[string]any{
						"address": "127.0.0.1",
						"id":      "slave1",
						"port":    int32(50051),
						"uuid":    "b4a9297f-1651-4b37-9c2e-6d6223cd72ee",
					},
					map[string]any{
						"address": "127.0.0.2",
						"id":      "slave2",
						"port":    int32(50051),
						"uuid":    "69a82f87-99ab-49d6-a9d6-f4948fd10d09",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cmproto.FromFlex(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFlex(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{
			name:  "String input",
			value: "test",
		},
		{
			name:  "Int32 input",
			value: int32(42),
		},
		{
			name:  "Nil input",
			value: nil,
		},
		{
			name:  "Array input",
			value: []any{"item1"},
		},
		{
			name:  "Map input",
			value: map[string]any{"key1": "value1", "key2": []any{"value2"}},
		},
		{
			name: "Complex",
			value: map[string]any{
				"MetricsBreakTime":     "10m",
				"MetricsInterval":      "5s",
				"RequestPerSlaveCount": int32(3000),
				"SlaveCount":           int32(2),
				"ThreadPerSlaveCount":  int32(50),
				"slaveLists": []any{
					map[string]any{
						"address": "127.0.0.1",
						"id":      "slave1",
						"port":    int32(50051),
						"uuid":    "b4a9297f-1651-4b37-9c2e-6d6223cd72ee",
					},
					map[string]any{
						"address": "127.0.0.2",
						"id":      "slave2",
						"port":    int32(50051),
						"uuid":    "69a82f87-99ab-49d6-a9d6-f4948fd10d09",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flex, err := cmproto.NewFlex(tt.value)
			assert.NoError(t, err)
			result, err := cmproto.FromFlex(flex)
			assert.NoError(t, err)
			assert.Equal(t, tt.value, result)
		})
	}
}
