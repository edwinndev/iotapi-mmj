package routes

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/edwinndev/iotapi-mmj/middlewares"
	"github.com/edwinndev/iotapi-mmj/models"
	"github.com/gorilla/mux"
)

func ConfigureUploadsRouter(router *mux.Router) {
	router.HandleFunc("", middlewares.AdminMiddleware(getUploadFiles)).Methods(http.MethodGet)
	router.HandleFunc("", middlewares.AdminMiddleware(uploadFileAWS)).Methods(http.MethodPost)
}

func getUploadFiles(w http.ResponseWriter, _ *http.Request) {
	var files []models.Upload
	err := database.Mysql.Find(&files)
	if err.Error != nil {
		commons.ApiNotFound(w, "Archivos no encontrados")
	} else {
		commons.ApiOK(w, files)
	}
}

func uploadFileAWS(w http.ResponseWriter, request *http.Request) {
	errorParsingForm := request.ParseMultipartForm(10 << 20)
	if errorParsingForm != nil {
		commons.ApiBadRequest(w, "Error, archivo no seleccionado")
		return
	}
	file, mult, err := request.FormFile("file")
	if err != nil {
		commons.ApiBadRequest(w, "Error al procesar archivo")
		return
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		commons.ApiBadRequest(w, "Error al leer archivo")
		return
	}

	config := &aws.Config{
		Region:      aws.String(commons.AwsRegion),
		Credentials: credentials.NewStaticCredentials(commons.AwsId, commons.AwsSecret, ""),
	}
	sess, _ := session.NewSession(config)

	uploader := s3manager.NewUploader(sess)
	var contentType = mult.Header.Get("Content-Type")
	input := &s3manager.UploadInput{
		Bucket:      aws.String(commons.AwsBucket),
		Key:         aws.String(commons.AwsPath + mult.Filename),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	}

	var response = models.Upload{
		FileName: mult.Filename,
		FileSize: mult.Size,
		FileType: contentType,
	}
	existsError := database.Mysql.First(&response, "file_name=?", mult.Filename).Error
	if existsError == nil {
		commons.ApiBadRequest(w, "El archivo ya existe")
		return
	}

	res, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		commons.ApiBadRequest(w, "Error al guardar archivo")
	} else {
		response.FilePath = res.Location

		_ = database.Mysql.Create(&response).Error
		commons.ApiCreated(w, response)
	}
}
