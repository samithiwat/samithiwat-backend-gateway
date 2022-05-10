package main

import (
	"context"
	"fmt"
	"github.com/samithiwat/samithiwat-backend-gateway/src/config"
	_ "github.com/samithiwat/samithiwat-backend-gateway/src/docs"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/router"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// @title Samithiwat Backend
// @version 1.0
// @description.markdown

// @contact.name Samithiwat
// @contact.email admin@samithiwat.dev
// @contact.url https://samithiwat.dev

// @schemes https http

// @securityDefinitions.apikey  Auth Token
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @tag.name user
// @tag.description.markdown

// @tag.name organization
// @tag.description.markdown

// @tag.name team
// @tag.description.markdown

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}

	userConn, err := grpc.Dial(conf.Service.User, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to user service: ", err.Error())
	}

	userClient := proto.NewUserServiceClient(userConn)
	userSrv := service.NewUserService(userClient)
	userHandler := handler.NewUserHandler(userSrv)

	orgConn, err := grpc.Dial(conf.Service.Organization, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to team service ", err.Error())
	}

	teamClient := proto.NewTeamServiceClient(orgConn)
	teamSrv := service.NewTeamService(teamClient)

	orgClient := proto.NewOrganizationServiceClient(orgConn)
	orgSrv := service.NewOrganizationService(orgClient)

	r := router.NewFiberRouter()

	r.GetUser("/user", userHandler.FindAll)
	r.GetUser("/user/:id", userHandler.FindOne)
	r.CreateUser("user", userHandler.Create)
	r.PatchUser("/user/:id", userHandler.Update)
	r.DeleteUser("user/:id", userHandler.Delete)

	r.GetTeam("/team", teamSrv.FindAll)
	r.GetTeam("/team/:id", teamSrv.FindOne)
	r.CreateTeam("team", teamSrv.Create)
	r.PatchTeam("/team/:id", teamSrv.Update)
	r.DeleteTeam("team/:id", teamSrv.Delete)

	r.GetOrganization("/organization", orgSrv.FindAll)
	r.GetOrganization("/organization/:id", orgSrv.FindOne)
	r.CreateOrganization("organization", orgSrv.Create)
	r.PatchOrganization("/organization/:id", orgSrv.Update)
	r.DeleteOrganization("organization/:id", orgSrv.Delete)

	go func() {
		if err := r.Listen(fmt.Sprintf(":%v", conf.App.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"server": func(ctx context.Context) error {
			return r.Shutdown()
		},
	})

	<-wait
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Printf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Printf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}
