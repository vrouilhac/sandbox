module vrouilhac/hello_world

go 1.18

require (
	rsc.io/quote v1.5.2
	vrouilhac/hello_module v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace vrouilhac/hello_module => ../hello_module
