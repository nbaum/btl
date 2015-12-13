all: $(patsubst %.go.gpp,%.gen.go,$(wildcard *.go.gpp))

%.gen.go: %.go.gpp macros.gpp Makefile
	rm -f *.gen.go
	gpp -r "s/.go.gpp$$/.gen.go/" *.go.gpp
	gofmt -w -s *.gen.go
