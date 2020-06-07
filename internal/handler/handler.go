package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/beruangcoklat/share-variables/internal/entity"
	"github.com/beruangcoklat/share-variables/internal/usecase"
)

type Handler struct {
}

var (
	once    sync.Once
	handler *Handler
)

func GetHandler() *Handler {
	once.Do(func() {
		handler = &Handler{}
	})
	return handler
}

func (*Handler) StoreVariable(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	data := &entity.Variable{}
	if err := json.Unmarshal(body, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if err := usecase.GetUseCase().UpdateVariable(context.Background(), data.Key, data.Value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
