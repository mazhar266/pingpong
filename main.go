package main

import (
  "fmt"
  "os"
  "log"
  "time"
  "context"
  "github.com/joho/godotenv"
  "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
  // load the .env file
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  // get the info from env
  token := os.Getenv("INFLUXDB_TOKEN")
  bucket := os.Getenv("INFLUXDB_BUCKET")
  org := os.Getenv("INFLUXDB_ORG")
  host := os.Getenv("INFLUXDB_HOST")
  siteurl := os.Getenv("SITE_URL")

  client := influxdb2.NewClient(host, token)

  // now check if the site is up

  // get non-blocking write client
  writeAPI := client.WriteAPI(org, bucket)

  // write line protocol
  writeAPI.WriteRecord("stat,site=x up=0")
  writeAPI.WriteRecord("stat,site=x up=1")
  // Flush writes
  writeAPI.Flush()

  // always close client at the end
  defer client.Close()
}
