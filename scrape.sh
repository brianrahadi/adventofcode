#!/usr/bin/env bash

YEAR="$1"
DAY="$2"
#create dir
mkdir -p "$YEAR/$DAY"
cd "$YEAR/$DAY" || exit

touch output

if [ ! -f main.go ]; then cp ../../go.template main.go; fi

if [ ! -f Makefile ]; then cat >Makefile <<EOF
main:
	go build -o main main.go

.PHONY: run clean

run: main
	./main <input

clean:
	rm -f main

EOF
fi

echo "Running with AOC_SESSION=$AOC_SESSION"
# download input files
http "https://adventofcode.com/$YEAR/day/$DAY/input" "Cookie:session=$AOC_SESSION;" >input

# download assignment
http "https://adventofcode.com/$YEAR/day/$DAY" | pup 'article.day-desc' >tmp.html
lynx -dump tmp.html -width 80 >assignment
rm -f tmp.html