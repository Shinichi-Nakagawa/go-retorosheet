GO-RETROSHEET
=============

Golang scripts for Retrosheet data downloading and parsing.

YE REQUIREMENTS
---------------

- Chadwick 0.6.2 http://chadwick.sourceforge.net/

- Go 1.5+

USAGE
-----

### Download

    python retrosheet_download.py [-f <from 4-digit-year>] [-t <to 4-digit-year>] [-c <config.ini path>]

### Parse

    python parse_csv.py [-f <from 4-digit-year>] [-t <to 4-digit-year>] [-c <config.ini path>]

### Into SQL

    python retrosheet_mysql.py [-f <from 4-digit-year>] [-t <to 4-digit-year>] [-c <config.ini path>]

### Migration(Download - Parse - Into SQL)

    python migration.py [-f <from 4-digit-year>] [-t <to 4-digit-year>] [-c <config.ini path>]


YE GRATITUDE
------------

Github user jeffcrow made many fixes and additions and added sqlite support

JUST THE DATA
-------------

If you're using PostgreSQL (and you should be), you can get a dump of all data up through 2014 (warning: 502MB) [here](https://www.dropbox.com/s/nv9712l1ylvh64n/retrosheet.psql?dl=0)
