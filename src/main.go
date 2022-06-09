package main

import (
	"context"
	"fmt"
	"github.com/samithiwat/samithiwat-backend-gateway/src/config"
	"github.com/samithiwat/samithiwat-backend-gateway/src/constant"
	_ "github.com/samithiwat/samithiwat-backend-gateway/src/docs"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/middleware"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/router"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/validator"
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

// @securityDefinitions.apikey  AuthToken
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @tag.name auth
// @tag.description.markdown

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

	v, err := validator.NewValidator()
	if err != nil {
		log.Fatal(err.Error())
	}

	smithConn, err := grpc.Dial(conf.Service.Samithiwat, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to samithiwat service: ", err.Error())
	}

	userClient := proto.NewUserServiceClient(smithConn)
	userSrv := service.NewUserService(userClient)
	userHandler := handler.NewUserHandler(userSrv, v)

	teamClient := proto.NewTeamServiceClient(smithConn)
	teamSrv := service.NewTeamService(teamClient)
	teamHandler := handler.NewTeamHandler(teamSrv, v)

	orgClient := proto.NewOrganizationServiceClient(smithConn)
	orgSrv := service.NewOrganizationService(orgClient)
	orgHandler := handler.NewOrganizationHandler(orgSrv, v)

	authConn, err := grpc.Dial(conf.Service.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to auth service: ", err.Error())
	}

	authClient := proto.NewAuthServiceClient(authConn)
	authSrv := service.NewAuthService(authClient)
	authHandler := handler.NewAuthHandler(authSrv, userSrv, v)

	authGuard := middleware.NewAuthGuard(authSrv, constant.AuthExcludePath)

	r := router.NewFiberRouter(authGuard)

	r.PostAuth("/register", authHandler.Register)
	r.PostAuth("/login", authHandler.Login)
	r.GetAuth("/logout", authHandler.Logout)
	r.PostAuth("/change-password", authHandler.ChangePassword)
	r.GetAuth("/me", authHandler.Validate)
	r.PostAuth("/token", authHandler.RefreshToken)

	r.GetUser("/", userHandler.FindAll)
	r.GetUser("/:id", userHandler.FindOne)
	r.CreateUser("/", userHandler.Create)
	r.PatchUser("/:id", userHandler.Update)
	r.DeleteUser("/:id", userHandler.Delete)

	r.GetTeam("/", teamHandler.FindAll)
	r.GetTeam("/:id", teamHandler.FindOne)
	r.CreateTeam("/", teamHandler.Create)
	r.PatchTeam("/:id", teamHandler.Update)
	r.DeleteTeam("/:id", teamHandler.Delete)

	r.GetOrganization("/", orgHandler.FindAll)
	r.GetOrganization("/:id", orgHandler.FindOne)
	r.CreateOrganization("/", orgHandler.Create)
	r.PatchOrganization("/:id", orgHandler.Update)
	r.DeleteOrganization("/:id", orgHandler.Delete)

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
