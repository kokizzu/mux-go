// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

import (
	"time"
)

type Configuration struct {
	basePath  string `json:"basePath,omitempty"`
	host      string `json:"host,omitempty"`
	userAgent string `json:"userAgent,omitempty"`
	username  string
	password  string
	timeout   time.Duration
}

// ConfigurationOption configures the Mux API Client.
type ConfigurationOption func(*Configuration)

func NewConfiguration(opts ...ConfigurationOption) *Configuration {
	cfg := &Configuration{
		basePath:  "https://api.mux.com",
		userAgent: "Mux Go | 3.3.0",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithBasicAuth(username, password string) ConfigurationOption {
	return func(c *Configuration) {
		c.username = username
		c.password = password
	}
}

func WithTimeout(t time.Duration) ConfigurationOption {
	return func(c *Configuration) {
		c.timeout = t
	}
}

func WithHost(host string) ConfigurationOption {
	return func(c *Configuration) {
		c.host = host
	}
}
