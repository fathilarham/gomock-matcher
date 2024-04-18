package gomockmatcher

import (
	"encoding/json"
	"reflect"
	"strings"
)

type structMatcher struct {
	data any

	checkedFields map[string]bool
	ignoredFields map[string]bool

	option Option
	err    error
}

// New creates a new gomockmatcher instance.
//
//		user := User{}
//		data := gomockmatcher.New(user)
//
//	 You can pass optional option by passing an Option struct.
func New(data any, option ...Option) structMatcher {
	return structMatcher{
		data: data,
	}
}

// Matches implements gomock Matcher interface.
func (s structMatcher) Matches(x interface{}) bool {
	for field := range s.checkedFields {
		childFields := strings.Split(field, ".")

		var actualValue, expectedValue any
		actualValue = reflect.ValueOf(s.data).FieldByName(childFields[0]).Interface()
		expectedValue = reflect.ValueOf(x).FieldByName(childFields[0]).Interface()

		for i := 1; i < len(childFields); i++ {
			actualValue = reflect.ValueOf(actualValue).FieldByName(childFields[i]).Interface()
			expectedValue = reflect.ValueOf(expectedValue).FieldByName(childFields[i]).Interface()
		}

		if !reflect.DeepEqual(expectedValue, actualValue) {
			return false
		}
	}

	return true
}

// String implements gomock Matcher interface.
func (s structMatcher) String() string {
	b, _ := json.Marshal(s.data)
	return string(b)
}

func (s structMatcher) Include(fields []string) structMatcher {
	if s.isIgnoreUsed() {
		s.err = ErrIgnoreMethodAlreadyUsed
	}

	s.checkedFields = make(map[string]bool)

	for _, field := range fields {
		s.checkedFields[field] = true
	}

	return s
}

func (s structMatcher) Ignore(fields []string) structMatcher {
	if s.isCheckUsed() {
		s.err = ErrCheckMethodAlreadyUsed
	}

	s.ignoredFields = make(map[string]bool)

	for _, field := range fields {
		s.ignoredFields[field] = true
	}

	return s
}

func (s structMatcher) isIgnoreUsed() bool {
	return s.ignoredFields != nil
}

func (s structMatcher) isCheckUsed() bool {
	return s.checkedFields != nil
}

func (s structMatcher) addError(field string, message string) {
}
