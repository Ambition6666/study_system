package clients

import (
	"fmt"
	pro "studysystem/api/proto/problem"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

var ProCli pro.PrivateServiceClient

func InitProGRPC() (*grpc.ClientConn, error) {
	// etcd
	etcdCli, err := clientv3.NewFromURL("192.168.1.67:2379")
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdCli)
	if err != nil {
		return nil, err
	}
	// dial
	conn, err := grpc.Dial("etcd:///"+"private",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		return conn, err
	}
	// new client
	ProCli = pro.NewPrivateServiceClient(conn)
	return conn, nil
}
