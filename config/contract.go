package config

type Contract struct {
	AlchemyRpcUrl          string `mapstructure:"alchemy_rpc_url" json:"alchemy_rpc_url" yaml:"alchemy_rpc_url"`
	VipcardContractAbi     string `mapstructure:"vipcard_contract_abi" json:"vipcard_contract_abi" yaml:"vipcard_contract_abi"`
	VipcardName            string `mapstructure:"vipcard_name" json:"vipcard_name" yaml:"vipcard_name"`
	VipcardSymbol          string `mapstructure:"vipcard_symbol" json:"vipcard_symbol" yaml:"vipcard_symbol"`
	Erc6551RegistryAddress string `mapstructure:"erc6551_registry_address" json:"erc6551_registry_address" yaml:"erc6551_registry_address"`
	VipcardAccountAddress  string `mapstructure:"vipcard_account_address" json:"vipcard_account_address" yaml:"vipcard_account_address"`
	PrivateKey             string `mapstructure:"private_key" json:"private_key" yaml:"private_key"`
	ChainId                string `mapstructure:"chain_id" json:"chain_id" yaml:"chain_id"`
}
