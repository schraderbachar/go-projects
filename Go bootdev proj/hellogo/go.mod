module github.com/schraderbachar/hellogo

go 1.20

replace github.com/schraderbachar/mystrings v0.0.0 => ../mystrings //instead of looking for the thing on the internet, it gives it for us here

require (
	github.com/schraderbachar/mystrings v0.0.0
)
