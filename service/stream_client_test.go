package service_test

import (
	"github.com/stretchr/testify/assert"
	"l0/service"
	"l0/service/models"
	"testing"
)

func TestValidateMessage(t *testing.T) {

	unvalidByteStream := []byte("sldkgslkgha;sldghas;lgh")
	isValid := service.ValidateMessage(unvalidByteStream)
	assert.False(t, isValid)

	unvalidUID := []byte("{\n    \"order_uid\": \"b563feb7b2b84b6test123456\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	isValid = service.ValidateMessage(unvalidUID)
	assert.False(t, isValid)

	noUID := []byte("{\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	isValid = service.ValidateMessage(noUID)
	assert.False(t, isValid)

	unvalidFields := []byte("{\n    \"order\": \"b563feb7b2b84b6test\",\n    \"tracking\": \"123\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"surname\": \"Test Testov\",\n        \"email\": \"+9720000000\",\n        \"code\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	isValid = service.ValidateMessage(unvalidFields)
	assert.False(t, isValid)

	unvalidFields = []byte("{\n    \"order_uid\": \"b563feb7b2b84b6test\",\n    \"track_number\": \"WBILMTESTTRACK\",\n}")
	isValid = service.ValidateMessage(unvalidFields)
	assert.False(t, isValid)

	valid := []byte("{\n    \"order_uid\": \"b563feb7b2b84b6test\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	isValid = service.ValidateMessage(valid)
	assert.True(t, isValid)

	valid = []byte("{\n    \"order_uid\": \"b563feb7b2b84b651sd\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}")
	isValid = service.ValidateMessage(valid)
	assert.True(t, isValid)

}

func TestParseMessage(t *testing.T) {

	correctOrder := &models.Order{
		Uid:  "b563feb7b2b84b651sd",
		Data: []byte("{\n    \"order_uid\": \"b563feb7b2b84b651sd\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"),
	}
	parsed := service.ParseMessage([]byte("{\n    \"order_uid\": \"b563feb7b2b84b651sd\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"))
	assert.Equal(t, correctOrder, parsed)

	correctOrder = &models.Order{
		Uid:  "b563feb7b2b84b65test",
		Data: []byte("{\n    \"order_uid\": \"b563feb7b2b84b65test\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"),
	}
	parsed = service.ParseMessage([]byte("{\n    \"order_uid\": \"b563feb7b2b84b65test\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"))
	assert.Equal(t, correctOrder, parsed)

	correctOrder = &models.Order{
		Uid:  "235jgffa951lsofbro2",
		Data: []byte("{\n    \"order_uid\": \"235jgffa951lsofbro2\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"),
	}
	parsed = service.ParseMessage([]byte("{\n    \"order_uid\": \"235jgffa951lsofbro2\",\n    \"track_number\": \"WBILMTESTTRACK\",\n    \"entry\": \"WBIL\",\n    \"delivery\": {\n        \"name\": \"Test Testov\",\n        \"phone\": \"+9720000000\",\n        \"zip\": \"2639809\",\n        \"city\": \"Kiryat Mozkin\",\n        \"address\": \"Ploshad Mira 15\",\n        \"region\": \"Kraiot\",\n        \"email\": \"test@gmail.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"b563feb7b2b84b6\",\n        \"request_id\": \"\",\n        \"currency\": \"USD\",\n        \"provider\": \"wbpay\",\n        \"amount\": 1817,\n        \"payment_dt\": 1637907727,\n        \"bank\": \"alpha\",\n        \"delivery_cost\": 100,\n        \"goods_total\": 1,\n        \"custom_fee\": 0\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 9934930,\n            \"track_number\": \"WBILMTESTTRACK\",\n            \"price\": 453,\n            \"rid\": \"ab4219087a764ae0btest\",\n            \"name\": \"Mascaras\",\n            \"sale\": 30,\n            \"size\": \"0\",\n            \"total_price\": 317,\n            \"nm_id\": 2389212,\n            \"brand\": \"Vivienne Sabo\",\n            \"status\": 202\n        }\n    ],\n    \"locale\": \"en\",\n    \"internal_signature\": \"\",\n    \"customer_id\": \"test\",\n    \"delivery_service\": \"meest\",\n    \"shardkey\": \"9\",\n    \"sm_id\": 99,\n    \"date_created\": \"2021-11-26T06:22:19Z\",\n    \"oof_shard\": \"1\"\n}"))
	assert.Equal(t, correctOrder, parsed)

}

func TestStreamClient_Close(t *testing.T) {

	s := service.TestStreamClient(t, "test_subscriber", "prod")

	err := s.Close()

	assert.NoError(t, err)

}

func TestStreamClient_Subscribe(t *testing.T) {

	s := service.TestStreamClient(t, "test_subscriber", "prod")

	if err := s.Subscribe("test_channel"); err != nil {
		t.Fatal(err)
	}

	if err := s.Unsubscribe(); err != nil {
		t.Fatal(err)
	}

	service.TestPublish(t, "test_publisher", "prod", "test_channel")

	err := s.Subscribe("test_channel")

	assert.NoError(t, err)

}

func TestStreamClient_Unsubscribe(t *testing.T) {

	s := service.TestStreamClient(t, "test_subscriber", "prod")

	if err := s.Subscribe("test_channel"); err != nil {
		t.Fatal(err)
	}

	err := s.Unsubscribe()

	assert.NoError(t, err)
}
