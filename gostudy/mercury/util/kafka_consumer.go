package util

import (
	"strings"
	"sync"

	"github.com/gnufree/gostudy/logger"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func InitKafkaConsumer(addr, topic string, consume func(message *sarama.ConsumerMessage)) (err error) {

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logger.Error("init kafka consumer failed, addr:%v, err:%v", addr, err)
		return
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logger.Error("get the list of partitions :%v", err)
		return
	}
	for partition := range partitionList {
		pc, errRet := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logger.Error("Failed to start consumer for partition %v: err:%v", partition, err)
			return
		}
		//defer pc.AsyncClose()
		wg.Add(1)
		go func(pc1 sarama.PartitionConsumer) {
			for msg := range pc1.Messages() {
				logger.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				consume(msg)
			}
			wg.Done()
		}(pc)
	}
	return
}
