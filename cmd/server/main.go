package main

import (
	"context"

	"github.com/TOIIIA86/image-previewer/internal/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var shaCommit = "local"

const DefaultCacheCapacity = 100

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.With().Str("sha_commit", shaCommit).Logger()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info().Err(err).Msg("Не удалось загрузить файл env")
	}

	cacheCapacity, ok := viper.Get("CACHE_CAPACITY").(int)
	if !ok {
		cacheCapacity = DefaultCacheCapacity
	}

	srv, err := app.NewServer(logger, cacheCapacity)
	if err != nil {
		logger.Fatal().Err(err).Msg("Ошибка старта сервера")
	}

	ctx := log.Logger.WithContext(context.Background())
	if err := srv.Listen(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Не удалось прослушать порт")
	}
}
