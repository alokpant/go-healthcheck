# HealthCheck

A simple command-line tool to check if a given domain is reachable on a specific TCP port.  
Useful for quick website or service availability checks.

---

## Features
- Test connectivity to any domain and TCP port.
- Default port is **80** if none is provided.
- Reports whether the target is **UP** or **DOWN** along with connection details or error messages.
- Lightweight, no external dependencies except `urfave/cli`.

---

## Installation

1. **Clone this repository**:
```bash
git clone https://github.com/yourusername/healthcheck.git
cd healthcheck
```

2. **Install dependencies**:
```bash
go mod tidy
```

3. **Build the executable:**:
```bash
go build -o healthcheck
```

## Usage

**Build the executable:**:
```bash
go run . --domain alokpant.com
```

**Check if a domain is up on a custom port:**:
```bash
go run . --domain alokpant.com --port 443
```

**Short flag version:**:
```bash
go run . -d example.com -p 443
```

## Example Output:

**When site is reachable:**:
```bash
[UP] example.com is reachable. 
 From: 192.168.1.10:55433 
 To: 93.184.216.34:443
```

**When site is down/unreachable:**:
```bash
[DOWN] example.com is unreachable: 
 Error: dial tcp 93.184.216.34:443: i/o timeout
```

