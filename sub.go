package main

import (
  "os"
  "fmt"
  "log"
  "time"

  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var g_sub_counter int = 0
var g_start_time int64 = 0
var g_end_time int64 = 0

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
  g_sub_counter += 1

  if g_sub_counter == 1 {
    g_start_time = time.Now().UnixNano()
    fmt.Println("start time : %d ns", g_start_time)
  } else if g_sub_counter == 100000 {
    g_end_time = time.Now().UnixNano()
    fmt.Println("  end time : %d ns", g_end_time)
    fmt.Println("delta time : %d ns", g_end_time - g_start_time)
  }

  // fmt.Printf("Received message on topic: %s\n", message.Topic())
  // fmt.Printf("Message: %s\n", message.Payload())
}

func Subscribe(client *MQTT.MqttClient) error {
  topic := "my/topic/string"
  qos := 0

  topicFilter, err := MQTT.NewTopicFilter(topic, byte(qos))
  if err != nil {
    return err
  }

  _, err = client.StartSubscription(onMessageReceived, topicFilter)
  if err != nil {
    return err
  }

  for {
    time.Sleep(1 * time.Second)
  }
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

  err = Subscribe(client)
  if err != nil {
    log.Fatal(err)
  }
}
