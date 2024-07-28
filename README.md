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

Remember to use space around plus **'+'** and minus **'-'** signs.

> **NOTE:** To exit the running program type `exit`, `quit` or simply press `CTRL+C`. 

### `Time formats`
```
h:m:s | m:s | s

h - any number of hours
m - any number of minutes
s - any number of seconds

'now' - current, local time
```

### `Date formats`
```
DD/MM/YYYY | DD
DD-MM-YYYY | DD
DD.MM.YYYY | DD

dd - days of the month / any number of days
mm - month of the year
yyyy - year

'now' - current, local date
```

> **NOTE:** Output type is defined by the last operation: `days` - last operation with date, `date` - last operation with days