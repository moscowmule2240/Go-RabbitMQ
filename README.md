# Go-RabbitMQ
sample golang rabbitmq

## RabbitMQ

### install
    brew install rabbitmq

### config
    ~/.profile
    export PATH=$PATH:/usr/local/sbin

### start
    rabbitmq-server

### admin page
    http://localhost:15672/
    Username:guest
    Password:guest

## go

### install
    brew install go
    
### gopath
    $ mkdir ~/Workspace/go/.go
    $ vi ~/.profile
      export GOPATH=~/Workspace/go/.go
      export PATH=$GOPATH/bin:$PATH

### requirements
    go get github.com/streadway/amqp
    go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git
