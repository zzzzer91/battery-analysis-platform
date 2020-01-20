package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"battery-analysis-platform/pkg/security"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 用户类型 Type
const (
	// 超级用户
	UserTypeSuperUser = 64
	// 普通用户
	UserTypeCommonUser = 0
)

// 用户状态 Status
const (
	UserStatusForbiddenLogin = 0
	UserStatusNormal         = 1
)

const (
	// redis 的键前缀
	redisKeyPrefix = "ubattery:users:"
	// 缓存中 key 过期时间
	redisKeyExpiration = time.Hour * 3
)

type User struct {
	Name     string `json:"userName" bson:"name"`
	Password string `json:"-" bson:"password"`
	Type     int    `json:"userType" bson:"type"`
	// *string 让其 json 时可以返回 null，否则只能返回字符串零值
	AvatarName    *string    `json:"avatarName" bson:"avatarName"`
	LastLoginTime jtime.Time `json:"lastLoginTime" bson:"lastLoginTime"`
	Comment       string     `json:"comment" bson:"comment"`
	LoginCount    int        `json:"loginCount" bson:"loginCount"`
	Status        int        `json:"userStatus" bson:"status"`
	CreateTime    jtime.Time `json:"createTime" bson:"createTime"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	s, err := security.GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	user.Password = s
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := security.CheckPasswordHash(user.Password, password)
	return err == nil
}

func (user *User) CheckStatusOk() bool {
	return user.Status == UserStatusNormal
}

func CreateUser(name, password, comment string) (*User, error) {
	now := jtime.Now()
	user := User{
		Name:          name,
		Comment:       comment,
		CreateTime:    now,
		LastLoginTime: now,
		Status:        1,
	}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	err = insertMongoCollection(mongoCollectionUser, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(name string) (*User, error) {
	var user User
	collection := db.Mongo.Collection(mongoCollectionUser)
	filter := bson.M{"name": name}
	projection := bson.M{"_id": false} // 注意 _id 默认会返回，需要手动过滤
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ListCommonUser() ([]User, error) {
	collection := db.Mongo.Collection(mongoCollectionUser)
	filter := bson.M{"type": bson.M{"$ne": UserTypeSuperUser}} // 过滤记录
	projection := bson.M{"_id": false}                         // 过滤字段
	ctx := context.TODO()
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	users := make([]User, 0)
	for cur.Next(ctx) {
		result := User{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, result)
	}
	_ = cur.Close(ctx)
	return users, nil
}

func SaveUserLoginTimeAndCount(user *User) error {
	collection := db.Mongo.Collection(mongoCollectionUser)
	filter := bson.M{"name": user.Name} // 过滤记录
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	update := bson.M{"$set": bson.M{
		"lastLoginTime": user.LastLoginTime,
		"loginCount":    user.LoginCount,
	}}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func SaveUserChange(user *User) error {
	collection := db.Mongo.Collection(mongoCollectionUser)
	filter := bson.M{"name": user.Name} // 过滤记录
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	update := bson.M{"$set": bson.M{
		"comment": user.Comment,
		"status":  user.Status,
	}}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func ChangeUserPassword(userName, password string) error {
	collection := db.Mongo.Collection(mongoCollectionUser)
	filter := bson.M{"name": userName} // 过滤记录
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	s, err := security.GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{
		"password": s,
	}}
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}

// ---------------------------cache---------------------------

func SaveUserToCache(user *User) error {
	// 存储 JSON 序列化的数据
	jd, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return db.Redis.Set(redisKeyPrefix+user.Name, jd, redisKeyExpiration).Err()
}

func DeleteUserFromCache(name string) error {
	return db.Redis.Del(redisKeyPrefix + name).Err()
}

func GetUserFromCache(name string) (*User, error) {
	val, err := db.Redis.Get(redisKeyPrefix + name).Bytes()
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(val, &user)
	if err != nil {
		return nil, err
	}
	// 刷新 key 的过期时间
	db.Redis.Expire(redisKeyPrefix+name, redisKeyExpiration)
	return &user, nil
}
