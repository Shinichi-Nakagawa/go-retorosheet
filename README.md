GO-RETROSHEET
=============

Golang scripts for Retrosheet data downloading and parsing.

REQUIREMENTS
---------------

- Chadwick 0.6.2 http://chadwick.sourceforge.net/

- Go 1.5+

USAGE
---------------

### Download(Get Retrosheet CSV)

    go run download.go [-f <from 4-digit-year>] [-t <to 4-digit-year>]

### Parse(Retrosheet CSV to dataset csv for events & games)

    go run dataset.go [-f <from 4-digit-year>] [-t <to 4-digit-year>]

THANKS
---------------

this project is based on https://github.com/wellsoliver/py-retrosheet