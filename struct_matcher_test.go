package gomockmatcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_StructMatcher(t *testing.T) {
	t.Run("should match first level struct", func(t *testing.T) {
		type someStruct struct {
			Name      string
			Age       int
			CreatedOn time.Time
		}

		var testStruct someStruct
		testStruct.Name = "Fathil"
		testStruct.Age = 30
		testStruct.CreatedOn = time.Now()

		s := StructMatcher(testStruct).Fields([]string{"Name", "Age"})

		assert.True(t, s.Matches(someStruct{
			Name:      "Fathil",
			Age:       30,
			CreatedOn: time.Now(),
		}))
	})

	t.Run("should match second level struct", func(t *testing.T) {
		type childStruct struct {
			Name string
			Age  int
		}

		type someStruct struct {
			Name      string
			Age       int
			Child     childStruct
			CreatedOn time.Time
		}

		var testStruct someStruct
		testStruct.Name = "Fathil"
		testStruct.Age = 30
		testStruct.CreatedOn = time.Now()
		testStruct.Child.Name = "Arham"
		testStruct.Child.Age = 10

		s := StructMatcher(testStruct).Fields([]string{"Name", "Child.Name"})

		assert.True(t, s.Matches(someStruct{
			Name:      "Fathil",
			Age:       30,
			CreatedOn: time.Now(),
			Child: childStruct{
				Name: "Arham",
				Age:  5,
			},
		}))
	})
}
