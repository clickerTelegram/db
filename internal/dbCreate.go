package internal

import (
	"fmt"
	"github.com/clickerTelegram/type/db"
	"github.com/go-playground/validator/v10"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Default db.ConfigEnv
var database *gorm.DB

var initializers = []func() error{
	initSettings,
}

func initSettings() error {
	err := database.AutoMigrate(db.Settings{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&db.Settings{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		data := make(map[string]string)

		data["start"] = "Hey, {FIRST_NAME}! Welcome to TapSwap!\nTap on the coin and see your balance rise.\n\nTapSwap is a Decentralized Exchange on the Solana Blockchain. The biggest part of TapSwap Token TAPS distribution will occur among the players here.\n\nGot friends, relatives, co-workers?\nBring them all into the game.\nMore buddies, more coins."
		data["start.button.1"] = "👋 Start now!"
		data["start.button.2"] = "💪💋 Join community"
		data["start.button.3"] = "📑 Help"

		data["help"] = "[Explore the complete guide](https://tapswap.ai/)\n\nTap to Earn:\nTapSwap is an addictive clicker game where you accumulate Shares by tapping the screen.\n\nLeagues:\nClimb the ranks by earning more Shares and outperforming others in the leagues.\n\nBoosts:\nUnlock boosts and complete tasks to maximize your Shares earnings.\n\nFriends:\nInvite others and both of you will receive bonuses. Assist your friends in advancing to higher leagues for bigger Shares rewards.\n\nThe Purpose:\nCollect as many Shares as possible and exchange them for TAPS, TapSwap Token on Solana Blockchain.\n\nType /help to access this guide."
		data["help.button.1"] = "👋 Start now!"

		data["profile"] = "{USERNAME} profile\n\n{LEAGUE} League\nTotal score: {TOTAL_SCORE}\nBalance: {BALANCE}\n\n/profile for personal stats"
		data["profile.button.1"] = "👥 Invite friends"
		data["profile.button.2"] = "👋 Play"

		data["friend"] = "Share with your friends and earn bonuses for each friend you invite and for their activity:\n\nYour referral link: `{LINK_REF}`"

		data["socials"] = "Join our socials so you do not miss any important news or updates."
		data["socials.button.1"] = "TapSwap Community"
		data["socials.button.2"] = "TapSwap on X"
		data["socials.button.3"] = "TapSwap site"
		data["socials.button.4"] = "👋 Play"

		var dbSettings []db.Settings
		for index, value := range data {
			st := db.Settings{}
			st.Type = index
			st.Value = value

			dbSettings = append(dbSettings, st)
		}

		database.Create(dbSettings)
	}

	return nil

}

func DbCreate() {
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
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		Default.PostgresUrl, Default.PostgresPort, Default.PostgresUsername, Default.PostgresPassword, Default.PostgresDb, Default.SslMode, Default.TimeZone)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("error 1")
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		panic("error 2")
	}
	sqlDb.SetMaxIdleConns(Default.SetMaxIdleConns)
	sqlDb.SetMaxOpenConns(Default.SetMaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(Default.SetConnMaxLifetime) * time.Hour)

	database = dbClient

	for _, initialize := range initializers {
		if err := initialize(); err != nil {
			logrus.WithError(err).Error("failed to initialize")
		}
	}
}

func validate(c *db.ConfigEnv) error {
	v := validator.New()

	return v.Struct(c)
}
