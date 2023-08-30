# ager

Calculate age from date

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/ager)](https://goreportcard.com/report/github.com/prongbang/ager)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

```shell
go get github.com/prongbang/ager
```

### How to use

```go
format := "2/1/2006"
todayDate, _ := time.Parse(format, "9/2/2022")
birthDate, _ := time.Parse(format, "8/1/2022")

age := ager.Age(birthDate, todayDate)
// Year: 0, Month: 1, Day: 1
```
