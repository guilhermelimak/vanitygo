# Tries Seconds Hashes per second

## Without print
 - 259 - 3,556s = 72.8h/s
 - 49 - 1,079s = 45.4h/s
 - 10 - 0,480 = 20.8h/s

## With print:
 - 20 - 0,599s = 33.3h/s
 - 17 - 0,590 = 28.8h/s
 - 40 - 0,851 = 47.0h/s
 - 127 - 2,002s = 63.4h/s

## Without print:

 - Tries: 259<br />
   `go run \*.go v np 3,02s user 0,93s system 110% cpu 3,556 total`

 - Tries: 49<br />
   `go run \*.go v np 0,78s user 1,09s system 173% cpu 1,079 total`

 - Tries: 10<br />
   `go run \*.go v np 0,34s user 0,32s system 136% cpu 0,480 total`

## With print:
 - Tries: 20<br />
   `go run \*.go v np 0,50s user 0,16s system 109% cpu 0,599 total`

 - Tries: 17<br />
   `go run \*.go v np 0,46s user 0,34s system 136% cpu 0,590 total`

 - Tries: 127<br />
   `go run \*.go v np 1,64s user 0,56s system 110% cpu 2,002 total`
