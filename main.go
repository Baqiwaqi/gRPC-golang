package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"time"

	pb "example.com/baqiwaqi/golang-grpc-servive/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50052"
)

var (
	ErrFlightNotFound = status.Error(codes.NotFound, "Flight not found")
)

// array of mock flights
var flights = []*pb.MyService_Flight{
	{
		Id:     1,
		Number: "AA123",
		Airline: &pb.MyService_Airline{
			Name: "American Airlines",
			Code: "AA",
		},
		Departure: &pb.MyService_Airport{
			Code:    "SFO",
			Name:    "San Francisco",
			City:    "San Francisco",
			Country: "USA",
		},
		Arrival: &pb.MyService_Airport{
			Code:    "LAX",
			Name:    "Los Angeles",
			City:    "Los Angeles",
			Country: "USA",
		},
	},
	{
		Id:     2,
		Number: "UA123",
		Airline: &pb.MyService_Airline{
			Name: "United Airlines",
			Code: "UA",
		},
		Departure: &pb.MyService_Airport{
			Code:    "AMS",
			Name:    "Schiphol",
			City:    "Amsterdam",
			Country: "Netherlands",
		},
		Arrival: &pb.MyService_Airport{
			Code:    "SFO",
			Name:    "San Francisco",
			City:    "San Francisco",
			Country: "USA",
		},
	},
}

type TrackingServer struct {
	pb.UnimplementedFlightTrackingServer
	flightCh chan *pb.MyService_Flight
}

func (t *TrackingServer) generateFakeFlights() {
	for {
		select {
		case fake := <-time.After(1 * time.Second):
			t.flightCh <- &pb.MyService_Flight{
				Id:     1,
				Number: strconv.Itoa(int(fake.Unix())),
			}
		}
	}
}

// function GetFlight
func (s *TrackingServer) GetFlight(ctx context.Context, req *pb.MyService_Flight_Request) (*pb.MyService_Flight_Response, error) {
	log.Printf("GetFlight: %v", req)
	for _, f := range flights {
		if f.Id == req.Id {
			return &pb.MyService_Flight_Response{Flight: f}, nil
		}
	}
	return nil, ErrFlightNotFound
}

func (s *TrackingServer) StreamFlights(req *pb.MyService_Flight_Request, stream pb.FlightTracking_StreamFlightsServer) error {
	log.Printf("Stream flights: %v", req)
	for {
		select {
		case flight := <-s.flightCh:
			//time.Sleep(1 * time.Second)
			if err := stream.Send(&pb.MyService_Flight_Response{Flight: flight}); err != nil {
				log.Println(err)
			}
		}
		if s.flightCh == nil {
			break
		}
	}
	return nil
}

func (s *TrackingServer) ListFlights(ctx context.Context, req *pb.MyService_Flight_Request) (*pb.MyService_Flight_ListResponse, error) {
	log.Printf("List flights: %v", req)
	return &pb.MyService_Flight_ListResponse{Flights: flights}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var trackingServer = &TrackingServer{}

	trackingServer.flightCh = make(chan *pb.MyService_Flight)
	defer close(trackingServer.flightCh)
	go trackingServer.generateFakeFlights()

	s := grpc.NewServer()
	pb.RegisterFlightTrackingServer(s, trackingServer)
	log.Printf("Server started on port %s", lis.Addr())
	s.Serve(lis)
}
