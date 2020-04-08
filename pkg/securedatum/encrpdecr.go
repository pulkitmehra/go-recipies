package secureddatum

import (
	"fmt"
)

type (
	//Option convinient way to input config
	Option func(c *Config)

	//config struct
	Config struct {
		EncryptionPassword string
		EncryptionSalt     string
		EncryptionBits     int
		EncryptionPoolSize int
	}

	//EnvConfigProvider provider
	EnvConfigProvider struct{}

	//VaultConfigProvider provider
	VaultConfigProvider struct{}
)

var (
	EnvCfgProvider EnvConfigProvider = EnvConfigProvider{}
)

//domain model
type (
	// Encryptor implementation producing Base64-encoded values.
	Encryptor interface {
		Encrypt(value string) (string, error)
	}

	// Decryptor implementation consuming Base64-encoded values.
	// Arbitrary levels of URL encoding should be stripped prior to decryption.
	Decryptor interface {
		Decrypt(value string) (string, error)
	}

	// EncryptorDecryptor implementation consuming/producing Base64-encoded values.
	// Composition of Encryptor and Decryptor interfaces.
	EncryptorDecryptor interface {
		Encryptor
		Decryptor
	}

	encryptedDatum struct {
		*Config
		password    string
		salt        string
		keyStrength int
	}
)

//errors
type (
	encryptorDecryptorError struct {
		val   string
		cause error
	}
	encryptorError encryptorDecryptorError
	decryptorError encryptorDecryptorError
)

func (e encryptorError) Error() string {
	return fmt.Sprintf("Failed to encrypt content %v: %v", e.val, e.cause)
}

func (e decryptorError) Error() string {
	return fmt.Sprintf("Failed to decrypt content %v: %v", e.val, e.cause)
}

// New allocates and configures EncryptorDecryptor based on "singleton" secrets and encryption configuration.
func New(options ...Option) (EncryptorDecryptor, error) {
	config := &Config{}
	for _, o := range options {
		o(config)
	}
	if err := validate(config); err != nil {
		return nil, err
	}
	return &encryptedDatum{
		Config: config,
	}, nil
}

func validate(cfg *Config) error {
	return nil
}

// Encrypt encrypts the value to create a Base64-encoded value.
func (e encryptedDatum) Encrypt(value string) (string, error) {
	return "", nil
}

// Decrypt strips arbitrary levels of URL encoding prior to decryption of Base64-encoded value.
func (e encryptedDatum) Decrypt(value string) (unencrypted string, err error) {
	return "", nil
}

func (EnvConfigProvider) EncryptionPassword() Option {
	return func(c *Config) {
		//c.EncryptionBits = os.Getenv("ENV_EncryptionPassword")
	}
}
func (EnvConfigProvider) EncryptionSalt() Option {
	return func(c *Config) {
		//c.EncryptionBits = os.Getenv("ENV_EncryptionSalt")
	}
}
func (EnvConfigProvider) EncryptionBits() Option {
	return func(c *Config) {
		//c.EncryptionBits = os.Getenv("ENV_EncryptionBits")
	}
}
func (EnvConfigProvider) EncryptionPoolSize() Option {
	return func(c *Config) {
		//c.EncryptionBits = os.Getenv("ENV_EncryptionPoolSize")
	}
}

func (e EnvConfigProvider) All() []Option {
	options := []Option{e.EncryptionPassword(), e.EncryptionBits(), e.EncryptionPoolSize(), e.EncryptionSalt()}
	return options
}
