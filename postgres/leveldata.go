package postgres

import "github.com/clickerTelegram/type/entity"

func GetDataEnergyLevel() []entity.EnergyLevelsDB {

	return []entity.EnergyLevelsDB{
		{
			Limit: 500,
			Price: 50,
		},
		{
			Limit: 1000,
			Price: 200,
		},
		{
			Limit: 1500,
			Price: 500,
		},
		{
			Limit: 2000,
			Price: 1000,
		},
		{
			Limit: 2500,
			Price: 2000,
		},
		{
			Limit: 3000,
			Price: 4000,
		},
		{
			Limit: 3500,
			Price: 8000,
		},
		{
			Limit: 4000,
			Price: 16000,
		},
		{
			Limit: 4500,
			Price: 25000,
		},
		{
			Limit: 5000,
			Price: 50000,
		},
		{
			Limit: 5500,
			Price: 100000,
		},
		{
			Limit: 6000,
			Price: 200000,
		},
		{
			Limit: 6500,
			Price: 300000,
		},
		{
			Limit: 7000,
			Price: 400000,
		},
		{
			Limit: 7500,
			Price: 500000,
		},
		{
			Limit: 8000,
			Price: 600000,
		},
		{
			Limit: 8500,
			Price: 700000,
		},
		{
			Limit: 9000,
			Price: 800000,
		},
		{
			Limit: 9500,
			Price: 900000,
		},
		{
			Limit: 10000,
			Price: 1000000,
		},
	}
}

func GetDataChargeLevel() []entity.ChargeLevelsDB {

	return []entity.ChargeLevelsDB{
		{
			Rate:  1,
			Price: 500,
		},
		{
			Rate:  2,
			Price: 2000,
		},
		{
			Rate:  3,
			Price: 10000,
		},
		{
			Rate:  4,
			Price: 100000,
		},
		{
			Rate:  5,
			Price: 250000,
		},
	}

}

func GetDataTapLevel() []entity.TapLevelsDB {

	return []entity.TapLevelsDB{
		{
			Rate:   1,
			Energy: 1,
			Price:  50,
		},
		{
			Rate:   2,
			Energy: 2,
			Price:  200,
		},
		{
			Rate:   3,
			Energy: 3,
			Price:  500,
		},
		{
			Rate:   4,
			Energy: 4,
			Price:  1000,
		},
		{
			Rate:   5,
			Energy: 5,
			Price:  2000,
		},
		{
			Rate:   6,
			Energy: 6,
			Price:  4000,
		},
		{
			Rate:   7,
			Energy: 7,
			Price:  8000,
		},
		{
			Rate:   8,
			Energy: 8,
			Price:  16000,
		},
		{
			Rate:   9,
			Energy: 9,
			Price:  25000,
		},
		{
			Rate:   10,
			Energy: 10,
			Price:  50000,
		},
		{
			Rate:   11,
			Energy: 11,
			Price:  100000,
		},
		{
			Rate:   12,
			Energy: 12,
			Price:  200000,
		},
		{
			Rate:   13,
			Energy: 13,
			Price:  300000,
		},
		{
			Rate:   14,
			Energy: 14,
			Price:  400000,
		},
		{
			Rate:   15,
			Energy: 15,
			Price:  500000,
		},
		{
			Rate:   16,
			Energy: 16,
			Price:  600000,
		},
		{
			Rate:   17,
			Energy: 17,
			Price:  700000,
		},
		{
			Rate:   18,
			Energy: 18,
			Price:  800000,
		},
		{
			Rate:   19,
			Energy: 19,
			Price:  900000,
		},
		{
			Rate:   20,
			Energy: 20,
			Price:  1000000,
		},
	}

}

func GetDataLeagues() []entity.LeaguesDB {

	return []entity.LeaguesDB{
		{
			Name:           "wood",
			Title:          "Wood League",
			Score:          0,
			Reward:         0,
			RewardReferral: 0,
		},
		{
			Name:           "bronze",
			Title:          "Bronze League",
			Score:          1,
			Reward:         1000,
			RewardReferral: 2000,
		},
		{
			Name:           "silver",
			Title:          "Silver League",
			Score:          5000,
			Reward:         5000,
			RewardReferral: 5000,
		},
		{
			Name:           "gold",
			Title:          "Gold League",
			Score:          50000,
			Reward:         10000,
			RewardReferral: 10000,
		},
		{
			Name:           "platinum",
			Title:          "Platinum League",
			Score:          250000,
			Reward:         30000,
			RewardReferral: 15000,
		},
		{
			Name:           "diamond",
			Title:          "Diamond League",
			Score:          500000,
			Reward:         50000,
			RewardReferral: 25000,
		},
		{
			Name:           "master",
			Title:          "Master League",
			Score:          1000000,
			Reward:         100000,
			RewardReferral: 50000,
		},
		{
			Name:           "grandmaster",
			Title:          "Grandmaster League",
			Score:          2500000,
			Reward:         250000,
			RewardReferral: 125000,
		},
		{
			Name:           "elite",
			Title:          "Elite League",
			Score:          5000000,
			Reward:         500000,
			RewardReferral: 250000,
		},
		{
			Name:           "legendary",
			Title:          "Legendary League",
			Score:          10000000,
			Reward:         1000000,
			RewardReferral: 500000,
		},
		{
			Name:           "mythic",
			Title:          "Mythic League",
			Score:          50000000,
			Reward:         5000000,
			RewardReferral: 2500000,
		},
	}

}

func GetDataReferences() []entity.ReferralRewardsDB {

	return []entity.ReferralRewardsDB{
		{
			Count:  0,
			Reward: 0,
		},
		{
			Count:  1,
			Reward: 2500,
		},
		{
			Count:  3,
			Reward: 50000,
		},
		{
			Count:  10,
			Reward: 200000,
		},
		{
			Count:  25,
			Reward: 250000,
		},
		{
			Count:  50,
			Reward: 300000,
		},
		{
			Count:  100,
			Reward: 500000,
		},
		{
			Count:  500,
			Reward: 2000000,
		},
		{
			Count:  1000,
			Reward: 2500000,
		},
		{
			Count:  10000,
			Reward: 10000000,
		},
		{
			Count:  100000,
			Reward: 100000000,
		},
	}

}
