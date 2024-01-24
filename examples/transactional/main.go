package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	serv := &service{
		db: NewDatabase(),
	}
	user := User{Name: "Marc"}
	err := serv.createNewUser(user)
	if err != nil {
		log.Fatalf("Error creating user %+v: %s\n", user, err)
	}
}

type service struct {
	db *Database
}

// START OMIT
func (s *service) createNewUser(user User) error {
	err := s.db.RunInTransaction(func(h *DatabaseHandle) error {
		// Your business logic in this anonymous function // HL
		log.Printf("Store user %v using database transaction: %s\n", user, h.GetTranactionId())
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) RunInTransaction(f func(h *DatabaseHandle) error) error {
	trx := db.CreateTransaction()
	err := f(trx.ToHandle()) // HL
	if err != nil {
		_ = db.RollbackTransaction(trx)
		return err
	}
	return db.CommitTransaction(trx)
}

// END OMIT

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) CreateTransaction() *Transaction {
	uid := uuid.New().String()
	fmt.Printf("Create transaction %s\n", uid)
	return &Transaction{
		uid: uid,
	}
}

func (db *Database) CommitTransaction(trx *Transaction) error {
	fmt.Printf("Commit transaction %s\n", trx.uid)
	return nil
}

func (db *Database) RollbackTransaction(trx *Transaction) error {
	fmt.Printf("Rollback transaction %s\n", trx.uid)
	return nil
}

type Transaction struct {
	uid string
}

func (t *Transaction) ToHandle() *DatabaseHandle {
	return &DatabaseHandle{
		transactionUid: t.uid,
	}
}

type DatabaseHandle struct {
	transactionUid string
}

func (dh *DatabaseHandle) GetTranactionId() string {
	return dh.transactionUid
}

type User struct {
	Name string
}
