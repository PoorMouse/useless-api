package main

type errAPI map[string]string

//Errors
var (
	noDataInDBErr = errAPI{"error": "There is no data for that ID"}
	noUserIDErr   = errAPI{"error": "No user ID was given or ID is incorrect"}
	tooSlowErr    = errAPI{"error": "Database responds too slow"}
	unexpErr      = errAPI{"error": "Unexpected error. Everything's doomed! Try later"}
)
