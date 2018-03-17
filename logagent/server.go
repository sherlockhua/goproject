package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type KafkaSender struct {
	client sarama.SyncProducer
	lineChan chan string
}

var kafkaSender *KafkaSender

func NewKafkaSender (kafkaAddr string) (kafka *KafkaSender, err error) {
	kafka = &KafkaSender{
		lineChan: make (chan string, 100000),
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	/*
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")
	*/
	client, err := sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		logs.Error("init kafka client failed, err:%v", err)
		return
	}

	kafka.client = client
	for i := 0; i < appConfig.KafkaThreadNum; i++ {
		go kafka.sendToKafka()
	}

	return
}

func initKafka() (err error) {
	kafkaSender, err = NewKafkaSender(appConfig.kafkaAddr)
	return
}

func (k *KafkaSender) sendToKafka() {

}

func (k *KafkaSender) addMessage(line string) (err error) {
	k.lineChan <- line
	return
}

func RunServer() {

}