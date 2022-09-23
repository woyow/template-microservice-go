package structure

import (
	"github.com/woyow/template-microservice-go/config"
)

type (
	File struct {
		Name string
		Path string
		Template string
	}

	Dirs struct {
	CmdDir
	ConfigDir
	DBDir
	InternalDir
	PkgDir
	ReadmeFile		File
	EnvExampleFile 	File
	GitIgnoreFile  	File
	GoModFile      	File
	DockerFile     	File
}

	// module/cmd
	CmdDir struct {
		Name string
		CmdServiceDir
	}

	// module/cmd/service
	CmdServiceDir struct {
		Name	string
		MainFile File
	}

	// module/config
	ConfigDir struct {
		Name string
		ConfigFile File
		LocalConfigYamlFile File
	}

	// module/db
	DBDir struct {
		Name string
		MigrationsDir
	}

	// module/db/migrations
	MigrationsDir struct {
		Name string
		ReadmeFile File
	}

	// module/internal
	InternalDir struct {
		Name string
		ReadmeFile File
		AppDir
		EntityDir
		QueueDir
		ServiceDir
		StorageDir
		TransportDir
	}

	// module/pkg
	PkgDir struct {
		Name string
		ReadmeFile File
	}

	AppDir struct {
		Name string
		AppFile File
		LoggerFile File
		MigrateFile File
	}

	EntityDir struct {
		Name string
		ReadmeFile File
	}

	QueueDir struct {
		Name string
		NatsDir
		RedisMQDir
	}

	NatsDir struct {
		Name string
		NatsFile File
		PublisherTestsFile File
		SubscriberTestsFile File
	}

	RedisMQDir struct {
		Name string
		RedisMQFile File
		ProducerTestsFile File
		ConsumerTestsFile File
	}

	ServiceDir struct {
		Name string
		ServiceFile File
		ReadmeFile File
	}

	StorageDir struct {
		Name string
		ReadmeFile File
		StorageFile File
		PsqlDir
		RedisDir
	}

	PsqlDir struct{
		Name string
		ConfigFile File
		PsqlFile File
	}

	RedisDir struct{
		Name string
		ConfigFile File
		RedisFile File
	}

	TransportDir struct {
		Name string
		GrpcDir
		HttpDir
	}

	GrpcDir struct {
		Name string
		ServerFile File
	}

	HttpDir struct {
		Name string
		ServerFile File
		HandlerDir
	}

	HandlerDir struct{
		Name string
		HandlerFile File
		ResponseFile File
	}
)

func NewStructure(cfg *config.Config) *Dirs {
	return &Dirs{
		ReadmeFile: File{
			Name: "README.md",
			Path: cfg.ProjectPath + "/",
			Template: "./src/templates/README.md",
		},
		EnvExampleFile: File{
			Name: ".env.example",
			Path: cfg.ProjectPath + "/",
			Template: "./src/templates/.env.example.template",
		},
		GitIgnoreFile: File{
			Name: ".gitignore",
			Path: cfg.ProjectPath + "/",
			Template: "./src/templates/.gitignore.template",
		},
		GoModFile: File{
			Name: "go.mod",
			Path: cfg.ProjectPath + "/",
			Template: "./src/templates/go.mod.template",
		},
		DockerFile: File{
			Name: "Dockerfile",
			Path: cfg.ProjectPath + "/",
			Template: "./src/templates/Dockerfile.template",
		},
		CmdDir: CmdDir{
			Name: "cmd",
			CmdServiceDir: CmdServiceDir{
				Name: cfg.ServiceName,
				MainFile: File{
					Name: "main.go",
					Path: cfg.ProjectPath + "/cmd/" + cfg.ServiceName + "/",
					Template: "./src/templates/cmd/main.go.template",
				},
			},
		},
		ConfigDir: ConfigDir{
			Name: "config",
			ConfigFile: File{
				Name: "config.go",
				Path: cfg.ProjectPath + "/config/",
				Template: "./src/templates/config/config.go.template",
			},
			LocalConfigYamlFile: File{
				Name: "local.yaml",
				Path: cfg.ProjectPath + "/config/",
				Template: "./src/templates/config/local.yaml.template",
			},
		},
		DBDir: DBDir{
			Name: "db",
			MigrationsDir: MigrationsDir{
				Name: "migrations",
				ReadmeFile: File{
					Name: "README.md",
					Path: cfg.ProjectPath + "/db/migrations/",
					Template: "./src/templates/db/migrations/README.md",
				},
			},
		},
		InternalDir: InternalDir{
			Name: "internal",
			ReadmeFile: File{
				Name: "README.md",
				Path: cfg.ProjectPath + "/internal/",
				Template: "./src/templates/internal/README.md",
			},
			AppDir: AppDir{
				Name: "app",
				AppFile: File{
					Name: "app.go",
					Path: cfg.ProjectPath + "/internal/app/",
					Template: "./src/templates/internal/app/app.go.template",
				},
				LoggerFile: File{
					Name: "logger.go",
					Path: cfg.ProjectPath + "/internal/app/",
					Template: "./src/templates/internal/app/logger.go.template",
				},
				MigrateFile: File{
					Name: "migrate.go",
					Path: cfg.ProjectPath + "/internal/app/",
					Template: "./src/templates/internal/app/migrate.go.template",
				},
			},
			EntityDir: EntityDir{
				Name: "entity",
				ReadmeFile: File{
					Name: "README.md",
					Path: cfg.ProjectPath + "/internal/entity/",
					Template: "./src/templates/internal/entity/README.md",
				},
			},
			QueueDir: QueueDir{
				Name: "queue",
				NatsDir: NatsDir{
					Name: "nats",
					NatsFile: File{
						Name: "nats.go",
						Path: cfg.ProjectPath + "/internal/queue/nats/",
						Template: "./src/templates/internal/queue/nats/nats.go.template",
					},
					PublisherTestsFile: File{
						Name: "publisher_tests.go",
						Path: cfg.ProjectPath + "/internal/queue/nats/",
						Template: "./src/templates/internal/queue/nats/publisher_tests.go.template",
					},
					SubscriberTestsFile: File{
						Name: "subscriber_tests.go",
						Path: cfg.ProjectPath + "/internal/queue/nats/",
						Template: "./src/templates/internal/queue/nats/subscriber_tests.go.template",
					},
				},
				RedisMQDir: RedisMQDir{
					Name: "redismq",
					RedisMQFile: File{
						Name: "redismq.go",
						Path: cfg.ProjectPath + "/internal/queue/redismq/",
						Template: "./src/templates/internal/queue/redismq/redismq.go.template",
					},
					ProducerTestsFile: File{
						Name: "producer_tests.go",
						Path: cfg.ProjectPath + "/internal/queue/redismq/",
						Template: "./src/templates/internal/queue/redismq/producer_tests.go.template",
					},
					ConsumerTestsFile: File{
						Name: "consumer_tests.go",
						Path: cfg.ProjectPath + "/internal/queue/redismq/",
						Template: "./src/templates/internal/queue/redismq/consumer_tests.go.template",
					},
				},
			},
			ServiceDir: ServiceDir{
				Name: "service",
				ServiceFile: File{
					Name: "service.go",
					Path: cfg.ProjectPath + "/internal/service/",
					Template: "./src/templates/internal/service/service.go.template",
				},
				ReadmeFile: File{
					Name: "README.md",
					Path: cfg.ProjectPath + "/internal/service/",
					Template: "./src/templates/internal/service/README.md",
				},
			},
			StorageDir: StorageDir{
				Name: "storage",
				ReadmeFile: File{
					Name: "README.md",
					Path: cfg.ProjectPath + "/internal/storage/",
					Template: "./src/templates/internal/storage/README.md",
				},
				StorageFile: File{
					Name: "storage.go",
					Path: cfg.ProjectPath + "/internal/storage/",
					Template: "./src/templates/internal/storage/storage.go.template",
				},
				PsqlDir: PsqlDir{
					Name: "psql",
					PsqlFile: File{
						Name: "psql.go",
						Path: cfg.ProjectPath + "/internal/storage/psql/",
						Template: "./src/templates/internal/storage/psql/psql.go.template",
					},
					ConfigFile: File{
						Name: "config.go",
						Path: cfg.ProjectPath + "/internal/storage/psql/",
						Template: "./src/templates/internal/storage/psql/config.go.template",
					},
				},
				RedisDir: RedisDir{
					Name: "redis",
					RedisFile: File{
						Name: "redis.go",
						Path: cfg.ProjectPath + "/internal/storage/redis/",
						Template: "./src/templates/internal/storage/redis/redis.go.template",
					},
					ConfigFile: File{
						Name: "config.go",
						Path: cfg.ProjectPath + "/internal/storage/redis/",
						Template: "./src/templates/internal/storage/redis/config.go.template",
					},
				},
			},
			TransportDir: TransportDir{
				Name: "transport",
				GrpcDir: GrpcDir{
					Name: "grpc",
					ServerFile: File{
						Name: "server.go",
						Path: cfg.ProjectPath + "/internal/transport/grpc/",
						Template: "./src/templates/internal/transport/grpc/server.go.template",
					},
				},
				HttpDir: HttpDir{
					Name: "http",
					ServerFile: File{
						Name: "server.go",
						Path: cfg.ProjectPath + "/internal/transport/http/",
						Template: "./src/templates/internal/transport/http/server.go.template",
					},
					HandlerDir: HandlerDir{
						Name: "handler",
						HandlerFile: File{
							Name: "handler.go",
							Path: cfg.ProjectPath + "/internal/transport/http/handler/",
							Template: "./src/templates/internal/transport/http/handler/handler.go.template",
						},
						ResponseFile: File{
							Name: "response.go",
							Path: cfg.ProjectPath + "/internal/transport/http/handler/",
							Template: "./src/templates/internal/transport/http/handler/response.go.template",
						},
					},
				},
			},
		},
		PkgDir: PkgDir{
			Name: "pkg",
			ReadmeFile: File{
				Name: "README.md",
				Path: cfg.ProjectPath + "/pkg/",
				Template: "./src/templates/pkg/README.md",
			},
		},
	}
}