package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func CheckVarExistsAndReturn(variable string, failonError bool) string {
	v, ok := os.LookupEnv(variable)
  if !ok {
    log.Printf("The Variable %s is not present\n", variable)
	if failonError{
		log.Fatalf("Cannot Proceed without Variable %s \n", variable)
	}
	return ""
  } else {
    log.Printf("The Variable %s is present\n", variable)
	return v 
  }
}

func ToDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

func GetWorkingDir()string{
	mydir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
	return mydir
}
func SetKubeConfig(variable string) {
	err := os.Setenv("KUBECONFIG", variable)
	if err != nil{
		fmt.Printf("Error changing Kubeconfig")
		os.Exit(1)
	}
}

func SetChaosNamespace(s string, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "chaos_ns", s)
	return ctx
}
func SetTestNamespace(s string, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "test-namespace", s)
	return ctx
}
func SetpodDeleteLabelKey(s string, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "podDeleteLabelKey", s)
	return ctx
}
func SetpodDeleteLabelValue(s string, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "podDeleteLabelValue", s)
	return ctx
}

func RandomString(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 4)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	suffix := strconv.Itoa(rand.Intn(int(9999 - 1000 + 1)) + 1000)
	return prefix + string(b) + suffix
}
