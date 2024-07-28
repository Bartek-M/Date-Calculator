# Datetime Calculator
Simple datetime calculator made using **GO**. Add / Subtract any **time**, any **date**, in any amount.

> If you find any bugs, feel free to create a new issue in this repository.

## Requirements
- Golang 1.22.1 

## Usage
Build an executable application
```bash
go build .
```

First pick a mode: `time` or `date`. To change the mode run: `/time` or `/date`.

### `Time formats`
```
h:m:s | m:s | s

h - any number of hours
m - any number of minutes
s - any number of seconds
```

### `Date formats`
```
DD/MM/YYYY | DD/MM | DD
DD-MM-YYYY | DD-MM | DD

dd - day of the month 
mm - month of the year
yyyy - year
```

> **NOTE:** To exit the running program type `exit`, `quit` or simply press `CTRL+C`. 