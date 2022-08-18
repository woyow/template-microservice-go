package generate

import (
	"fmt"

	"github.com/woyow/template-microservice-go/config"
	"github.com/woyow/template-microservice-go/internal/fs"
	"github.com/woyow/template-microservice-go/internal/structure"
)

const (
	sep = "/"
)

func DirsGenerate(s *structure.Dirs, cfg *config.Config) error {
	var dirs []string

	// Cmd directory
	cmdDir := cfg.ProjectPath + sep + s.CmdDir.Name + sep + s.CmdDir.CmdServiceDir.Name
	dirs = append(dirs, cmdDir)

	// Config directory
	configDir := cfg.ProjectPath + sep + s.ConfigDir.Name
	dirs = append(dirs, configDir)

	// DB directory
	dbDir := cfg.ProjectPath + sep + s.DBDir.Name + sep + s.DBDir.MigrationsDir.Name
	dirs = append(dirs, dbDir)

	// Internal directory
	appDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.AppDir.Name
	entityDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.EntityDir.Name
	serviceDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.ServiceDir.Name
	psqlDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.StorageDir.Name + sep + s.InternalDir.StorageDir.PsqlDir.Name
	redisDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.StorageDir.Name + sep + s.InternalDir.StorageDir.RedisDir.Name
	handlerDir := cfg.ProjectPath + sep + s.InternalDir.Name + sep + s.InternalDir.TransportDir.Name + sep + s.InternalDir.TransportDir.HttpDir.Name + sep + s.InternalDir.TransportDir.HttpDir.HandlerDir.Name

	dirs = append(dirs, appDir)
	dirs = append(dirs, entityDir)
	dirs = append(dirs, serviceDir)
	dirs = append(dirs, psqlDir)
	dirs = append(dirs, redisDir)
	dirs = append(dirs, handlerDir)

	// Pkg directory
	pkgDir := cfg.ProjectPath + sep + s.PkgDir.Name
	dirs = append(dirs, pkgDir)

	for _, dir := range dirs {
		// Create dir if not exist
		err := fs.CreateDir(dir); if err != nil {
			return err
		}
	}

	return nil
}

func FilesGenerate(s *structure.Dirs, words *Words) error {

	// .env.example
	if err := FileGeneration(words,	s.EnvExampleFile); err != nil {
		return err
	}

	// .gitignore
	if err := FileGeneration(words,	s.GitIgnoreFile); err != nil {
		return err
	}

	// go.mod
	if err := FileGeneration(words,	s.GoModFile); err != nil {
		return err
	}

	// Dockerfile
	if err := FileGeneration(words,	s.DockerFile); err != nil {
		return err
	}

	// Cmd directory
	{
		// Service directory
		{
			// main.go
			if err := FileGeneration(words,	s.CmdDir.CmdServiceDir.MainFile); err != nil {
				return err
			}
		}
	}

	// Config directory
	{
		// config.go
		if err := FileGeneration(words, s.ConfigDir.ConfigFile); err != nil {
			return err
		}

		// local.yaml
		if err := FileGeneration(words, s.ConfigDir.LocalConfigYamlFile); err != nil {
			return err
		}
	}

	// DB directory
	{
		// README.md
		if err := FileGeneration(words, s.DBDir.ReadmeFile); err != nil {
			return err
		}
	}

	// Internal directory
	{
		// README.md
		if err := FileGeneration(words, s.InternalDir.ReadmeFile); err != nil {
			return err
		}

		// App directory
		{
			// app.go
			if err := FileGeneration(words, s.InternalDir.AppDir.AppFile); err != nil {
				return err
			}

			// migrate.go
			if err := FileGeneration(words, s.InternalDir.AppDir.MigrateFile); err != nil {
				return err
			}
		}

		// Entity directory
		{
			// README.md
			if err := FileGeneration(words, s.InternalDir.EntityDir.ReadmeFile); err != nil {
				return err
			}
		}

		// Service directory
		{
			// service.go
			if err := FileGeneration(words, s.InternalDir.ServiceDir.ServiceFile); err != nil {
				return err
			}

			// README.md
			if err := FileGeneration(words, s.InternalDir.ServiceDir.ReadmeFile); err != nil {
				return err
			}
		}

		// Storage directory
		{
			// storage.go
			if err := FileGeneration(words, s.InternalDir.StorageDir.StorageFile); err != nil {
				return err
			}

			// README.md
			if err := FileGeneration(words, s.InternalDir.StorageDir.ReadmeFile); err != nil {
				return err
			}

			// Psql directory
			{
				// config.go
				if err := FileGeneration(words, s.InternalDir.StorageDir.PsqlDir.ConfigFile); err != nil {
					return err
				}

				// psql.go
				if err := FileGeneration(words, s.InternalDir.StorageDir.PsqlDir.PsqlFile); err != nil {
					return err
				}
			}

			// Redis directory
			{
				// config.go
				if err := FileGeneration(words, s.InternalDir.StorageDir.RedisDir.ConfigFile); err != nil {
					return err
				}

				// redis.go
				if err := FileGeneration(words, s.InternalDir.StorageDir.RedisDir.RedisFile); err != nil {
					return err
				}
			}
		}

		// Transport directory
		{

			// Http directory
			{
				// server.go
				if err := FileGeneration(words, s.InternalDir.TransportDir.HttpDir.ServerFile); err != nil {
					return err
				}

				// Handler directory
				{
					// handler.go
					if err := FileGeneration(words, s.InternalDir.TransportDir.HttpDir.HandlerDir.HandlerFile); err != nil {
						return err
					}

					// response.go
					if err := FileGeneration(words, s.InternalDir.TransportDir.HttpDir.HandlerDir.ResponseFile); err != nil {
						return err
					}
				}
			}
		}
	}

	// Pkg directory
	{
		// README.md
		if err := FileGeneration(words, s.PkgDir.ReadmeFile); err != nil {
			return err
		}
	}

	return nil
}

func FileGeneration(words *Words, file structure.File) error {
	f, err := fs.ReadFile(file.Template)
	if err != nil {
		return err
	}

	f = ReplaceWords(words, &f)

	fPath := file.Path + file.Name

	if err = fs.WriteFile(fPath, f); err != nil {
		return err
	}

	return nil
}

func CompleteGeneration(cfg *config.Config) {
	fmt.Printf("Project create on %s directory\n", cfg.ProjectPath)
}

func CodeGenerate(s *structure.Dirs, cfg *config.Config) {
	DirsGenerate(s, cfg)

	words := NewWords(cfg)

	FilesGenerate(s, words)

	CompleteGeneration(cfg)
}