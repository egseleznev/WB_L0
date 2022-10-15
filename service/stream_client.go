package service

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"l0/service/store"
)

type StreamClient struct {
	config       *Config
	connection   stan.Conn
	subscription stan.Subscription
	logger       *logrus.Logger
	cache        *cache.Cache
	database     *store.Database
}

func New(config *Config, cache *cache.Cache) *StreamClient {
	return &StreamClient{
		config: config,
		logger: logrus.New(),
		cache:  cache,
	}
}

func (s *StreamClient) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *StreamClient) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	con, err := stan.Connect(s.config.ClusterID, s.config.ClientID)
	if err != nil {
		return err
	}
	s.connection = con
	s.logger.Info("Connected to nats streaming server")
	return nil
}

func (s *StreamClient) Close() error {
	if err := s.connection.Close(); err != nil {
		return err
	}
	s.logger.Info("Disconnected from nats streaming server")
	return nil
}

func (s *StreamClient) Subscribe(subject string) error {
	sub, err := s.connection.Subscribe(subject, func(m *stan.Msg) {
		if validateMessage(m) == 0 {
			s.logger.Info("Valid message recieved")
			key, jsonValue := parseMessage(m)
			s.cache.Set(key, jsonValue, cache.NoExpiration)
			s.logger.Info("Cache set with order_uid: " + string(key))
		} else {
			s.logger.Warn("Unvalid message recieved")
		}
	})
	if err != nil {
		return err
	}
	s.subscription = sub
	s.logger.Info("Subscribed on channel")
	return nil
}

func (s *StreamClient) Unsubscribe() error {
	if err := s.subscription.Unsubscribe(); err != nil {
		return err
	}
	s.logger.Info("Unsubscribed from channel")
	return nil
}

func validateMessage(m *stan.Msg) int {
	f := make(map[string]interface{})
	err := json.Unmarshal(m.Data, &f)
	if err != nil || len(f) != 14 || len(f["order_uid"].(string)) != 19 {
		return -1
	}
	return 0
}

func parseMessage(m *stan.Msg) (string, map[string]interface{}) {
	jsonValue := make(map[string]interface{})
	json.Unmarshal(m.Data, &jsonValue)
	key := jsonValue["order_uid"]
	return key.(string), jsonValue
}
