package gymstore

var users map[string]User = make(map[string]User)

type User struct {
	Name      string
	Exercises map[string]Exercise
}

type Exercise struct {
	ID         int               //This only identifies the exercise via an ID....not the iteration of the exercise itself.
	Iterations map[int]Iteration //Should this also be made a map? --> Has been made into a map of an iteration.
}

type Iteration struct {
	Reps                  []float64
	Weights               []float64
	Variances             []float64
	ID                    int
	Sets                  int
	Weight                float64
	Date                  string
	Note                  string
	TotalWeight           float64
	AverageWeight         float64
	AverageRep            float64
	AverageWeightRepTotal float64
	DropSet               []Dropset
}

type Dropset struct {
	Reps                  []float64
	Weights               []float64
	Variances             []float64
	Sets                  int
	Weight                float64
	TotalWeight           float64
	AverageWeight         float64
	AverageRep            float64
	AverageWeightRepTotal float64
}
