package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"

	pb "github.com/Its-Alex/flatsharing/internal/api-flatsharing/v1"
	"github.com/Its-Alex/flatsharing/internal/core/database"
	"github.com/Its-Alex/flatsharing/internal/core/helper"
	"github.com/gobuffalo/packr"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Service current grpc service
type service struct {
	db *database.DB
}

type runner func() error

// Global env variables
var (
	db *sqlx.DB
	// runners contain function that will be launch when program start
	runners = []runner{
		RunGRPC,
		RunJSON,
	}
)

func init() {
	viper.SetEnvPrefix("api-flatsharing")

	viper.BindEnv("grpc_listen_addr")
	viper.BindEnv("grpc_listen_port")
	viper.BindEnv("json_listen_addr")
	viper.BindEnv("json_listen_port")

	viper.BindEnv("psql_host")
	viper.BindEnv("psql_port")
	viper.BindEnv("psql_user")
	viper.BindEnv("psql_password")
	viper.BindEnv("psql_dbname")
	viper.BindEnv("psql_sslmode")
	viper.BindEnv("psql_max_retry")

	viper.SetDefault("grpc_listen_addr", "0.0.0.0")
	viper.SetDefault("grpc_listen_port", "8080")
	viper.SetDefault("json_listen_addr", "0.0.0.0")
	viper.SetDefault("json_listen_port", "8081")

	viper.SetDefault("psql_sslmode", "disable")
	viper.SetDefault("psql_max_retry", 10)

	helper.InitLogger()
}

// RunGRPC services
func RunGRPC() error {
	// Open socket connection for GRPC
	listen, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%s",
		viper.GetString("grpc_listen_addr"),
		viper.GetString("grpc_listen_port"),
	))
	if err != nil {
		return err
	}

	// Create GRPC server
	server := grpc.NewServer()

	// Register services
	s := &service{}
	s.db, err = database.Connect()
	if err != nil {
		return err
	}

	pb.RegisterFlatServicesServer(server, s)
	pb.RegisterPurchaseServicesServer(server, s)

	// Enable reflection
	reflection.Register(server)

	// Start gRPC
	helper.Logger.Infof("gRPC service start at %s:%s",
		viper.GetString("grpc_listen_addr"),
		viper.GetString("grpc_listen_port"))
	return server.Serve(listen)
}

// RunJSON run JSON api
func RunJSON() error {
	// Create context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	SwaggerUIBox := packr.NewBox("../../../assets/swagger-ui")
	mux.Handle(
		"/",
		http.FileServer(SwaggerUIBox),
	)
	SwaggerFileBox := packr.NewBox("../swagger")
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		file, err := SwaggerFileBox.FindString("flatsharing.swagger.json")
		if err != nil {
			helper.Logger.Error(err)
		}
		io.Copy(w, strings.NewReader(file))
	})

	// Create mux
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// Setup JSON service
	err := pb.RegisterFlatServicesHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf(
		"%s:%s",
		viper.GetString("grpc_listen_addr"),
		viper.GetString("grpc_listen_port"),
	), opts)
	if err != nil {
		return err
	}
	err = pb.RegisterPurchaseServicesHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf(
		"%s:%s",
		viper.GetString("grpc_listen_addr"),
		viper.GetString("grpc_listen_port"),
	), opts)
	if err != nil {
		return err
	}
	mux.Handle("/v1/", gwmux)

	// Start JSON service
	helper.Logger.Infof("JSON service start at %s:%s",
		viper.GetString("json_listen_addr"),
		viper.GetString("json_listen_port"))
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(
		"%s:%s",
		viper.GetString("json_listen_addr"),
		viper.GetString("json_listen_port"),
	))
	return nil
}

func main() {
	// Start services
	var wg sync.WaitGroup
	for key, f := range runners {
		wg.Add(1)
		go func(key int, f runner) {
			defer wg.Done()
			if err := f(); err != nil {
				helper.Logger.Errorf("Something went wrong with runner %d : %s", key, err)
			}
		}(key, f)
	}
	wg.Wait()

	// Sync logger before exit
	err := helper.Logger.Sync()
	if err != nil {
		fmt.Printf("Sync logger failed: %s", err)
	}
}
