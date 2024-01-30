package utils

import (
	"context"
	"net/http"
	"petclinic/prisma/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO:  verficar se o prisma trabalha com Pool de Conexões
func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	return client
}

// mock para testes
type Pool struct {
	minConns int32
	maxConns int32
}

// mock para testes
type ResultSet struct {
	rows int32
}

// mock para testes
func (p *Pool) Exec(ctx context.Context, sql string, arguments ...any) (*ResultSet, error) {
	rows := 0
	return &ResultSet{rows: int32(rows)}, nil
}

func GetPool(databaseURL string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		panic(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	return pool
}

/*
func AppMiddleware(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pool", pool)
		c.Next()
	}
}
*/
