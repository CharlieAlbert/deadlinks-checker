# deadlinks-checker

A fast and simple command-line tool to check for dead (broken) links on any website. Built in Go, it fetches a web page, extracts all anchor links, normalizes them, and checks their HTTP status concurrently.

## Features

- **Fetches HTML** from a given website URL
- **Extracts all anchor links** (`<a href=...>`) from the page
- **Normalizes and deduplicates links** (handles relative, absolute, and ignores non-HTTP links)
- **Checks link status concurrently** for fast results
- **Clear output**: shows which links are alive (✅) or dead (❌) with HTTP status

## Usage

```sh
$ go run main.go
Enter webiste URL: https://example.com

Checking 42 links...

✅ https://example.com/about (200 OK)
❌ https://example.com/broken (404 Not Found)
...
```

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/CharlieAlbert/deadlinks-checker.git
   cd deadlinks-checker
   ```
2. **Run the tool:**
   ```sh
   go run main.go
   ```
   Or build a binary:
   ```sh
   go build -o deadlinks-checker main.go
   ./deadlinks-checker
   ```

## Project Structure

- `main.go` — Entry point, handles user input and orchestrates the workflow
- `internal/`
  - `fetch/` — Fetches HTML content from a URL
  - `parser/` — Extracts anchor links from HTML using [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)
  - `utils/` — Normalizes and deduplicates links, resolves relative URLs
  - `checker/` — Checks link status (HTTP HEAD/GET), runs checks concurrently

## How it Works

1. **User Input:** Prompts for a website URL.
2. **Fetch HTML:** Downloads the HTML content of the page.
3. **Extract Links:** Parses the HTML and collects all `<a href=...>` links.
4. **Normalize Links:** Resolves relative links, removes duplicates, and filters out non-HTTP links (e.g., `mailto:`, `tel:`, `javascript:`).
5. **Check Links:** Sends HTTP HEAD (or GET) requests to each link concurrently, reporting status.
6. **Output:** Prints a list of all links with their status (alive/dead and HTTP status code).

## Dependencies

- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html) — HTML5-compliant parser for extracting links
- Standard Go libraries: `net/http`, `net/url`, `sync`, `time`, `os`, `bufio`, `strings`

## Example Output

```
Enter webiste URL: https://example.com

Checking 3 links...

✅ https://example.com/ (200 OK)
❌ https://example.com/broken (404 Not Found)
✅ https://www.iana.org/domains/example (200 OK)
```

## License

[BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause) (see dependency licenses)

---

_Built with Go. Contributions and issues welcome!_
