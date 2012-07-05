stathat-line-counter
====================

__stathat-line-counter__ is a [Go](http://golang.org) program that reads stdin
and periodically logs lines read to StatHat.

It's useful for tracking statistics like web server hits, DNS queries etc.

Installation
------------

	go get github.com/jbarham/stathat-line-counter
	go install github.com/jbarham/stathat-line-counter

Example Usage
-------------

To log the number of Nginx hits every minute:

	tail -F /var/log/nginx/access.log | stathat-line-counter -stat "web hits" -ezkey yourname@example.com

(Note the use of the -F vs. -f flag to re-open the log file when it's rotated.)

About
-----

__stathat-line-counter__ was written by John Barham (jbarham@gmail.com).

It assumes you have an account with [StatHat](http://www.stathat.com), which is
a great statistics tracking service, and coincidentally is [largely written
in Go](http://blog.golang.org/2011/12/building-stathat-with-go.html).