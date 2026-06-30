# 🐹 GopherPulse

**GopherPulse** is a high-concurrency URL health monitoring tool built in Go. It uses a disciplined Worker Pool architecture to verify the status and latency of multiple targets simultaneously without overwhelming system resources.

## Features
* **Concurrency Control:** Uses a fixed Worker Pool to prevent "Thundering Herd" issues.
* **Polite Pinging:** Custom HTTP Client with a `User-Agent` to avoid being flagged as a bot.
* **Zombie Protection:** Strict 5-second timeouts to ensure workers never get stuck on hanging connections.
* **Real-time Analytics:** Tracks success rates and calculates average latency across all targets.
* **Zero Dependencies:** Built entirely using the Go Standard Library.

## Architecture
GopherPulse operates on a **Producer-Consumer** pipeline:
1.  **The Feeder (Producer):** A non-blocking goroutine that pushes URLs from CLI arguments into the Job Channel.
2.  **The Worker Pool (Consumers):** A fixed set of 5 goroutines that pull jobs, execute the network logic, and report results.
3.  **The Collector:** The main thread gathers results, updates real-time stats, and prints a final summary.



## Installation & Usage

### 1. Clone and Build
```bash
git clone [https://github.com/nilanshucodes/gopherpulse](https://github.com/nilanshucodes/gopherpulse)
cd gopherpulse
go build -o gopherpulse
```
### 2. Run
Pass any number of URLs as command-line arguments:
```bash
./gopherpulse [https://google.com](https://google.com) [https://github.com](https://github.com) [https://codeforces.com](https://codeforces.com)
```
## Technical Deep Dive

### Custom HTTP Client
Unlike the default `http.Get`, GopherPulse utilizes a manual `http.Client` configuration:

- **Timeout:** 5s (Prevents resource leaking on slow/dead servers).
- **Headers:** Sets `User-Agent: GopherPulse/1.0` to ensure compatibility with DDoS protection layers like Cloudflare.

### Non-Blocking Feed
By wrapping the job distribution in a goroutine, the program begins processing the first URL the moment it's available, rather than waiting to load the entire list into memory first.

<img width="804" height="218" alt="Screenshot 2026-03-22 at 14 55 32" src="https://github.com/user-attachments/assets/94ab0d93-bd79-4502-9160-db26e36cd9a7" />

