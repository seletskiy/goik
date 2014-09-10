goik
====

Small tool for parsing iCal format.

Can parse even utterly broken input.

Install
=======

Just `go get` package:

```
go get github.com/seletskiy/goik
```

Usage
=====

`goik` eats iCal on standard input and outputs it in a pretty format.

For example:

```
$ go run main.go < ~/meeting.ics
      August 2014    
 Su Mo Tu We Th Fr Sa
                 1  2 
  3  4  5  6  7[ 8] 9 
 10 11 12 13 14 15 16 
 17 18 19 20 21 22 23 
 24 25 26 27 28 29 30 
 31                   
 
Time :: 13:00 (in local time)
Summary :: Interview with someone http://intranet/browse/RESUME-1153
Organizer :: s.seletskiy[at]office.ngs.ru
```

Integration
===========

With using `mutt` as e-mail client it's pretty easy to get use of `goik`:

Just add following line to the `~/.mailcap`:

```
text/calendar;                  goik; copiousoutput
```

All e-mails with ICS contents will be automagically rendered in pretty format. Wow.
