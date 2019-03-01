package controllers

import "report/delivery"

var Consumer delivery.Delivery

const keyApp = "91fc328d734261ff45775ebe00fdb13c"
const token = "0859930bb7d2177e2fd3a1a8c4bed14f253ba2f93b044e4b5e9e59241b9f61d3"
const username = "eban6"

func init() {
	// var keyApp, token, username string
	// fmt.Print("Enter keyApp: ")
	// fmt.Scanln(&keyApp)
	// fmt.Print("Enter token: ")
	// fmt.Scanln(&token)
	// fmt.Print("Enter username: ")
	// fmt.Scanln(&username)
	Consumer = delivery.NewDelivery(keyApp, token, username)
}
