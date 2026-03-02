package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"andy": {
		AuthToken: "123ABC",
		Username:  "andy",
	},
	"bob": {
		AuthToken: "456DEF",
		Username:  "bob",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"andy": {
		Coins:    100,
		Username: "andy",
	},
	"bob": {
		Coins:    200,
		Username: "bob",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1) // Simulate DB Call

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
