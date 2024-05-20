package main

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/aronrichter/monorepo-with-earthly/libs/log"
)

func main() {
	localLogger := log.New()
	defer func(localLogger *zap.Logger) {
		_ = localLogger.Sync()
	}(localLogger)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		localLogger.Info("iniciei aqui", zap.String("path", request.URL.Path))
		_, err := writer.Write([]byte("Hello, Earth!"))
		if err != nil {
			localLogger.Error("deu ruim aqui", zap.String("path", request.URL.Path), zap.Int("status", http.StatusInternalServerError))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		localLogger.Info("terminei aqui", zap.String("path", request.URL.Path), zap.Int("status", http.StatusOK))

		writer.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		localLogger.Info("erro em providenciar a API", zap.Error(err))
	}
}
