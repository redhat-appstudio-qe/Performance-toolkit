package utils

import (
	"context"
	"fmt"
	"os"
)

func CheckVarExistsAndReturn(variable string) string {
	v, ok := os.LookupEnv(variable)
  if !ok {
    fmt.Printf("The Variable %s is not present\n", variable)
	os.Exit(1)
  } else {
    fmt.Printf("The Variable %s is present\n", variable)
	return v 
  }
  return v
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
