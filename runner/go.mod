module runner

go 1.20

replace MircoHandler4 => ../MircoHandler4

require (
	MircoHandler4 v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require MicroData v0.0.0-00010101000000-000000000000 // indirect

replace MicroData => ../MicroData
