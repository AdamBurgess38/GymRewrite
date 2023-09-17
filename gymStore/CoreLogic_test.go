package gymstore

import (
	"fmt"
	"testing"
)

func TestAddExercise(t *testing.T) {
	tests := []struct {
		name     string
		username string
		request  AddExerciseRequest
	}{
		{
			name:     "Add a single exercise to user with dropset",
			username: "Adam",
			request: AddExerciseRequest{
				ExerciseIdentifier: ExerciseRequest{Username: "Adam", ExerciseName: "Bench Press"},
				MainSet: AddUserInput{
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
