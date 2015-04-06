package main
import (
    "github.com/streadway/amqp"
    "fmt"
)


func main() {
    println("start...")

//    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        fmt.Errorf("Dial: %s", err)
        return
    }
    defer connection.Close()

//    channel = connection.channel()
    channel, err := connection.Channel()
    if err != nil {
        fmt.Errorf("Channel: %s", err)
        return
    }
    defer channel.Close()

//    channel.queue_declare(queue='hello')
    queue, err := channel.QueueDeclare("hello", false, false, false, false, nil)
    if err != nil {
        fmt.Errorf("QueueDeclare: %s", err)
        return
    }
    fmt.Println("created queue: %s", queue.Name)

//    channel.basic_publish(exchange='', routing_key='hello', body='Hello World!')
    err = channel.Publish(
        "",
        "hello",
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte("Hello World!"),
    })
    if err != nil {
        fmt.Errorf("Publish: %s", err)
        return
    }
    println("[x] Sent 'Hello World!'")

    println("end...")
}
