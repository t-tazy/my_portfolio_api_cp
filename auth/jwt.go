package auth

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/t-tazy/my_portfolio_api/clock"
	"github.com/t-tazy/my_portfolio_api/entity"
)

// JWTの独自クレーム名
const (
	RoleKey     = "role"
	UserNameKey = "user_name"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

// 鍵として読み込んだデータを保持する
type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker // JWTの時刻情報を操作
}

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c
	return j, nil
}

// 単一のキーを解析
func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

// 引数で渡されたユーザーに対して署名済みのJWTを発行する
func (j *JWTer) GenerateToken(ctx context.Context, u *entity.User) ([]byte, error) {
	token, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/t-tazy/my_portfolio_api`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(RoleKey, u.Role).     // 独自クレーム
		Claim(UserNameKey, u.Name). // 独自クレーム
		Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}
	// Redisに保存
	if err := j.Store.Save(ctx, token.JwtID(), u.ID); err != nil {
		return nil, err
	}

	// 署名
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}
