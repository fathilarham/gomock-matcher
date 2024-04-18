package gomockmatcher

import (
	"encoding/json"
	"reflect"
	"strings"
)

type structMatcher struct {
	fields   map[string]bool
	s        any
	hasError bool
	option   Option
}

func (s structMatcher) Matches(x interface{}) bool {
	for field := range s.fields {
		childFields := strings.Split(field, ".")

		var actualValue, expectedValue any
		actualValue = reflect.ValueOf(s.s).FieldByName(childFields[0]).Interface()
		expectedValue = reflect.ValueOf(x).FieldByName(childFields[0]).Interface()

		for i := 1; i < len(childFields); i++ {
			actualValue = reflect.ValueOf(actualValue).FieldByName(childFields[i]).Interface()
			expectedValue = reflect.ValueOf(expectedValue).FieldByName(childFields[i]).Interface()
		}

		if !reflect.DeepEqual(expectedValue, actualValue) {
			if s.option.BailError {
				return false
			}

			s.hasError = true
		}
	}

	return !s.hasError
}

func (s structMatcher) String() string {
	b, _ := json.Marshal(s.s)
	return string(b)
}

func (s structMatcher) Fields(fields []string) structMatcher {
	for _, field := range fields {
		s.fields[field] = true
	}

	return s
}

func New(s any) structMatcher {
	return structMatcher{
		fields: make(map[string]bool),
		s:      s,
	}
}
