package main

import (
  "os"
  "fmt"
  "log"
  "time"

  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var PUBLISH_CNT int = 1000000
var g_pub_counter int = 0
var g_start_time int64 = 0
var g_end_time int64 = 0


func Publish(client *MQTT.MqttClient, cnt int) error {
  topic := "my/topic/string"
  qos := 0
  message := fmt.Sprintf("hello%d", cnt)

  result := client.Publish(MQTT.QoS(qos), topic, message)
  <-result

  g_pub_counter += 1
  if g_pub_counter == 1 {
    g_start_time = time.Now().UnixNano()
    fmt.Printf("start time : %d ns\n", g_start_time)
  } else if g_pub_counter == PUBLISH_CNT {
    g_end_time = time.Now().UnixNano()
    fmt.Printf("  end time : %d ns\n", g_end_time)
    fmt.Printf("delta time : %d ns\n", g_end_time - g_start_time)
  }

  return nil
}

func main() {
  mqtt_server_addr := os.Getenv("MQTT_SERVER_ADDR")
  if mqtt_server_addr == "" {
    mqtt_server_addr = "localhost"
  }
  fmt.Println("MQTT server:", mqtt_server_addr)
  port := 1883

  opts := MQTT.NewClientOptions()

  brokerUri := fmt.Sprintf("tcp://%s:%d", mqtt_server_addr, port)
  opts.AddBroker(brokerUri)

  client := MQTT.NewClient(opts)

  _, err := client.Start()
  if err != nil {
      log.Fatal(err)
  }

  for i := 0; i < PUBLISH_CNT; i++ {
    err = Publish(client, i)
    if err != nil {
      log.Fatal(err)
    }
  }
}
