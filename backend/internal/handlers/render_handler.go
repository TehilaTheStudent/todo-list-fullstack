package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RenderHandler struct{}

func NewRenderHandler() *RenderHandler {
	return &RenderHandler{}
}

func (h *RenderHandler) GetServices(c *gin.Context) {
	url := "https://api.render.com/v1/services?includePreviews=true&limit=20"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer rnd_mro2eEG0i3XVNmLcwMHo0izf5jMM")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	c.String(http.StatusOK, string(body))
}
