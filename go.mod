module hellogo

go 1.16

require(
    hellogo.com/testing v0.0.0
)

replace (
    hellogo.com/testing v0.0.0 => ./testing
)