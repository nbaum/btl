# Golem

Golem is a programming language in the Lisp family. Its current status is best described as "broken toy".

Golem is written in Go, and the interpreter has no external dependencies.

## Usage

Compile it:

    go get github.com/nbaum/golem/cmd/golem

Pipe a script through the golem binary.

    echo '(prn "hello, world")' | golem

I'll make it hash-bangable when I can be bothered.

## The language

Golem is currently a mostly-complete implementation of Arc's foundation. I don't plan to keep it that way.
