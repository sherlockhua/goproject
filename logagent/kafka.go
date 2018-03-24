package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type Message struct {
	line string
	topic string
}

type KafkaSender struct {
	client sarama.SyncProducer
	lineChan chan *Message
}

var kafkaSender *KafkaSender

func NewKafkaSender (kafkaAddr string) (kafka *KafkaSender, err error) {
	kafka = &KafkaSender{
		lineChan: make (chan *Message, 100000),
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
	for v := range k.lineChan {
		msg := &sarama.ProducerMessage{}
		msg.Topic = v.topic
		msg.Value = sarama.StringEncoder(v.line)

		_, _, err := k.client.SendMessage(msg)
		if err != nil {
			logs.Error("send message to kafka failed, err:%v", err)
		}
	}
}

func (k *KafkaSender) addMessage(line string, topic string) (err error) {
	k.lineChan <- &Message{line:line, topic:topic}
	return
}
