package gymstore

import (
	"fmt"
	"testing"
)

func TestFullFlow(t *testing.T) {
	tests := []struct {
		name     string
		username string
		request  AddExerciseRequest
	}{
		{
			name:     "Add a single exercise to user with dropset",
			username: "AdamBurgessFull",
			request: AddExerciseRequest{
				ExerciseIdentifier: ExerciseRequest{Username: "AdamBurgessFull", ExerciseName: "Bench Press"},
				MainSet: &AddUserInput{
					Reps:    []float64{1, 1, 1, 1},
					Weights: []float64{1, 1, 1, 1},
					Sets:    4,
					Weight:  1,
					Date:    "01-01-2000",
					Note:    "Test note"},
				Dropsets: []AddUserInput{
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					},
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					}},
			},
		},
	}

	StartUp("Testing/users/")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			//Phase 1: Add
			AddExercise(test.request)

			//Check if add worked

			user, ok := users[test.username]

			if !ok {
				t.Fatalf("User with username %s, does not exist: they should", test.username)
			}

			exerciseInstance, ok := user.Exercises[test.request.ExerciseIdentifier.ExerciseName]

			if !ok {
				t.Fatalf("Exercise instance expected of %s, does not exist", test.request.ExerciseIdentifier.ExerciseName)
			}

			fmt.Println(exerciseInstance)

			//Phase 2: get

			exerciseInstance, err := GetExercise(test.request.ExerciseIdentifier)

			if err != nil {
				t.Fatalf("Exercise instance expected of %s, was not returned form GetExercise(request\n", test.request.ExerciseIdentifier.ExerciseName)
			}

			set := exerciseInstance.Iterations[0]

			//Need to add better checking
			if set.Sets != 4 {
				t.Fatalf("Exercise instance expected of was returned incorrect format\n")
			}

			//Phase 3: Test Delete

			err = DeleteExercise(test.request.ExerciseIdentifier)

			if err != nil {
				t.Fatalf("unexpected error when deleting exercise which does exist, error thrown %s, from request %v\n", err, test.request.ExerciseIdentifier)
			}

			//Phase 4: test if exercise still exists

			exerciseInstance, err = GetExercise(test.request.ExerciseIdentifier)

			if err == nil {
				t.Fatalf("Exercise instance expected of %s, was returned form GetExercise(request). It should not now exist\n", test.request.ExerciseIdentifier.ExerciseName)
			}

		})
	}

}

func TestAddExercise(t *testing.T) {
	tests := []struct {
		name     string
		username string
		request  AddExerciseRequest
	}{
		{
			name:     "Add a single exercise to user with dropset",
			username: "AdamAdd",
			request: AddExerciseRequest{
				ExerciseIdentifier: ExerciseRequest{Username: "Adam", ExerciseName: "Bench Press"},
				MainSet: &AddUserInput{
					Reps:    []float64{1, 1, 1, 1},
					Weights: []float64{1, 1, 1, 1},
					Sets:    4,
					Weight:  1,
					Date:    "01-01-2000",
					Note:    "Test note"},
				Dropsets: []AddUserInput{
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					},
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					}},
			},
		},
	}

	StartUp("Testing/users/")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			AddExercise(test.request)

			user, ok := users[test.username]

			if !ok {
				t.Fatalf("User with username %s, does not exist: they should", test.username)
			}

			exerciseInstance, ok := user.Exercises[test.request.ExerciseIdentifier.ExerciseName]

			if !ok {
				t.Fatalf("Exercise instance expected of %s, does not exist", test.request.ExerciseIdentifier.ExerciseName)
			}

			fmt.Println(exerciseInstance)
		})
	}

}

func TestDeleteExercise(t *testing.T) {
	tests := []struct {
		name           string
		username       string
		addRequest     AddExerciseRequest
		deleteRequest  ExerciseInstanceRequest
		expectedLength int
	}{
		{
			name:     "Add a single exercise to user with dropset",
			username: "AdamDelete",
			addRequest: AddExerciseRequest{
				ExerciseIdentifier: ExerciseRequest{Username: "AdamDelete", ExerciseName: "Bench Press"},
				MainSet: &AddUserInput{
					Reps:    []float64{1, 1, 1, 1},
					Weights: []float64{1, 1, 1, 1},
					Sets:    4,
					Weight:  1,
					Date:    "01-01-2000",
					Note:    "Test note"},
				Dropsets: []AddUserInput{
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					},
					{
						Reps:    []float64{1, 1, 1, 1},
						Weights: []float64{1, 1, 1, 1},
						Sets:    4,
						Weight:  1,
						Date:    "01-01-2000",
						Note:    "Test note",
					}},
			},
			deleteRequest: ExerciseInstanceRequest{
				ExerciseIdentifier: ExerciseRequest{Username: "AdamDelete", ExerciseName: "Bench Press"},
				ID:                 0,
			},
			expectedLength: 1,
		},
	}

	StartUp("Testing/users/")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			//Not testing this
			AddExercise(test.addRequest)
			AddExercise(test.addRequest)

			DeleteExerciseIteration(test.deleteRequest)

			//test one delete

			user, ok := users[test.username]

			if !ok {
				t.Fatalf("User with username %s, does not exist: they should\n", test.username)
			}

			exerciseInstance, ok := user.Exercises[test.addRequest.ExerciseIdentifier.ExerciseName]

			if !ok {
				t.Fatalf("Exercise instance expected of %s, does not exist\n", test.addRequest.ExerciseIdentifier.ExerciseName)
			}

			result := len(exerciseInstance.Iterations)

			if result != test.expectedLength {
				t.Fatalf("Unexpected exercise instance length, got %d, expected %d\n", result, test.expectedLength)
			}

		})
	}

}
