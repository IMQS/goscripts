# goscripts

goscripts is a tiny package intended to make it easy to write little Go programs
that fulfill the need of a shell script, but for situations where it's just
extremely painful to get the job done with a shell script alone.

* Functions do not return errors - they panic if they encounter an error