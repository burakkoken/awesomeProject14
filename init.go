package main

import (
	"github.com/burakkoken/procyon-helloworld/controller"
	"github.com/burakkoken/procyon-helloworld/model"
	core "github.com/procyon-projects/procyon-core"
)

func init()  {
	core.Register(controller.NewHelloController)
	core.Register(model.NewTest)
}
