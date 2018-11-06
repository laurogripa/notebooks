package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"

    "github.com/magrathealabs/warren-buffett/app/services"
)


type MarketValues struct {
    Coin   string                 `json:"coin"`
    Values []services.MarketValue `json:"values"`
}

func main() {
    file := os.Args[1]
    file_name := strings.Split(file, ".json")[0]

    data, err := ioutil.ReadFile(file_name + ".json")
    if err != nil {
        fmt.Println(err)
    }

    mv := MarketValues{}
    json.Unmarshal(data, &mv)

    fmt.Println(mv)

    newFile, err := os.Create(file_name + ".csv")
    if err != nil {
        fmt.Println(err)
    }
    defer newFile.Close()

    writer := csv.NewWriter(newFile)

    writer.Write([]string{"Date","Open","High","Low","Close","Volume","MarketCap"})
    for _, el := range(mv.Values) {
        var record []string
        record = append(record, el.Date.String())
        record = append(record, strconv.FormatFloat(el.Open, 'f', 1, 64))
        record = append(record, strconv.FormatFloat(el.High, 'f', 1, 64))
        record = append(record, strconv.FormatFloat(el.Low, 'f', 1, 64))
        record = append(record, strconv.FormatFloat(el.Close, 'f', 1, 64))
        record = append(record, strconv.FormatInt(int64(el.Volume), 10))
        record = append(record, strconv.FormatInt(int64(el.MarketCap), 10))
        writer.Write(record)
    }
    writer.Flush()
}
