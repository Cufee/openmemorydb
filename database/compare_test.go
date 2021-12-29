package database

import "testing"

type TestCase struct {
	name          string
	struct1       map[string]interface{}
	struct2       map[string]interface{}
	expectedMatch bool
}

func TestCompare(t *testing.T) {
	var testCases = []TestCase{
		{
			name: "Maps 1 and 2 are the same",
			struct1: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: true,
		},
		{
			name: "Map 1 has multiple fields and some are not the same as Map 2",
			struct1: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test2",
			},
			expectedMatch: false,
		},
		{
			name: "Map 1 has multiple fields and some are not the same as Map 2, Map 2 has a field that is not in Map 1",
			struct1: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field3": "test2",
			},
			expectedMatch: false,
		},
		{
			name: "Map 1 has a field that is matching a field in Map 2",
			struct1: map[string]interface{}{
				"Field1": 1,
			},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: true,
		},
		{
			name: "Map 1 has a field that is not matching a field in Map 2",
			struct1: map[string]interface{}{
				"Field1": 2,
			},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: false,
		},
		{
			name:    "Map 1 is empty",
			struct1: map[string]interface{}{},
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: true,
		},
		{
			name:    "Map 1 is nil",
			struct1: nil,
			struct2: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: true,
		},
		{
			name:    "Map 2 is nil",
			struct2: nil,
			struct1: map[string]interface{}{
				"Field1": 1,
				"Field2": "test",
			},
			expectedMatch: false,
		},
		{
			name: "Map 1 and Map 2 are nested and are the same",
			struct1: map[string]interface{}{
				"Field1": map[string]interface{}{
					"Field1": 1,
					"Field2": "test",
				},
				"Field2": "test",
			},
			struct2: map[string]interface{}{
				"Field1": map[string]interface{}{
					"Field1": 1,
					"Field2": "test",
				},
				"Field2": "test",
			},
			expectedMatch: true,
		},
	}

	// Iterate through the test cases
	for _, testCase := range testCases {
		// Compare the structs
		match := CompareMaps(testCase.struct1, testCase.struct2)
		if match != testCase.expectedMatch {
			t.Errorf("%s: Expected match to be %t, got %t", testCase.name, testCase.expectedMatch, match)
		}
	}
}
