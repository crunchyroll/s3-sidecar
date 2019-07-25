package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/crunchyroll/s3-sidecar/handler"
	"github.com/crunchyroll/s3-sidecar/logging"
	"github.com/gorilla/mux"
)

type ItemController struct {
	mediaHandler *handler.S3Handler
	logger       *logging.Logger
}

func NewItemController(region, bucket string, logger *logging.Logger) *ItemController {
	return &ItemController{mediaHandler: handler.NewS3Handler(region, bucket), logger: logger}
}

// GetContent - gets the content from the Vod Media Bucket based on the request
func (it *ItemController) GetContent(resp http.ResponseWriter, req *http.Request) {
	nowTime := time.Now()
	vars := mux.Vars(req)
	assetID := vars["asset"] + vars["item"]

	it.logger.Info("looking for ", logging.DataFields{"item": assetID})
	// validate the asset id
	err, content := it.mediaHandler.GetItemContent(assetID)
	if err != nil {
		it.logger.Error("ItemController:Failure", logging.DataFields{"reason": "invalid asset-id", "error": err.Error()})
		http.Error(resp, fmt.Sprintf("400 Bad Request, %s ", err.Error()), http.StatusBadRequest)
		return
	}

	//TODO: measure transaction with newrelic

	duration := time.Now().Sub(nowTime).Seconds() //use nano seconds for now
	it.logger.Info("success", logging.DataFields{"asset-id": assetID, "content-length": len(content), "duration": duration})
	resp.Write([]byte(content))
	resp.WriteHeader(http.StatusOK)
}
