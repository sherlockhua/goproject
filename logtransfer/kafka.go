package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var consumer sarama.Consumer


type LogData struct {
	line string
	topic string
}

var msgChan chan *LogData

func initKafka(addr string) (err error) {
	consumer, err = sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		logs.Error("Failed to start consumer: %v", err)
		return
	}

	msgChan = make(chan *LogData, 10000)
	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		logs.Error("Failed to get the list of partitions, err:%v", err)
		return
	}

	logs.Debug("get partition from kafka succ, partition:%v", partitionList)
	for partition := range partitionList {
		pc, errRet := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		go func(p sarama.PartitionConsumer) {
			for msg := range p.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				logData := &LogData{
					line:string(msg.Value),
					topic: msg.Topic,
				}

				msgChan <- logData
			}
		}(pc)
	}
	
	return
}

func GetMessage() chan *LogData {
	return msgChan
}