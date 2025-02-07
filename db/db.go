package db

import (
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

import (
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // No password set
		DB:       0,  // Default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
