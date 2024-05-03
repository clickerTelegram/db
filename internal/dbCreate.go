package internal

import (
	"github.com/clickerTelegram/type/db"
	"github.com/go-playground/validator/v10"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/sirupsen/logrus"
)

var Default db.ConfigEnv

func DbCreate() *db.ConfigEnv {
	envFeeder := feeder.DotEnv{Path: ".env"}

	c := config.New()
	c.AddFeeder(envFeeder)
	c.AddStruct(&Default)

	if err := c.Feed(); err != nil {
		logrus.Errorln(err)
	}

	if err := validate(&Default); err != nil {
		logrus.Errorln(err)
	}
	return &Default
}

func validate(c *db.ConfigEnv) error {
	v := validator.New()

	return v.Struct(c)
}
