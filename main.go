package main

import (
  "fmt"
  "os"
  "log"
  // "strings"
  // "time"
  // "context"
  "net/http"
  "github.com/joho/godotenv"
  "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
  // load the .env file
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  // log
  log.Println("Loaded the .env file")

  // get the info from env
  token := os.Getenv("INFLUXDB_TOKEN")
  bucket := os.Getenv("INFLUXDB_BUCKET")
  org := os.Getenv("INFLUXDB_ORG")
  host := os.Getenv("INFLUXDB_HOST")
  site_url := os.Getenv("SITE_URL")
  stat := 1

  client := influxdb2.NewClient(host, token)

  // log
  log.Println("Connected to the influxdb")

  // now check if the site is up
  resp, err := http.Get(site_url)
  if err != nil {
    stat = 0
  } else {
    defer resp.Body.Close()

    // log
    log.Println(fmt.Sprintf("Site's status code: %d", resp.StatusCode))
  }

  // log
  log.Println("Checked the site-url")

  line := fmt.Sprintf("stat,site=%s up=%d", site_url, stat)

  // get non-blocking write client
  writeAPI := client.WriteAPI(org, bucket)
  // write line protocol
  writeAPI.WriteRecord(line)
  // Flush writes
  writeAPI.Flush()

  // log
  log.Println("Inserted into influxdb")

  // always close client at the end
  defer client.Close()
}
