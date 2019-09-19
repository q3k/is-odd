is-odd
======

A microservice to check whether a number is odd.

    $ go get github.com/q3k/is-odd
    $ go generate github.com/q3k/is-odd/...
    $ go build github.com/q3k/is-odd
    $ ./is-odd

    $ grpcurl -plaintext -format text -d 'number: 2138' 127.0.0.1:2137 isodd.IsOdd.IsOdd
    result: RESULT_NON_ODD
    $ grpcurl -plaintext -format text -d 'number: 2137' 127.0.0.1:2137 isodd.IsOdd.IsOdd
    result: RESULT_ODD

