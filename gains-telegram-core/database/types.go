package database

type HootUser struct {
	ID         float64 `bson:"_id"`
	PrivateKey string  `bson:"private_key"`
	UserName   string  `bson:"user_name"`
	FirstName  string  `bson:"first_name"`
	LastName   string  `bson:"last_name"`
	WalletSet  bool    `bson:"wallet_set"`
	PublicKey  string  `bson:"public_key"`
	TotalXP    int64   `bson:"total_xp"`
}
