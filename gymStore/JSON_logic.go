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
	return []Dropset{}
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

	x := request.MainSet

	var totalWeight float64 = 0
	var totalWeightRep float64
	var totalReps float64 = 0
	for i, w := range x.Weights {
		totalWeightRep += w * x.Reps[i]
		totalReps += x.Reps[i]
		totalWeight += w
	}

	exerciseInstance.Iterations[len(exerciseInstance.Iterations)] = *NewIteration(x.Reps, x.Weights, Map(x.Weights, func(item float64) float64 { return (item - x.Weight) }), newID,
		x.Sets,
		x.Weight,
		x.Date,
		x.Note, totalWeightRep, totalReps/float64(len(x.Reps)), totalWeight/float64(len(x.Weights)), totalWeightRep/float64(len(x.Weights)), dropsets)

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
