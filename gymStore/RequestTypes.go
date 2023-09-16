package gymstore

type ExerciseRequest struct {
	Username     string `json:"username"`
	ExerciseName string `json:"ExerciseName"`
}

type AddExerciseRequest struct {
	ExerciseIdentifier ExerciseRequest `json:"exerciseIdentifer"`
	MainSet            AddUserInput    `json:"mainSet"`
	Dropsets           []AddUserInput  `json:"Dropsets"`
}

type AddUserInput struct {
	Reps    []float64
	Weights []float64
	Sets    int
	Weight  float64
	Date    string
	Note    string
}
