package usecase

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/beruangcoklat/share-variables/internal/repository"
	sharevariablespb "github.com/beruangcoklat/share-variables/proto"
	"google.golang.org/grpc"
)

type Usecase struct {
	client sharevariablespb.ShareServiceClient
	stream sharevariablespb.ShareService_UpdateVariableClient
}

var (
	once    sync.Once
	usecase *Usecase
)

func GetUseCase() *Usecase {
	once.Do(func() {
		clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		client := sharevariablespb.NewShareServiceClient(clientConn)
		stream, err := client.UpdateVariable(context.Background())

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		usecase = &Usecase{
			client: client,
			stream: stream,
		}
	})
	return usecase
}

func (uc *Usecase) UpdateVariable(ctx context.Context, key string, value string) error {
	if err := uc.stream.Send(&sharevariablespb.ShareRequest{
		Key:   key,
		Value: value,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (uc *Usecase) PrintVariable(ctx context.Context) {
	r := repository.GetRepository()
	key := "var1"
	for {
		fmt.Printf("%v : %v\n", key, r.GetVariable(key))
		time.Sleep(time.Second * 1)
	}
}

func (uc *Usecase) ReceiveVariable(ctx context.Context) error {
	for {
		res, err := uc.stream.Recv()
		if err != nil {
			fmt.Println(err)
			return err
		}

		repository.GetRepository().UpdateVariable(res.GetKey(), res.GetValue())
	}
}
