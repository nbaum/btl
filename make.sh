set -e
for X in $(find -name '*.gop'); do
  gopp $X > ${X%.gop}.generated.go
done
go install github.com/nbaum/golem/cmd/...
cat test.gl | golem
