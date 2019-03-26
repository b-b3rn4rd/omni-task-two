package home

import "github.com/emicklei/go-restful"

var (
	// Version application version, will be replaced during build
	Version = "1.0.0"
	// Lastcommitsha current commit sha1, will be replaced during build
	Lastcommitsha = "unknown"
)

type healthcheck struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	Lastcommitsha string `json:"lastcommitsha"`
}

// New creates new home resource
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("").To(Home))
	service.Route(service.GET("healthcheck").To(Healthcheck))

	return service
}

// Home home action
func Home(request *restful.Request, response *restful.Response) {
	response.WriteEntity(struct {
		Message string `json:"message"`
	}{Message: "Hello World"})
}

// Healthcheck healtcheck action
func Healthcheck(request *restful.Request, response *restful.Response) {
	response.WriteEntity(healthcheck{
		Description:   "pre-interview technical test",
		Version:       Version,
		Lastcommitsha: Lastcommitsha,
	})
}
