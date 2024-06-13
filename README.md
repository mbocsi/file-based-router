# File based router

A server that routes incoming http requests to a file structure that contains html, javascript, css, and cgi files.

## Usage

### Build

`make build` \
Outputs binary executable to `bin/fbr`

### Run

```
./fbr
-port int
Listening port (default 8080)
-root string
The root route directory (default "routes")
```

### File structure

Any file structure with .js, .css, .html, and .cgi files.
Router looks at index.[js/css/html/cgi] files by default unless file name is explicit in the request URL.

#### Example

routes/ -- (root)\
&emsp; api/\
&emsp;&emsp; index.cgi \
&emsp;&emsp; foo.cgi \
&emsp; index.html\
&emsp; index.js\
&emsp; style.css

- Request: GET / \
  Returns content of `index.html`
- Request: GET /styles.css \
  Returns content of `styles.css`
- Request: GET /api \
  Executes and returns STDOUT of `index.cgi`
- Request: POST /api/foo.cgi \
  Executes and returns STDOUT of `foo.cgi`
