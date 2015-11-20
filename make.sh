set -e
rm -f $(find -name '*.generated.go')
rm -f $(find -name '*.generated.gop')
for X in $(find -name '*.gop.m4'); do
  m4 $X > ${X%.m4}.generated.gop
done
for X in $(find -name '*.gop'); do
  m4 $X | gopp $X > ${X%.gop}.generated.go
done
go install github.com/nbaum/golem/cmd/...
cat lib.gl test.gl | golem
