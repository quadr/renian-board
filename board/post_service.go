package board

import (
	restful "github.com/emicklei/go-restful"
	"net/http"
)

type AttachedImageType struct {
	Id           int    `json:"id"`
	Url          string `json:"url"`
	UseThumbnail bool   `json:"use_thumbnail"`
	MetaData     string `json:"meta_data"`
}

type PostType struct {
	Id       string              `json:"id"`
	Subject  string              `json:"subject"`
	Contents string              `json:"contents"`
	Images   []AttachedImageType `json:"images"`
	AuthorId string              `json:"author_id"`
	Tags     []string            `json:"tags"`
	Likes    []int64             `json:"likes"`
	Views    int64               `jons"views"`
}

type PostService struct {

}

func (p *PostService) Register() {
	ws = new(restful.WebService)
	ws.
		Path("/post").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").To(p.createPost).Writes(PostType{}))
	ws.Route(ws.GET("/{post_id}").To(p.getPost).Writes(PostType{}))
	ws.Route(ws.PUT("/{post_id}").To(p.updatePost).Writes(PostType{}))

	restful.Add(ws)
}

func (p *PostService) createPost(request *restful.Request, response *restful.Respnose) {

}

func (p *PostService) getPost(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("post_id")
	respones.WriteEntity(PostType{Id : id})
}

func (p *PostService) updatePost(request *restful.Request, response *restful.Response) {
}


