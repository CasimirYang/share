package main

import (
	"share/log"
)

func main() {
	log.SugarLogger.Info("SugarLogger info message")
	log.SugarLogger.Debug("debug message")
	log.SugarLogger.Error("error message")
	fun()
	defer log.SugarLogger.Sync()
}

func fun() {
	log.SugarLogger.Info("iner 01")
	fun2()
}

func fun2() {
	log.SugarLogger.Info("iner 02")
}
