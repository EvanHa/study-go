module mytesting

go 1.16

require (
    gopkg.in/yaml.v2 v2.4.0 // indirect
    // hellogo.com/mytesting/calculator v0.0.0
    // hellogo.com/mytesting/yamltohtml v0.0.0
)

replace (
    // hellogo.com/mytesting/calculator v0.0.0 => ./calculator
    // hellogo.com/mytesting/yamltohtml v0.0.0 => ./yamltohtml
)
