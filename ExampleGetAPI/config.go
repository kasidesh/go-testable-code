package main

import (
	"os"
	"strconv"
)

type config struct {
	requestTimeout  int
	databaseTimeout int
}

func newConfig() *config {
	return &config{}
}

func (c *config) DatabaseTimeout() int {
	if c.databaseTimeout == 0 {
		timeout, err := strconv.Atoi(os.Getenv("DATABASE_TIMEOUT"))
		if err != nil {
			timeout = 5
		}
		c.databaseTimeout = timeout
	}
	return c.databaseTimeout
}

func (c *config) RequestTimeout() int {
	if c.requestTimeout == 0 {
		timeout, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
		if err != nil {
			timeout = 5
		}
		c.requestTimeout = timeout
	}
	return c.requestTimeout
}
