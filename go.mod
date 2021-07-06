module hellogo

go 1.16

require (
	hellogo.com/mytesting v0.0.0
)

replace (
	hellogo.com/mytesting v0.0.0 => ./mytesting
)
