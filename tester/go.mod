module example.com/tester

go 1.20

replace example.com/practice => ../practice

require (
	example.com/backend v0.0.0-00010101000000-000000000000
	example.com/practice v0.0.0-00010101000000-000000000000
)

replace example.com/backend => ../backend
