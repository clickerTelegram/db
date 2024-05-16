package postgres

import (
	"github.com/clickerTelegram/type/db"
	"github.com/clickerTelegram/type/entity"
	"github.com/go-playground/validator/v10"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Default db.ConfigEnv
var database *gorm.DB

var initializers = []func() error{
	initSettings,
	initAccessToken,
	initBoosts,
	initChargeLevel,
	initClaims,
	initEnergyLevel,
	initLastTapPlayer,
	initLeagues,
	initMissions,
	initMissionsItemsDB,
	initMissionsCompleted,
	initOnlineWeb,
	initPlayer,
	initPlayerLevel,
	initReferralPlayerData,
	initReferralRewards,
	initStatsData,
	initTabBots,
	initTapLevel,
	initPlayerData,
}

func initTapLevel() error {
	err := database.AutoMigrate(entity.TapLevelsDB{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&entity.TapLevelsDB{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		database.Create(GetDataTapLevel())
	}
	return err
}

func initTabBots() error {
	return database.AutoMigrate(entity.TabBotsDB{})
}

func initStatsData() error {
	return database.AutoMigrate(entity.StatsData{})
}

func initReferralRewards() error {
	err := database.AutoMigrate(entity.ReferralRewardsDB{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&entity.ReferralRewardsDB{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		database.Create(GetDataReferences())
	}
	return err
}

func initReferralPlayerData() error {
	return database.AutoMigrate(entity.ReferralPlayerData{})
}

func initPlayerLevel() error {
	return database.AutoMigrate(entity.PlayerLevelsDB{})
}

func initPlayerData() error {
	return database.AutoMigrate(entity.PlayerData{})
}

func initPlayer() error {
	return database.AutoMigrate(entity.Player{})
}

func initOnlineWeb() error {
	return database.AutoMigrate(entity.OnlineWebPlayer{})
}

func initMissionsCompleted() error {
	return database.AutoMigrate(entity.MissionsComplete{})
}

func initMissions() error {
	return database.AutoMigrate(entity.MissionsDB{})
}

func initMissionsItemsDB() error {
	return database.AutoMigrate(entity.MissionsItemsDB{})
}

func initLeagues() error {
	err := database.AutoMigrate(entity.LeaguesDB{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&entity.LeaguesDB{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		database.Create(GetDataLeagues())
	}
	return err
}

func initLastTapPlayer() error {
	return database.AutoMigrate(entity.LastTapPlayer{})
}

func initEnergyLevel() error {
	err := database.AutoMigrate(entity.EnergyLevelsDB{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&entity.EnergyLevelsDB{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		database.Create(GetDataEnergyLevel())
	}
	return err
}

func initAccessToken() error {
	return database.AutoMigrate(entity.AccessToken{})
}

func initBoosts() error {
	return database.AutoMigrate(entity.Boosts{})
}

func initChargeLevel() error {
	err := database.AutoMigrate(entity.ChargeLevelsDB{})
	if err != nil {
		return err
	}

	var count int64

	err = database.Model(&entity.ChargeLevelsDB{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		database.Create(GetDataChargeLevel())
	}
	return err

}

func initClaims() error {
	return database.AutoMigrate(entity.ClaimsData{})
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

		data["start"] = "Hey, [{USERNAME|FIRST_NAME}] ! Welcome to TapSwap!\nTap on the coin and see your balance rise.\n\nTapSwap is a Decentralized Exchange on the Solana Blockchain. The biggest part of TapSwap Token TAPS distribution will occur among the players here.\n\nGot friends, relatives, co-workers?\nBring them all into the webView.\nMore buddies, more coins."
		data["start.button.1"] = "👋 Start now!"
		data["start.button.2"] = "💪💋 Join community"
		data["start.button.3"] = "📑 Help"
		data["start.image"] = "https://i.ibb.co/K0NPrpC/photo-2024-05-02-20-40-38.jpg"

		data["help"] = "[Explore the complete guide](https://tapswap.ai/)\n\nTap to Earn:\nTapSwap is an addictive clicker webView where you accumulate Shares by tapping the screen.\n\nLeagues:\nClimb the ranks by earning more Shares and outperforming others in the leagues.\n\nBoosts:\nUnlock boosts and complete tasks to maximize your Shares earnings.\n\nFriends:\nInvite others and both of you will receive bonuses. Assist your friends in advancing to higher leagues for bigger Shares rewards.\n\nThe Purpose:\nCollect as many Shares as possible and exchange them for TAPS, TapSwap Token on Solana Blockchain.\n\nType /help to access this guide."
		data["help.button.1"] = "👋 Start now!"
		data["help.image"] = "https://i.ibb.co/Mg6ndcs/photo-2024-05-02-20-40-37.jpg"

		data["profile"] = "[{USERNAME|FIRST_NAME}] profile\n\n[{LEAGUE}]\nTotal score: [{TOTAL_SCORE}]\nBalance: [{BALANCE}]\n\n/profile for personal stats"
		data["profile.button.1"] = "👥 Invite friends"
		data["profile.button.2"] = "👋 Play"

		data["friend"] = "Share with your friends and earn bonuses for each friend you invite and for their activity:\n\nYour referral link: `{LINK_REF}`"

		data["socials"] = "Join our socials so you do not miss any important news or updates."
		data["socials.button.1"] = "TapSwap Community"
		data["socials.button.1.url"] = ""
		data["socials.button.2"] = "TapSwap on X"
		data["socials.button.2.url"] = ""
		data["socials.button.3"] = "TapSwap site"
		data["socials.button.3.url"] = ""
		data["socials.button.4"] = "👋 Play"

		var dbSettings []db.Settings
		st := db.Settings{}
		for index, value := range data {

			st.Type = index
			st.Value = value

			dbSettings = append(dbSettings, st)
			st = db.Settings{}
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
		logrus.Panic(err)
	}

	if err := validate(&Default); err != nil {
		logrus.Panic(err)
	}

	database = ConnectPostgres(Default)

	for _, initialize := range initializers {
		if err := initialize(); err != nil {
			logrus.Panic(err)
		}
	}

}

func validate(c *db.ConfigEnv) error {
	v := validator.New()

	return v.Struct(c)
}
