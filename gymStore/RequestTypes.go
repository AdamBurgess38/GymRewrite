package gymstore

import "fmt"

type Validator interface {
	validate() (bool, error)
}

type ExerciseRequest struct {
	Username     string `json:"username"`
	ExerciseName string `json:"ExerciseName"`
}

type AddExerciseRequest struct {
	ExerciseIdentifier ExerciseRequest `json:"exerciseIdentifer"`
	MainSet            *AddUserInput   `json:"mainSet"`
	Dropsets           []AddUserInput  `json:"Dropsets"`
}

type ExerciseInstanceRequest struct {
	ExerciseIdentifier ExerciseRequest `json:"exerciseIdentifer"`
	ID                 int
}

type AddUserInput struct {
	Reps    []float64
	Weights []float64
	Sets    int
	Weight  float64
	Date    string
	Note    string
}

func (exerciseRequest AddExerciseRequest) validate() (bool, error) {
	valid, err := exerciseRequest.ExerciseIdentifier.validate()

	if err != nil {
		return valid, err
	}

	if exerciseRequest.MainSet == nil {
		return false, fmt.Errorf("no mainset is provided for %s - %s", exerciseRequest.ExerciseIdentifier.Username, exerciseRequest.ExerciseIdentifier.ExerciseName)
	}

	valid, err = exerciseRequest.MainSet.validate()

	if err != nil {
		return valid, err
	}

	for _, dropsets := range exerciseRequest.Dropsets {
		valid, err = dropsets.validate()

		if err != nil {
			return valid, err
		}
	}

	return true, nil
}

func (exerciseRequest AddUserInput) validate() (bool, error) {
	if len(exerciseRequest.Reps) != len(exerciseRequest.Weights) {
		return false, fmt.Errorf("Reps and Weights per set do not allign")
	}

	return true, nil
}

func (exerciseRequest ExerciseRequest) validate() (bool, error) {
	if (len(exerciseRequest.Username)) <= 0 {
		return false, fmt.Errorf("username is empty")
	}
	if (len(exerciseRequest.Username)) <= 0 {
		return false, fmt.Errorf("exerciseName is empty")
	}
	return true, nil
}

func (exerciseRequest ExerciseInstanceRequest) validate() (bool, error) {

	valid, err := exerciseRequest.ExerciseIdentifier.validate()

	if err != nil {
		return valid, err
	}

	if exerciseRequest.ID < 0 {
		return false, fmt.Errorf("ID requested is less than 0")
	}

	return true, nil
}
