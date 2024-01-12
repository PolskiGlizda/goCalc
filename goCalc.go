package main

import(
    "fmt"
    "os"
    "net/http"
    "net/url"
    "io"
    "encoding/json"
    "github.com/TwiN/go-color"
)

type response struct{
    Queryresult struct{
        Success bool `json:"success"`
        Error bool `json:"error"`
        Numpods int `json:"numpods"`
        Datatypes string `json:"datatypes"`
        Timedout string `json:"timedout"`
        Timedoutpods string `json:"timedoutpods"`
        Timing float64 `json:"timing"`
        Parsetiming float64 `json:"parsetiming"`
        Parsetimedout bool `json:"parsetimedout"`
        Recalculate string `json:"recalculate"`
        Id string `json:"id"`
        Host string `json:"host"`
        Server string `json:"server"`
        Related string `json:"related"`
        Version string `json:"version"`
        Input string `json:"input"`
        Pods []struct{
            Title string `json:"title"`
            Scanner string `json:"scanner"`
            Id string `json:"id"`
            Position int `json:"position"`
            Error bool `json:"error"`
            Numsubpods int `json:"numsubpods"`
            Subpods []struct{
                Title string `json:"title"`
                Plaintext string `json:"plaintext"`
            } `json:"subpods"`
        } `json:"pods"`
    } `json:"queryresult"`
}



func main(){
    if len(os.Args) == 1 {
        fmt.Println("Usage: goCalc <query>")
        return
    }
    query := os.Args[1]
    query = url.QueryEscape(query)
    appid := "QVTRL7-JU2829Q55V"
    url := "http://api.wolframalpha.com/v2/query?input=" + query + "&appid=" + appid + "&output=JSON"
    resp, err := http.Get(url)
    if err != nil{
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil{
        fmt.Println(err)
        return
    }
    var response response
    err = json.Unmarshal(body, &response)
    if err != nil{
        fmt.Println(err)
        return
    }
    for i := 0; i < len(response.Queryresult.Pods); i++{
        fmt.Println(color.Ize(color.Yellow, response.Queryresult.Pods[i].Title))
        for j := 0; j < len(response.Queryresult.Pods[i].Subpods); j++{
            fmt.Println(color.Ize(color.Cyan, response.Queryresult.Pods[i].Subpods[j].Title))
            fmt.Println(color.Ize(color.Red, response.Queryresult.Pods[i].Subpods[j].Plaintext))
            fmt.Println()
        }
        fmt.Println()
    }
}
