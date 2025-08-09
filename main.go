package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Servir arquivos estáticos
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Endpoint para gerar QR Code
	http.HandleFunc("/qrcode", func(w http.ResponseWriter, r *http.Request) {
		// Log da requisição recebida
		log.Printf("Recebida requisição para gerar QR Code do IP: %s", r.RemoteAddr)

		userURL := r.URL.Query().Get("url")
		if userURL == "" {
			// Log do erro
			log.Println("Erro: Parâmetro 'url' não foi fornecido.")
			http.Error(w, "Parâmetro 'url' é obrigatório", http.StatusBadRequest)
			return
		}
		log.Printf("URL recebida: %s", userURL)

		// Adiciona https:// se faltar
		if !strings.HasPrefix(userURL, "http://") && !strings.HasPrefix(userURL, "https://") {
			userURL = "https://" + userURL
		}

		// Valida URL
		parsed, err := url.ParseRequestURI(userURL)
		if err != nil {
			// Log do erro de validação
			log.Printf("Erro: URL inválida fornecida '%s'. Detalhe: %v", userURL, err)
			http.Error(w, "URL inválida", http.StatusBadRequest)
			return
		}

		// Gera QR Code
		qr, err := qrcode.Encode(parsed.String(), qrcode.Medium, 256)
		if err != nil {
			// Log do erro na geração
			log.Printf("Erro ao gerar QR Code para a URL '%s'. Detalhe: %v", parsed.String(), err)
			http.Error(w, "Erro ao gerar QR Code", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(qr)

		// Log de sucesso
		log.Printf("QR Code gerado com sucesso para: %s", parsed.String())
	})

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}