package xendpoint

import (
	"boostPuzzle/server/models"
	"boostPuzzle/server/xservice"
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type EndpointFactory interface {
	GetAlbumHTML() func(c *gin.Context)
	GetAlbumRest() func(c *gin.Context)
}

type endpointFactory struct {
	service xservice.Service
}

func NewEndpointFactory(service xservice.Service) EndpointFactory {
	return endpointFactory{
		service: service,
	}
}

type GetAlbumRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (factory endpointFactory) GetAlbumRest() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req GetAlbumRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		offset, _ := c.GetQuery("offset")
		mediaAlbum, err := factory.service.GetAlbumByUser(req.Username, offset)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"res": mediaAlbum})
	}
}

type LoadMore struct {
	Username string
	Offset   string
}

type MediaTmpl struct {
	*models.MediaAlbum
	LoadMore
}

func (factory endpointFactory) GetAlbumHTML() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req GetAlbumRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		mediaAlbum, err := factory.service.GetAlbumByUser(req.Username, "")
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		// Parse the three templates into a single *template.Template
		t, err := template.ParseFiles("server/static/header.html", "server/static/index.html", "server/static/footer.html", "server/static/loadMore.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing templates: %s", err)
			return
		}

		// Render the templates and get the resulting HTML output
		var buf bytes.Buffer
		tmpl := MediaTmpl{
			MediaAlbum: mediaAlbum,
			LoadMore: LoadMore{
				Username: req.Username,
				Offset:   mediaAlbum.Extra.Offset,
			},
		}
		err = t.ExecuteTemplate(&buf, "index", tmpl)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error rendering templates: %s", err)
			return
		}
		c.Data(200, "text/html", buf.Bytes())
	}
}
