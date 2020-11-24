package _interface

import (
	"golang-clean-architecture/interface/kafka/producer"
	"golang-clean-architecture/interface/repository/testMysql"
)

var(
 MySqlRegistryRepo = new(testMysql.TestRegistryRepo)
 MySqlFetchRepo = new(testMysql.FetchRepo)
 KafkaProducer = new (producer.TestProducer)
)
