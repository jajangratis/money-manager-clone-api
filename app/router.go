package app

import (
	"fmt"
	"github.com/jajangratis/money-manager-clone-api/controller"
	"github.com/jajangratis/money-manager-clone-api/exception"
	"github.com/julienschmidt/httprouter"
	"github.com/zc2638/swag"
	"io"
	"net/http"
)

type Category struct {
	ID     int64  `json:"category"`
	Name   string `json:"name" enum:"dog,cat" required:""`
	Exists *bool  `json:"exists" required:""`
}

type Pet struct {
	ID        int64     `json:"id"`
	Category  *Category `json:"category" desc:"分类"`
	Name      string    `json:"name" required:"" example:"张三" desc:"名称"`
	PhotoUrls []string  `json:"photoUrls"`
	Tags      []string  `json:"tags" desc:"标签"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, fmt.Sprintf("[%s] Hello World!", r.Method))
}

func NewRouter(mstCategoryController controller.MstCategoryController, mstAccountTypeController controller.MstAccountTypeController, mstTransactionMethod controller.MstTransactionMethodController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/mst-category", mstCategoryController.FindAll)
	router.DELETE("/api/mst-category", mstCategoryController.Delete)
	router.POST("/api/mst-category", mstCategoryController.Create)

	router.GET("/api/mst-account-type", mstAccountTypeController.FindAll)
	router.POST("/api/mst-account-type", mstAccountTypeController.Create)
	router.DELETE("/api/mst-account-type/:id", mstAccountTypeController.Delete)

	router.GET("/api/mst-transaction-method", mstTransactionMethod.FindAll)
	router.DELETE("/api/mst-transaction-method", mstTransactionMethod.Delete)
	router.POST("/api/mst-transaction-method", mstTransactionMethod.Create)

	docs := NewDocs(router)
	router.Handler(http.MethodGet, "/swagger/json", docs.Handler())
	router.Handler(http.MethodGet, "/swagger/ui/*any", swag.UIHandler("/swagger/ui", "/swagger/json", true))

	router.PanicHandler = exception.ErrorHandler
	return router
}
