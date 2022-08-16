package main

import (
	"github.com/woyow/template-microservice-go/config"
	"github.com/woyow/template-microservice-go/internal/structure"
	_ "github.com/woyow/template-microservice-go/internal/fs"
	"github.com/woyow/template-microservice-go/internal/generate"
)

func main() {
	cfg := config.NewConfig()

	s := structure.NewStructure(cfg)

	generate.CodeGenerate(s, cfg)
}
