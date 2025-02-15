package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	authn "github.com/tx7do/kratos-authn/engine"
	authnEngine "github.com/tx7do/kratos-authn/engine"

	userV1 "kratos-cms/api/gen/go/user/service/v1"
)

type UserToken struct {
	log           *log.Helper
	rdb           *redis.Client
	authenticator authnEngine.Authenticator

	accessTokenKeyPrefix  string
	refreshTokenKeyPrefix string
}

func NewUserToken(
	rdb *redis.Client,
	authenticator authnEngine.Authenticator,
	logger log.Logger,
	accessTokenKeyPrefix string,
	refreshTokenKeyPrefix string,
) *UserToken {
	l := log.NewHelper(log.With(logger, "module", "user-token/cache"))
	return &UserToken{
		rdb:                   rdb,
		log:                   l,
		authenticator:         authenticator,
		accessTokenKeyPrefix:  accessTokenKeyPrefix,
		refreshTokenKeyPrefix: refreshTokenKeyPrefix,
	}
}

// createAccessJwtToken 生成JWT访问令牌
func (r *UserToken) createAccessJwtToken(_ string, userId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(userId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}

	return signedToken
}

// createRefreshToken 生成刷新令牌
func (r *UserToken) createRefreshToken() string {
	return uuid.New().String()
}

// GenerateToken 创建令牌
func (r *UserToken) GenerateToken(ctx context.Context, user *userV1.User) (accessToken string, refreshToken string, err error) {
	if accessToken = r.createAccessJwtToken(user.GetUserName(), user.GetId()); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		return
	}

	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// GenerateAccessToken 创建访问令牌
func (r *UserToken) GenerateAccessToken(ctx context.Context, userId uint32, userName string) (accessToken string, err error) {
	if accessToken = r.createAccessJwtToken(userName, userId); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, userId, accessToken, 0); err != nil {
		return
	}

	return
}

// GenerateRefreshToken 创建刷新令牌
func (r *UserToken) GenerateRefreshToken(ctx context.Context, user *userV1.User) (refreshToken string, err error) {
	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// RemoveToken 移除所有令牌
func (r *UserToken) RemoveToken(ctx context.Context, userId uint32) error {
	var err error
	if err = r.deleteAccessTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user access token failed: [%v]", err)
	}

	if err = r.deleteRefreshTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user refresh token failed: [%v]", err)
	}

	return err
}

// GetAccessToken 获取访问令牌
func (r *UserToken) GetAccessToken(ctx context.Context, userId uint32) string {
	return r.getAccessTokenFromRedis(ctx, userId)
}

// GetRefreshToken 获取刷新令牌
func (r *UserToken) GetRefreshToken(ctx context.Context, userId uint32) string {
	return r.getRefreshTokenFromRedis(ctx, userId)
}

// IsExistAccessToken 访问令牌是否存在
func (r *UserToken) IsExistAccessToken(ctx context.Context, userId uint32) bool {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	n, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return n > 0
}

// IsExistRefreshToken 刷新令牌是否存在
func (r *UserToken) IsExistRefreshToken(ctx context.Context, userId uint32) bool {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	n, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return n > 0
}

func (r *UserToken) setAccessTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	return r.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserToken) getAccessTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("get redis user access token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserToken) deleteAccessTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	return r.rdb.Del(ctx, key).Err()
}

func (r *UserToken) setRefreshTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	return r.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserToken) getRefreshTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("get redis user refresh token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserToken) deleteRefreshTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	return r.rdb.Del(ctx, key).Err()
}
