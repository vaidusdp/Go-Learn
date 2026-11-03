package main

import "github.com/sirupsen/logrus"

func main(){
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("Info Msg")
	log.Warn("Warn Msg")
	log.Error("Error Msg")

	log.WithFields((logrus.Fields{
		"username": "Jhon Doe",
		"method": "GET",
	})).Info("User Logged In")
}