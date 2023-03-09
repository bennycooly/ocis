package config

// TokenManager is the config for using the reva token manager
type TokenManager struct {
	JWTSecret string `yaml:"jwt_secret" env:"OCIS_JWT_SECRET;STORAGE_PUBLICLINK_JWT_SECRET" desc:"The secret to mint and validate jwt tokens."`
}