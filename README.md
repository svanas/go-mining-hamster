# go-mining-hamster

Package go-mining-hamster is a Golang SDK for the mininghamster.com trading signals API

## Installation

```
$ go get github.com/svanas/go-mining-hamster
```

## Usage

```
import (
	mininghamster "github.com/svanas/go-mining-hamster"
)
```

## Example

```
client := mininghamster.New(mininghamster.DemoKey)
signals, err := client.Get()
if err == nil {
	out, _ := json.Marshal(signals)
	fmt.Println(string(out))
}
```
