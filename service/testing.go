package service

import (
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"l0/service/store"
	"testing"
)

func TestStreamClient(t *testing.T, ClientID, ClusterID string) *StreamClient {
	t.Helper()

	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	config := NewConfig()
	config.ClientID = ClientID
	config.ClusterID = ClusterID
	config.Database = &store.Config{URL: "host=localhost port=5432 user=WB password=Qwe123!!! dbname=WB_L0_test sslmode=disable"}

	s := New(config, c)

	if err := s.Start(); err != nil {
		t.Fatal(err)
	}

	return s

}

func TestPublish(t *testing.T, ClientID, ClusterID, channel string) {
	t.Helper()

	sc, err := stan.Connect(ClusterID, ClientID)

	if err != nil {
		t.Fatal(err)
	}

	data := []byte("{\n    \"order_uid\": \"b563feb7b2b84b6test\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	if err := sc.Publish(channel, data); err != nil {
		t.Fatal(err)
	}

	data = []byte("{\n    \"order_uid\": \"b563feb7b2b84b6test123\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	if err := sc.Publish(channel, data); err != nil {
		t.Fatal(err)
	}

	if err := sc.Close(); err != nil {
		t.Fatal(err)
	}
}
