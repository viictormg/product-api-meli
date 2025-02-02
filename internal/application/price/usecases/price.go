package usecases

import (
	"encoding/csv"
	"mime/multipart"
)

type PriceUsecaseIF interface {
	UploadPriceFile(*multipart.FileHeader) error
}

func NewPriceUsecase() PriceUsecaseIF {
	return &PriceUsecase{}
}

type PriceUsecase struct{}

func (h *PriceUsecase) UploadPriceFile(file *multipart.FileHeader) error {
	_, err := extracDataFile(file)

	if err != nil {
		return err
	}

	PushPrice("price", []byte("price"))

	return nil
}

func chunkData(data [][]string, chunkSize int) [][][]string {
	var chunks [][][]string
	for chunkSize < len(data) {
		data, chunks = data[chunkSize:], append(chunks, data[0:chunkSize:chunkSize])
	}
	return append(chunks, data)
}

func extracDataFile(file *multipart.FileHeader) ([][][]string, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(src)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return chunkData(records, 500), nil
}
