package types

type NewUser struct {
	FirstName string "json:firstName"
	LastName  string "json:lastName"
	Email     string "json:email"
	Password  string "json:password"
}
