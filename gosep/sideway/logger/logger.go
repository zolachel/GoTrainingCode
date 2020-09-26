package main

import "go.uber.org/zap"

func main() {
	l := zap.NewExample()
	l = l.With(zap.Namespace("hometic"), zap.String("I'm", "gopher"))
	l.Info("Hello")
}
