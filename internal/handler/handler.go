package handler

import (
	"crud/internal/domain"
	"crud/internal/pkg/authclient"
	"crud/internal/service"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

func ServerHandler(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, "*")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodPost)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodGet)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodDelete)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderContentType)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderAuthorization)

	if ctx.IsOptions() {
		return
	}

	token := ctx.Request.Header.Peek(fasthttp.HeaderAuthorization)
	log.Println(string(token) == "", !authclient.ValidateToken(string(token)), string(token) == "" || !authclient.ValidateToken(string(token)))
	if string(token) == "" || !authclient.ValidateToken(string(token)) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		log.Println("Get request", string(ctx.Method()), string(token), "error", fasthttp.StatusUnauthorized)
		return
	}

	switch {
	case ctx.IsGet():
		GetHandler(ctx)
	case ctx.IsDelete():
		DeleteHandler(ctx)
	case ctx.IsPost():
		PostHandler(ctx)
	}

}

// GetHandler godoc
// @Tags Recipes
// @Summary Чтение рецепта
// @Param id query string true "Айди рецепта"
// @Param Authorization header string true "Токен пользователя"
// @Success 200 {object} domain.Recipe
// @failure 400
// @failure 401
// @failure 404
// @failure 500
// @Router  / [get]
func GetHandler(ctx *fasthttp.RequestCtx) {
	isEntry := func(entryPoint string) bool {
		return strings.Contains(string(ctx.Path()), entryPoint)
	}
	log.Printf("request path:%s\n", ctx.Path())
	switch {
	case isEntry("/count-recipes"):
		getCountRecipes(ctx)
	case isEntry("/recipes"):
		getRecipes(ctx)
	default:
		getRecipeByID(ctx)
	}
}
func getRecipes(ctx *fasthttp.RequestCtx) {
	indBytes := ctx.QueryArgs().Peek("page")
	if len(indBytes) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Printf("len(indBytes) == 0")
		return
	}
	index, err := strconv.Atoi(string(indBytes))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Printf(" strconv.Atoi(string(indBytes))")
		return
	}
	sizBytes := ctx.QueryArgs().Peek("size")
	if len(sizBytes) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Printf("len(sizBytes) == 0")
		return
	}

	size, err := strconv.Atoi(string(sizBytes))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Printf(" strconv.Atoi(string(sizBytes))")
		return
	}

	rec, err := service.GetRecipes(index, size)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("service.GetRecipes(%d, %d):%s", index, size, err)
		return
	}
	log.Printf("recipes:%#v", rec)
	marshal, err := json.Marshal(rec)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
func getCountRecipes(ctx *fasthttp.RequestCtx) {

	rec := service.Count()
	marshal, err := json.Marshal(rec)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	log.Printf("count:%d", rec.Count)
	ctx.SetStatusCode(fasthttp.StatusOK)

}
func getRecipeByID(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	rec, err := service.Get(string(id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	marshal, err := json.Marshal(rec)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func DeleteHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := service.Delete(string(id)); err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func PostHandler(ctx *fasthttp.RequestCtx) {
	var rec domain.Recipe
	log.Println(string(ctx.PostBody()))
	if err := json.Unmarshal(ctx.PostBody(), &rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := service.AddOrUpd(&rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	resp := IdResponse{ID: rec.ID}

	marshal, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
