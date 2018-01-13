package main


type Pay interface {
	pay(user_id int64, money float32) error
}