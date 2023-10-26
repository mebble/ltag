# https://unix.stackexchange.com/questions/46051/how-to-run-time-on-multiple-commands-and-write-the-time-output-to-file
/usr/bin/time -p sh -c "seq 100000000000 | ./ltag | head -n 3"
