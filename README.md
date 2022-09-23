## Template-microservice-go
A tool for quickly generating a template for a new microservice.

## Usage

1. Clone project:
```bash
cd ~
git clone https://github.com/woyow/template-microservice-go.git
```

2. Go into project directory:
```bash
cd ~/template-microservice-go
```

3. To create a template for a new microservice, enter the following command:
```bash
go run ./cmd/main.go \ 
    --path ./my-dir-for-project \ 
    --module-name github.com/example-nickname/example-service \ 
    --name example-service-name
```

or run binary file:
```bash
./micro-svc-builder \ 
    --path ./my-dir-for-project \ 
    --module-name github.com/example-nickname/example-service \ 
    --name example-service-name
```