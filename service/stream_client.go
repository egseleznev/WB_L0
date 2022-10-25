package service

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"l0/service/models"
	"l0/service/store"
	"log"
	"reflect"
	"strings"
	"time"
)

// StreamClient represents an entity of nats streaming server client
// config is specified streaming server client config
// connection is a handle to established connection
// subscription is a handle to active message subscription
// logger is a handle to logging service conditions
// cache is a handle to an application cache
// database is a handle to interact with database
type StreamClient struct {
	config       *Config
	logger       *logrus.Logger
	cache        *cache.Cache
	database     *store.Database
	connection   stan.Conn
	subscription stan.Subscription
}

// New creates and return a streaming server client with the specified config and application cache
func New(config *Config, cache *cache.Cache) *StreamClient {

	return &StreamClient{
		config: config,
		logger: logrus.New(),
		cache:  cache,
	}
}

// Start starts the streaming server client
// establishes database connection, then loads cache from it
// logger is also configured
// return nil if successfully starting
// otherwise return error
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
		s.logger.Fatal("Can't restore cache with error: " + err.Error())
		return err
	}

	return nil
}

// Close closes the connection to nats streaming server
// also disconnects from the database
// return nil if successfully closed
// otherwise return error
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

// Subscribe inputs channel name and subscribes to receive messages from the channel
// return nil if successfully subscribed
// otherwise return error
func (s *StreamClient) Subscribe(subject string) error {

	aw, _ := time.ParseDuration("1s") // receive the message again after 1s if ack not received

	sub, err := s.connection.Subscribe(subject, s.handleMessage,
		stan.DurableName("subscriber_durable"), stan.SetManualAckMode(), stan.AckWait(aw))

	if err != nil {
		s.logger.Fatal("Can't subscribe on channel with error: " + err.Error())
		return err
	}

	s.subscription = sub
	s.logger.Info("Subscribed on channel")

	return nil
}

// Unsubscribe ends a subscription from receiving messages from a streaming server channel
// return nil if successfully closed
// otherwise return error
func (s *StreamClient) Unsubscribe() error {

	if err := s.subscription.Unsubscribe(); err != nil {
		s.logger.Warn("Can't unsubscribe from channel with error: " + err.Error())
		return err
	}

	s.logger.Info("Unsubscribed from channel")

	return nil
}

// configureLogger gets the logging level from the config and sets it to the logger
// return nil if successfully closed
// otherwise return error
func (s *StreamClient) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// configureDatabase configures the database using the database config and connects to it
// return nil if successfully closed
// otherwise return error
func (s *StreamClient) configureDatabase() error {

	db := store.New(s.config.Database)

	if err := db.Connect(); err != nil {
		return err
	}

	s.database = db
	s.logger.Info("Connected to DB")

	return nil
}

// handleMessage handles messages receiving from streaming server
// then if receipt of the message, validation occurs, checking for existence in the cache
// and adding to the database and cache if doesn't exist
func (s *StreamClient) handleMessage(m *stan.Msg) {

	m.Ack() // send an ack of receipt

	if valid := ValidateMessage(m.Data); valid == true { // validate message
		s.logger.Info("Valid message recieved")

		o := ParseMessage(m.Data) // parse message to order structure

		if _, found := s.cache.Get(o.Uid); found != true { // check for presence in the cache

			s.cache.Set(o.Uid, o.Data, cache.NoExpiration) // add order to the cache with no expire
			s.logger.Info("Added to cache with order_uid: " + string(o.Uid))

			if err := s.database.Order().Create(o); err != nil { // add order to the database
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

}

// ValidateMessage checks the recieved message for mismatch
// return true if message valid
// otherwise return false
func ValidateMessage(m []byte) bool {

	var jsonData models.JsonData

	if err := json.Unmarshal(m, &jsonData); err != nil {
		return false
	}

	allFields := reflect.ValueOf(&jsonData).Elem()

	for i := 0; i < allFields.NumField(); i++ {
		if allFields.Field(i).IsZero() && strings.Contains(allFields.Type().Field(i).Tag.Get("validate"), "required") || len(jsonData.OrderUID) != 19 {
			return false
		}
	}

	return true
}

// ParseMessage inputs message received from the channel and unmarshal it into the order
// return order
func ParseMessage(m []byte) *models.Order {

	jsonValue := make(map[string]interface{})

	json.Unmarshal(m, &jsonValue)

	o := &models.Order{
		Uid:  jsonValue["order_uid"].(string),
		Data: m,
	}

	return o
}

// restoreCacheFromDB get all records from the database and add them to the cache
// return nil if successfully added
// otherwise return error
func (s *StreamClient) restoreCacheFromDB() error {

	orders, err := s.database.Order().FindAll()

	if err != nil {
		return err
	}

	for _, v := range orders {
		s.cache.Set(v.Uid, v.Data, cache.NoExpiration)
		s.logger.Info("Restored from db to cache with order_uid: " + string(v.Uid))
	}

	return nil
}
