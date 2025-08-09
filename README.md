# Gerador de QR Code

![Pipeline](https://github.com/dellabeneta/qrcode-gen/actions/workflows/main.yaml/badge.svg)

Um serviço web simples para geração de QR Codes a partir de URLs, desenvolvido em Go. Digite um endereço na interface web e obtenha instantaneamente o QR Code correspondente.

## Funcionalidades

- Geração de QR Code para qualquer URL
- Interface web responsiva e fácil de usar
- Backend em Go com endpoint `/qrcode`
- Pronto para rodar em Docker e Kubernetes
- CI/CD com GitHub Actions

## Começando

### Pré-requisitos

- Go 1.24.x ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

### Instalação Local

```bash
git clone https://github.com/dellabeneta/qrcode-gen.git
cd qrcode-gen
go run main.go
```

Acesse em [http://localhost:8080](http://localhost:8080)

### Usando Docker

```bash
docker build -t qrcode-gen .
docker run -p 8080:8080 qrcode-gen
```

### Deploy no Kubernetes

```bash
kubectl apply -f k3s/namespace.yaml
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

A aplicação estará disponível na porta `30084` do cluster.

## Estrutura do Projeto

```
.
├── Dockerfile
├── go.mod
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── main.go
├── nuke.sh
├── README.md
└── static
    ├── favicon.ico
    ├── index.html
    ├── script.js
    └── style.css
```

## Como Funciona

1. O usuário digita o endereço (URL) na interface web.
2. O frontend envia a URL para o backend via `/qrcode`.
3. O backend valida e gera o QR Code (PNG).
4. O QR Code é exibido na tela.

## Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes).

```bash
chmod +x nuke.sh
./nuke.sh
```

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.