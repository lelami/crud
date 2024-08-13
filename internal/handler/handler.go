package handler

import (
	"crud/internal/domain"
	"crud/internal/pkg/authclient"
	"crud/internal/service/owner"
	"crud/internal/service/recipe"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
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
	tokenIsValid, userId := authclient.ValidateToken(string(token))

	log.Println(string(token) == "", !tokenIsValid, string(token) == "" || !tokenIsValid)
	if string(token) == "" || !tokenIsValid {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		log.Println("Get request", string(ctx.Method()), string(token), "error", fasthttp.StatusUnauthorized)
		return
	}

	switch {
	case ctx.IsGet():
		GetHandler(ctx)
	case ctx.IsDelete():
		DeleteHandler(ctx, userId)
	case ctx.IsPost():
		PostHandler(ctx, userId)
	}

}

func GetHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	rec, err := recipe.Get(string(id))
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

func DeleteHandler(ctx *fasthttp.RequestCtx, userId string) {
	id := ctx.QueryArgs().Peek("id")

	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := recipe.Delete(string(id)); err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func PostHandler(ctx *fasthttp.RequestCtx, userId string) {
	var rec domain.Recipe
	var recOwner domain.RecipeOwner

	log.Println(string(ctx.PostBody()))
	if err := json.Unmarshal(ctx.PostBody(), &rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	rec.CreatedBy = userId

	if err := recipe.AddOrUpd(&rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	resp := IdResponse{ID: rec.ID}

	recOwner.OwnerId = userId

	if err := owner.AddOrUpd(rec.ID, userId); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

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
