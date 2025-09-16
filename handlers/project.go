package handlers

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/codetesla51/portBackend/models"
)

func GetProjects(c *gin.Context) {
    page := 1
    if p := c.Query("page"); p != "" {
        fmt.Sscanf(p, "%d", &page)
    }
    limit := 8
    offset := (page - 1) * limit

    projects, err := models.GetVisibleProjects(limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": projects,
        "page": page,
        "per_page": limit,
    })
}
func GetProject (c *gin.Context){
  slug := c.Param("slug")
  project , err := models.GetProjectBySlug(slug)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
return
  }
  c.JSON(http.StatusOK, project )
}
