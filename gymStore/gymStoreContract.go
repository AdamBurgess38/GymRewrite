package gymstore

//This store will think hey this user is already logged in, this is fine.
//Data should never leave this package.
/*Load exercise
Only when needed, add this to the map of the user. Will I need seperate JSON files for this? Probs...would be alot easier via DB

Add exercise

Delete exercise

Delete iteration

modify exercise? (One for the future*/
type GymStoreContract interface {
	LoadUser(username string) error
	SaveUser(user User) error
}
