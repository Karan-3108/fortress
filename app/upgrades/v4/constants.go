package v4

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v0.1.2-alpha"
	// MainnetUpgradeHeight defines the Fortress mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 257_850
	// TestnetUpgradeHeight defines the Fortress testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 1_200_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/Karan-3108/fortress/releases/download/v0.1.2-alpha/fortress_0.1.2-alpha_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/Karan-3108/fortress/releases/download/v0.1.2-alpha/fortress_0.1.2-alpha_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/Karan-3108/fortress/releases/download/v0.1.2-alpha/fortress_0.1.2-alpha_Linux_arm64.tar.gz","linux/x86_64":"https://github.com/Karan-3108/fortress/releases/download/v0.1.2-alpha/fortress_0.1.2-alpha_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/Karan-3108/fortress/releases/download/v0.1.2-alpha/fortress_0.1.2-alpha_Windows_x86_64.zip"}}'`

	// ExpiredOsmosisClient defines the client ID of the expired Osmosis IBC client
	ExpiredOsmosisClient = "07-tendermint-0"
	// ActiveOsmosisClient defines the client ID of the active Osmosis IBC client
	ActiveOsmosisClient = "07-tendermint-26"
	// ExpiredCosmosHubClient defines the client ID of the expired Cosmos Hub IBC client
	ExpiredCosmosHubClient = "07-tendermint-3"
	// ActiveCosmosHubClient defines the client ID of the active CosmosHub IBC client
	ActiveCosmosHubClient = "07-tendermint-29"
)
