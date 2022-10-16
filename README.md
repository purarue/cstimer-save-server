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
