package handler

import (
	"github.com/asim/go-micro/v3/util/log"
	"golang.org/x/net/context"
	helloworld "micro/helloworld/proto"
)

type Helloworld struct{}

func (this *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Info("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (this *Helloworld) Stream(ctx context.Context, req *helloworld.StreamingRequest, stream helloworld.Helloworld_StreamStream) error {
	log.Infof("Received Helloworld.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&helloworld.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (this *Helloworld) PingPong(ctx context.Context, stream helloworld.Helloworld_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&helloworld.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
