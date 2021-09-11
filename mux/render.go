package mux

import (
	"encoding/xml"
	"net/http"

	"github.com/clbanning/mxj"
	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/proxy"
)

// Render marshals the proxy response and passes the resulting xml to the response writer
func Render(w http.ResponseWriter, response *proxy.Response) {
	if response == nil {
		xml.NewEncoder(w).Encode(nil)
		w.Header().Add("Content-Type", gin.MIMEXML)
		return
	}

	mv := mxj.Map(response.Data)
	w.Header().Add("Content-Type", gin.MIMEXML)
	mv.XmlWriter(w)
}
