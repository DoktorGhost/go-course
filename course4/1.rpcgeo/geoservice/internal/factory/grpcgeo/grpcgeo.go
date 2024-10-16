package grpcgeo

import (
	"context"
	"geoservice/internal/entities"
	pb "geoservice/internal/factory/grpcgeo/generated"
	"geoservice/internal/metrics"
	"google.golang.org/grpc"
	"log"
	"time"
)

type GeoServiceGrpc struct {
	Client pb.GeoServiceClient
}

func NewGeoServiceGrpc() *GeoServiceGrpc {
	conn, err := grpc.NewClient("grpcgeo:1236", grpc.WithInsecure())
	if err != nil {
		log.Println("Error connecting to gRPC server:", err)
		return nil
	}
	client := pb.NewGeoServiceClient(conn)

	return &GeoServiceGrpc{client}
}

func (g *GeoServiceGrpc) AddressSearch(input entities.SearchRequest) (entities.Response, error) {
	start := time.Now()

	var addressSearchResp entities.Response

	req := &pb.SearchRequest{Query: input.Query}
	resp, err := g.Client.AddressSearch(context.Background(), req)
	if err != nil {
		log.Println("Error calling AddressSearch:", err)
		return addressSearchResp, err
	}

	for _, addr := range resp.Addresses {
		addressSearchResp.Addresses = append(addressSearchResp.Addresses, entities.Address{
			City:   addr.City,
			Street: addr.Street,
			House:  addr.House,
			Lat:    addr.Lat,
			Lon:    addr.Lon,
		})
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("AddressSearch").Observe(duration)

	return addressSearchResp, nil
}

func (g *GeoServiceGrpc) GeoCode(input entities.GeocodeRequest) (entities.Response, error) {
	start := time.Now()

	var geoCodeResp entities.Response

	req := &pb.GeoCodeRequest{Lat: input.Lat, Lng: input.Lng}
	resp, err := g.Client.GeoCode(context.Background(), req)
	if err != nil {
		log.Println("Error calling GeoCode:", err)
		return geoCodeResp, err
	}

	for _, addr := range resp.Addresses {
		geoCodeResp.Addresses = append(geoCodeResp.Addresses, entities.Address{
			City:   addr.City,
			Street: addr.Street,
			House:  addr.House,
			Lat:    addr.Lat,
			Lon:    addr.Lon,
		})
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("GeoCode").Observe(duration)

	return geoCodeResp, nil
}
