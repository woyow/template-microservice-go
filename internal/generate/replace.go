package generate

import (
	"bytes"
	"github.com/woyow/template-microservice-go/config"
)

type ReplaceWord struct {
	OldWord string
	NewWord string
}

type Words struct {
	ModuleName ReplaceWord
	ServiceName ReplaceWord
	AppEnv ReplaceWord
	PostgresEnv ReplaceWord
	RedisEnv ReplaceWord
	RabbitEnv ReplaceWord
	GoVersion ReplaceWord
}

func NewWords(cfg *config.Config) *Words {
	return &Words{
		ModuleName: ReplaceWord{
			OldWord: "{{MODULE_NAME}}",
			NewWord: cfg.ModuleName,
		},
		ServiceName: ReplaceWord{
			OldWord: "{{SERVICE_NAME}}",
			NewWord: cfg.ServiceName,
		},
		AppEnv: ReplaceWord{
			OldWord: "{{APP_ENV}}",
			NewWord: "APP_ENV: \"\"",
		},
		PostgresEnv: ReplaceWord{
			OldWord: "{{POSTGRES_ENV}}",
			NewWord: "PG_PASSWORD: \"\"",
		},
		RedisEnv: ReplaceWord{
			OldWord: "{{REDIS_ENV}}",
			NewWord: "REDIS_PASSWORD: \"\"",
		},
		RabbitEnv: ReplaceWord{
			OldWord: "{{RABBIT_ENV}}",
			NewWord: "RABBIT_PASSWORD: \"\"",
		},
		GoVersion: ReplaceWord{
			OldWord: "{{GO_VERSION}}",
			NewWord: "1.17",
		},
	}
}

func ReplaceWords(words *Words, target *[]byte) []byte {
	*target = bytes.Replace(*target, []byte(words.ModuleName.OldWord), []byte(words.ModuleName.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.ServiceName.OldWord), []byte(words.ServiceName.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.AppEnv.OldWord), []byte(words.AppEnv.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.PostgresEnv.OldWord), []byte(words.PostgresEnv.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.RedisEnv.OldWord), []byte(words.RedisEnv.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.RabbitEnv.OldWord), []byte(words.RabbitEnv.NewWord), -1)
	*target = bytes.Replace(*target, []byte(words.GoVersion.OldWord), []byte(words.GoVersion.NewWord), -1)

	return *target
}