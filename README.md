A simple `userscript -> local server` which saves my <http://cstimer.net> data every time I open cstimer in my browser

The userscript runs the underlying javascript export code whenever I open cstimer, and then `POST`s it to `localhost:8553`, which the server receives and then writes to a file

Requires:

- `go` <https://go.dev/>
- a userscript manager. If you don't already have a userscript manager, I recommend [Violentmonkey](https://violentmonkey.github.io/)

To run:

- install the userscript from [here](https://greasyfork.org/en/scripts/453183-cstimer-auto-download), (or manually install it from `cstimer_auto_download.js`)
- `go install 'github.com/seanbreckenridge/cstimer-save-server@latest`
- Run `cstimer-save-server` in the background somewhere:

`cstimer-save-server -save-to ~/Documents/cstimer`

That saves to `~/Documents/cstimer/cstimer.json`

```
$ cstimer-save-server -help
usage: cstimer-save-server [FLAG...]

  -port int
    	port to serve server on (default 8553)
  -save-to string
    	path to save datafile to.
  -timestamped
    	instead of writing to the same 'cstimer.json' file, write to a new file each time
```

Now whenever you open <https://cstimer.net/>, it should save your current solves to a local file:

```
$ go run ./server.go -save-to . -timestamped
2022/10/16 10:20:53 cstimer-save-server saving to '.' on port 8553
2022/10/16 10:20:55 Saving data to '1665940855868.json'
2022/10/16 10:20:56 Saving data to '1665940856882.json'
```

If you have a port conflict, change the `PORT` variable in `cstimer_auto_download.js`, and supply the `-port` flag to change what port the server launches on

### bleanser

I personally run this with [this script](https://sean.fish/d/cstimer-server?dark) - saving `-timestamped` filenames, and then run have a corresponding [bleanser module](https://github.com/seanbreckenridge/bleanser) script to remove 'useless' backups every few hours:

```
[ ~ ] $ python3 -m bleanser_sean.modules.cstimer prune ~/data/cubing/cstimer --remove --yes
[INFO    2022-10-16 10:57:04 bleanser.core.common main.py:144] processing 4 files (/home/sean/data/cubing/cstimer/1665942943939.json ... /home/sean/data/cubing/cstimer/1665943018015.json)
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:95] using 1 workers
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:373] processing /home/sean/data/cubing/cstimer/1665942943939.json (0/4)
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:387] cleanup(/home/sean/data/cubing/cstimer/1665942943939.json): took 0.01 seconds
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:373] processing /home/sean/data/cubing/cstimer/1665942944781.json (1/4)
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:387] cleanup(/home/sean/data/cubing/cstimer/1665942944781.json): took 0.01 seconds
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:373] processing /home/sean/data/cubing/cstimer/1665943017337.json (2/4)
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:387] cleanup(/home/sean/data/cubing/cstimer/1665943017337.json): took 0.01 seconds
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:373] processing /home/sean/data/cubing/cstimer/1665943018015.json (3/4)
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:387] cleanup(/home/sean/data/cubing/cstimer/1665943018015.json): took 0.01 seconds
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:468] emitting group pivoted on ['/home/sean/data/cubing/cstimer/1665942943939.json', '/home/sean/data/cubing/cstimer/1665943018015.json'], size 4
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:1015] 0  /4   /home/sean/data/cubing/cstimer/1665942943939.json : Keep
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:1015] 1  /4   /home/sean/data/cubing/cstimer/1665942944781.json : Prune
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:1015] 2  /4   /home/sean/data/cubing/cstimer/1665943017337.json : Prune
[DEBUG   2022-10-16 10:57:04 bleanser.core.common processor.py:1015] 3  /4   /home/sean/data/cubing/cstimer/1665943018015.json : Keep
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1076] processing    0/   4 /home/sean/data/cubing/cstimer/1665942943939.json : will keep          ; pruned so far:    0 Mb /   0 Mb ,   0 /  1 files
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1076] processing    1/   4 /home/sean/data/cubing/cstimer/1665942944781.json : REMOVE             ; pruned so far:    0 Mb /   0 Mb ,   1 /  2 files
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1076] processing    2/   4 /home/sean/data/cubing/cstimer/1665943017337.json : REMOVE             ; pruned so far:    0 Mb /   0 Mb ,   2 /  3 files
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1076] processing    3/   4 /home/sean/data/cubing/cstimer/1665943018015.json : will keep          ; pruned so far:    0 Mb /   0 Mb ,   2 /  4 files
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1078] SUMMARY: pruned so far:    0 Mb /   0 Mb ,   2 /  4 files
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1122] rm /home/sean/data/cubing/cstimer/1665942944781.json
[INFO    2022-10-16 10:57:04 bleanser.core.common processor.py:1122] rm /home/sean/data/cubing/cstimer/1665943017337.json
```
