package models

// Order represents the storage model of the received message in the Сache and in the Database
// Uid is unique record identifier
// Data is set of bytes to store json data
type Order struct {
	Uid  string
	Data []byte
}
