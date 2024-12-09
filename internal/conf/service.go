package conf

type ServiceConfig struct {
	KeyPair KeyPairConfig `yaml:"key_pair"`
}

type KeyPairConfig struct {
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}
