package config

type Contract struct {
	AlchemyRpcUrl string `mapstructure:"alchemy_rpc_url" json:"alchemy_rpc_url" yaml:"alchemy_rpc_url"`
}
