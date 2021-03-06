package image

import (
	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/server/router"
)

// imageRouter is a router to talk with the image controller
type imageRouter struct {
	backend Backend
	decoder httputils.ContainerDecoder
	routes  []router.Route
}

// NewRouter initializes a new image router
func NewRouter(backend Backend, decoder httputils.ContainerDecoder) router.Router {
	r := &imageRouter{
		backend: backend,
		decoder: decoder,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the image controller
func (r *imageRouter) Routes() []router.Route {
	return r.routes
}

// initRoutes initializes the routes in the image router
func (r *imageRouter) initRoutes() { //Docker镜像存储相关数据结构 见http://www.infocool.net/kb/OtherCloud/201705/346489.html
	r.routes = []router.Route{
		// GET
		router.NewGetRoute("/images/json", r.getImagesJSON),
		router.NewGetRoute("/images/search", r.getImagesSearch),
		router.NewGetRoute("/images/get", r.getImagesGet),
		router.NewGetRoute("/images/{name:.*}/get", r.getImagesGet),
		router.NewGetRoute("/images/{name:.*}/history", r.getImagesHistory),
		router.NewGetRoute("/images/{name:.*}/json", r.getImagesByName),
		// POST
		router.NewPostRoute("/commit", r.postCommit),
		router.NewPostRoute("/images/load", r.postImagesLoad),
		router.Cancellable(router.NewPostRoute("/images/create", r.postImagesCreate)),  //例如客户端docker pull对应的就是这里，daemon会发送create到数据仓库
		router.Cancellable(router.NewPostRoute("/images/{name:.*}/push", r.postImagesPush)),
		router.NewPostRoute("/images/{name:.*}/tag", r.postImagesTag),
		router.NewPostRoute("/images/prune", r.postImagesPrune),
		// DELETE
		router.NewDeleteRoute("/images/{name:.*}", r.deleteImages),
	}
}
