# ☕ Coffee Tracker CLI

A small CLI written in Go to track my coffee consumption.  

Because apparently the human body has a limit to how much caffeine it can handle, and I’m curious how close I am to discovering it.

---

## Features

- Log coffee by **type**, **size**, and **venue**  
- View today’s coffee stats  
- Designed to be **entirely unnecessary** and **over-engineered**  
- JSON-based storage (so you can inspect the chaos anytime)  

---

## Accepted Values

**Coffee Types:**  
`espresso`, `latte`, `cappucino`, `flatWhite`, `mocha`, `americano`, `filter`, `caffèCrema`, `instant`  

**Coffee Sizes:**  
`small`, `medium`, `large`  

**Venues:**  
`home`, `coffeeShop`, `office`  

---

## Usage

Clone the repo, build, and run:  

```bash
go build -o coffee
./coffee add -type=espresso -size=small -venue=home
./coffee today
```

---

## Commands
| Command | Description |
|---------|-------------|
| `add`   | Add a new coffee log |
| `today` | Show coffee stats for today |
| `week`  | Show coffee stats for this week |
| `month` | Show coffee stats for this month |
| `year`  | Show coffee stats for this year |
| `all`   | Show all-time coffee stats |

---

## Future Ideas

- Daily totals, graphs, and ASCII visualizations  
- “Set the coffee down” warning  
- build an gRPC version of the coffee tracker
- build a webapp for the tracker
- have the tracker output to an analogue display
---

## Contributing

Pull requests welcome, but don’t judge me too harshly.  
This project is a playground, not a production system.
