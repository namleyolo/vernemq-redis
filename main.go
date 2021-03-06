package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bbf5099c19728139a0c0bc77de98e241ddbe13a6f923b3557a6a035e45f05d1c27ec728af982b629b682d47fc043db5b48e2573880f3da3be527782ab7246181",
		DB:       0,
	})
	log.Print(client.Ping())
	password := "namle1234567890"
	hash, _ := HashPassword(password)

	maps := make([]map[string]string, 0)
	mapSub := make([]map[string]string, 0)

	//type Value map[string]interface{}

	pattern := map[string]string{"pattern": "a/b/#"}
	patterns := map[string]string{"pattern": "d/e/f/#"}
	patternSub := map[string]string{"pattern": "hg/v1/#"}
	maps = append(maps, pattern)
	maps = append(maps, patterns)
	mapSub = append(mapSub, patternSub)
	mapSub = append(mapSub, pattern)

	value := map[string]interface{}{"passhash": hash, "publish_acl": maps, "subscribe_acl": mapSub}

	key := []string{"", "namle", "namle12345678"}

	bTest, err := json.Marshal(key)
	bNono, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = client.Set(string(bTest),
		bNono, 0).Err()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
