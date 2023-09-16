package gymstore

import "fmt"

var store GymStoreContract

func GetExercise(request ExerciseRequest) (Exercise, error) {

	//Validate request

	//Fetch object
	//Might already exists....

	user, err := fetchUser(request)
	if err != nil {
		return Exercise{}, err
	}

	exercise, ok := user.Exercises[request.ExerciseName]
	if !ok {
		return Exercise{}, fmt.Errorf("exercise: %s does not exist for user: %s", request.Username, request.ExerciseName)
	}

	return exercise, nil

	//If it doesn't, try and fetch it from the current store implementation.
	//This will be an interface to fetch the object....
	//Via JSON
	//Via DB

	//If still doesn't exist return error
}

func generateDropSets(inputs []AddUserInput) []Dropset {

	dropsets := []Dropset{}

	makeDropset := func(reps []float64, weights []float64, variances []float64, sets int, weight float64, totalWeight float64, averageRep float64, averageWeight float64, averageWeightRepTotal float64) Dropset {
		return Dropset{
			Reps:                  reps,
			Weights:               weights,
			Sets:                  sets,
			Variances:             variances,
			Weight:                weight,
			TotalWeight:           totalWeight,
			AverageWeight:         averageWeight,
			AverageRep:            averageRep,
			AverageWeightRepTotal: averageWeightRepTotal,
		}
	}

	for _, dropset := range inputs {
		var totalWeight float64 = 0
		var totalWeightRep float64
		var totalReps float64 = 0
		for i, w := range dropset.Weights {
			totalWeightRep += w * dropset.Reps[i]
			totalReps += dropset.Reps[i]
			totalWeight += w
		}
		dropset.Sets = len(dropset.Reps)

		dropset := makeDropset(dropset.Reps, dropset.Weights, Map(dropset.Weights, func(item float64) float64 { return (item - dropset.Weight) }),
			dropset.Sets,
			dropset.Weight,
			totalWeightRep, totalReps/float64(len(dropset.Reps)), totalWeight/float64(len(dropset.Weights)), totalWeightRep/float64(len(dropset.Weights)))

		dropsets = append(dropsets, dropset)

	}

	return dropsets
}

func AddExercise(request AddExerciseRequest) error {

	user, err := fetchUser(request.ExerciseIdentifier)
	if err != nil {
		return err
	}

	exerciseInstance, ok := user.Exercises[request.ExerciseIdentifier.ExerciseName]

	if !ok {
		//Will need to add something to add a unique ID...or boot up IDs as we load the user. Hence change everytime allowing us to use the length
		exerciseInstance = Exercise{ID: 0, Iterations: make(map[int]Iteration)}
	}

	dropsets := generateDropSets(request.Dropsets)

	mainset := request.MainSet

	var totalWeight float64 = 0
	var totalWeightRep float64
	var totalReps float64 = 0
	for i, w := range mainset.Weights {
		totalWeightRep += w * mainset.Reps[i]
		totalReps += mainset.Reps[i]
		totalWeight += w
	}

	exerciseInstance.Iterations[len(exerciseInstance.Iterations)] = *NewIteration(mainset.Reps, mainset.Weights, Map(mainset.Weights, func(item float64) float64 { return (item - mainset.Weight) }), newID,
		mainset.Sets,
		mainset.Weight,
		mainset.Date,
		mainset.Note, totalWeightRep, totalReps/float64(len(mainset.Reps)), totalWeight/float64(len(mainset.Weights)), totalWeightRep/float64(len(mainset.Weights)), dropsets)

	user.Exercises[request.ExerciseIdentifier.ExerciseName] = exerciseInstance
	//Save to DB

	return nil
}

func GetAllExercises(request AddExerciseRequest) {
	//Validate request

	//Fetch object
	//Might already exists....
	//If it doesn't, try and fetch it from the current store implementation.
	//This will be an interface to fetch the object....
	//Via JSON
	//Via DB
}

//For now this will just create a new user, in the future you have to manually make one
func fetchUser(request ExerciseRequest) (User, error) {
	user, ok := users[request.Username]

	if !ok {
		err := loadUserDB(request, &user)

		if err != nil {
			//needs changing to not create new user
			users[request.Username] = User{Name: request.Username,
				Exercises: make(map[string]Exercise)}

			return User{}, err
		}
	}
	//Try and assign the user object
	return user, nil
}

//For now this will just create a new user, in the future you have to manually make one
func loadUserDB(request ExerciseRequest, user *User) error {

	//Try and assign the user object
	return nil
}

func NewIteration(reps []float64, weights []float64, variances []float64, ID int, sets int, weight float64, date string, note string, totalWeight float64, averageRep float64, averageWeight float64, averageWeightRepTotal float64, dropsets []Dropset) *Iteration {
	return &(Iteration{
		Reps:                  reps,
		Weights:               weights,
		Variances:             variances,
		ID:                    ID,
		Sets:                  sets,
		Weight:                weight,
		Date:                  date,
		Note:                  note,
		TotalWeight:           totalWeight,
		AverageRep:            averageRep,
		AverageWeight:         averageWeight,
		AverageWeightRepTotal: averageWeightRepTotal,
		DropSet:               dropsets,
	})
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
