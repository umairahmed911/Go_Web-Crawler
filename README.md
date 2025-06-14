# 🕷️ Go Web Crawler

*A Concurrent Web Crawler Built with Go*

![Go](https://img.shields.io/badge/Language-Go-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## 📌 Overview

This project is a concurrent web crawler built in **Go (Golang)**. It traverses web pages up to a specified depth, respects rate limiting, avoids duplicate visits, and collects metadata such as page titles and HTTP status codes.

The goal of this crawler is to showcase Go’s powerful **concurrency model** using goroutines, channels, and synchronization mechanisms such as `sync.Map` and `sync.WaitGroup`.

---

## ✨ Features

* 🌐 Crawl web pages up to a defined depth
* 🔁 Avoids revisiting the same URLs (thread-safe)
* 🚦 Respects rate limiting (prevents overload)
* 📝 Extracts page titles and HTTP status codes
* 📝 Store them in the xlsx file
* ⚙️ Built using lightweight and efficient Go routines
* 📦 No third-party libraries required (only Go standard library)

---

## 🏗️ Project Structure

```
go-web-crawler/
├── crawler.go       # Main crawler logic
├── go.mod           # Go module definition
├── README.md        # Project documentation
```

---

## 🚀 Getting Started

### 🔧 Prerequisites

* Go 1.16+ installed
* Internet connection

### 📥 Installation

```bash
git clone https://github.com/yourusername/go-web-crawler.git
cd go-web-crawler
go mod tidy
```

### ▶️ Run the Crawler

```bash
go run crawler.go
```

You can modify the **start URL** and **crawl depth** in the `main` function inside `crawler.go`.

---

## 🧠 How It Works

1. The crawler starts with a root URL.
2. It fetches the content of the page and extracts metadata (title, status).
3. It finds new links on the page and recursively visits them (up to a max depth).
4. It avoids revisiting the same link using a `sync.Map`.
5. Rate limiting ensures that requests are throttled to avoid overwhelming servers.

---

## 🔍 Example Output

```bash
Crawling: https://example.com
Status: 200 OK
Title: Example Domain

Crawling: https://example.com/about
Status: 200 OK
Title: About Us
```

---

## 🔐 Security Considerations

* The crawler does not currently check `robots.txt`—use responsibly.
* Rate limiting is implemented to reduce load on target servers.

---

## 📈 Future Enhancements

* [ ] Add support for `robots.txt`
* [ ] Implement URL filtering (e.g., domain-only crawling)
* [ ] Save results to file or database
* [ ] Parallelize across machines (distributed crawling)
* [ ] Add CLI support for custom flags (URL, depth, delay)

---

## 🧾 License

This project is licensed under the **MIT License**.
See the [LICENSE](./LICENSE) file for details.

---

## 🙋‍♂️ Author

**Umair Ahmed**


---

## ⭐ Show Your Support

If you found this project helpful, feel free to:

* ⭐ Star the repo
* 📥 Fork it
* 🛠️ Contribute to improvements

---

