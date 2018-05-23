.PHONY: gofuzzbuild
gofuzzbuild: 
	go-fuzz-build github.com/francoispqt/gojay/examples/fuzz

.PHONY: gofuzz
gofuzz: 
	go-fuzz -bin=fuzz-fuzz.zip -workdir=.

.PHONY: gofuzzclean
gofuzzclean: 
	rm -rf corpus crashers suppressions fuzz-fuzz.zip