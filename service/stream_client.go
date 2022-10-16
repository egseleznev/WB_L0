package service

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"l0/service/models"
	"l0/service/store"
	"log"
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

func (s *StreamClient) Start() error {
	if err := s.configureLogger(); err != nil {
		log.Println("Can't configure logger with error: " + err.Error())
		return err
	}

	if err := s.configureDatabase(); err != nil {
		s.logger.Fatal("Can't configure DB with error: " + err.Error())
		return err
	}

	con, err := stan.Connect(s.config.ClusterID, s.config.ClientID)
	if err != nil {
		s.logger.Fatal("Can't connect to nats streaming server with error: " + err.Error())
		return err
	}
	s.connection = con
	s.logger.Info("Connected to nats streaming server")
	if err := s.restoreCacheFromDB(); err != nil {
		s.logger.Warn("Can't restore cache with error: " + err.Error())
		return err
	}
	return nil
}

func (s *StreamClient) Close() error {
	if err := s.connection.Close(); err != nil {
		s.logger.Fatal("Cant close the connection to nats streaming server: " + err.Error())
		return err
	}
	s.logger.Info("Disconnected from nats streaming server")
	s.database.Disconnect()
	s.logger.Info("Disconnected from DB")
	return nil
}

func (s *StreamClient) Subscribe(subject string) error {
	sub, err := s.connection.Subscribe(subject, func(m *stan.Msg) {
		if validateMessage(m) == 0 {
			s.logger.Info("Valid message recieved")
			o := parseMessage(m)
			if _, found := s.cache.Get(o.Uid); found != true {
				s.cache.Set(o.Uid, o.Data, cache.NoExpiration)
				s.logger.Info("Added to cache with order_uid: " + string(o.Uid))
				if err := s.database.Order().Create(o); err != nil {
					s.logger.Warn("Error while adding to DB with order_uid: " + string(o.Uid))
				} else {
					s.logger.Info("Added to DB with order_uid: " + string(o.Uid))
				}
			} else {
				s.logger.Warn("Order already exists with order_uid: " + string(o.Uid))
			}
		} else {
			s.logger.Warn("Unvalid message recieved")
		}
	})
	if err != nil {
		s.logger.Fatal("Can't subscribe on channel with error: " + err.Error())
		return err
	}
	s.subscription = sub
	s.logger.Info("Subscribed on channel")
	return nil
}

func (s *StreamClient) Unsubscribe() error {
	if err := s.subscription.Unsubscribe(); err != nil {
		s.logger.Warn("Can't unsubscribe from channel with error: " + err.Error())
		return err
	}
	s.logger.Info("Unsubscribed from channel")
	return nil
}

func (s *StreamClient) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *StreamClient) configureDatabase() error {
	db := store.New(s.config.Database)
	if err := db.Connect(); err != nil {
		return err
	}
	s.database = db
	s.logger.Info("Connected to DB")
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

func parseMessage(m *stan.Msg) *models.Order {
	jsonValue := make(map[string]interface{})
	json.Unmarshal(m.Data, &jsonValue)
	o := &models.Order{
		Uid:  "{" + jsonValue["order_uid"].(string) + "}",
		Data: m.Data,
	}
	return o
}

func (s *StreamClient) restoreCacheFromDB() error {
	orders, err := s.database.Order().FindAll()
	if err != nil {
		return err
	}
	for _, v := range orders {
		s.cache.Set(v.Uid, v.Data, cache.NoExpiration)
		s.logger.Info("Added to cache with order_uid: " + string(v.Uid))
	}
	return nil
}
