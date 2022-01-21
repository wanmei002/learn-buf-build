package main

import (
    "context"
    "github.com/sirupsen/logrus"
    "google.golang.org/genproto/googleapis/type/datetime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/backoff"
    "google.golang.org/grpc/keepalive"
    "google.golang.org/grpc/metadata"
    "gt/pb"
    "time"
)

func main(){
    // 客户端的option
    logrus.SetLevel(logrus.DebugLevel)
    dl, err := grpc.Dial(":9989", opt()...)
    if err != nil {
        logrus.Debug("dial failed; err:", err)
        return
    }

    cli := pb.NewFarmServiceClient(dl)

    md := metadata.Pairs("metadata", "zzh-zyn")
    ctx := metadata.NewOutgoingContext(context.Background(), md)

    // 下面通过tags传递元数据是不行了 只能通过 metadata 来传播元数据
    //tgs := grpc_ctxtags.NewTags()
    //tgs.Set("request_id", "zzh-good")
    //ctx = grpc_ctxtags.SetInContext(ctx, tgs)
    //
    //tgg := grpc_ctxtags.Extract(ctx)
    //logrus.Printf("ggg:%+v\n", tgg)
    ret, err := cli.AskPrice(ctx, &pb.Pig{
        Category: "1",
        Weight:   "2",
        Birthday: &datetime.DateTime{
            Year:       2022,
            Month:      01,
            Day:        21,
            Hours:      0,
            Minutes:    0,
            Seconds:    0,
            Nanos:      0,
        },
    })

    if err != nil {
        logrus.Debug("ask price failed; err:", err)
        return
    }

    logrus.Debugf("val:%+v\n", ret)
}

func opt() []grpc.DialOption {
    var list []grpc.DialOption
    list = append(list, grpc.WithInsecure())
    list = append(list, grpc.WithDefaultCallOptions(grpc.WaitForReady(false)))
    list = append(list, grpc.WithKeepaliveParams(keepalive.ClientParameters{
        Time:                10*time.Second,
        Timeout:             30*time.Second,
        PermitWithoutStream: true,
    }))

    list = append(list, grpc.WithConnectParams(grpc.ConnectParams{
        Backoff:           backoff.Config{
            BaseDelay:  2 * time.Second,
            Multiplier: 1,
            Jitter:     0.2,
            MaxDelay:   3 * time.Second,
        },
        MinConnectTimeout: 10 * time.Second,
    }))

    return list
}
