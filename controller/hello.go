package controller

import (
	"github.com/burakkoken/procyon-helloworld/model"
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type RequestHello struct {
	PathVariables struct{
		Name string `json:"name" yaml:"name"`
	} `request:"path"`
	Header struct{
		ContentType string `json:"Content-Type" yaml:"Content-Type"`
	} `request:"header"`
	Param struct{
		Query string `json:"query" yaml:"query"`
	} `request:"param"`
}

type HelloController struct {
	logger 		context.Logger
	testConfig 	model.TestConfiguration
}

func NewHelloController(logger context.Logger, testConfig model.TestConfiguration) HelloController  {
	return HelloController{
		logger: logger,
		testConfig: testConfig,
	}
}

func (controller HelloController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.Register(
		web.Get(controller.HelloWorld, web.Path("/hello/:name/test"), web.RequestObject(RequestHello{})),
		)
}

func (controller HelloController) HelloWorld(context *web.WebRequestContext)  {
	request := &RequestHello{}
	err := context.BindRequest(request)
	if err != nil {
		context.SetHTTPError(web.HttpErrorBadRequest)
	} else {
		context.SetModel("Hello " + request.PathVariables.Name)
	}
}