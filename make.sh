set -e
rm -f $(find -name '*.generated.go')
for X in $(find -name '*.gop'); do
  gopp $X > ${X%.gop}.generated.go
done
go install github.com/nbaum/golem/cmd/...
cat lib.gl test.gl | golem
