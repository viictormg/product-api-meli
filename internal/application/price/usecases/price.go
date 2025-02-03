package usecases

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/shopspring/decimal"
	"github.com/viictormg/product-api-meli/internal/application/price/dto"
	"github.com/viictormg/product-api-meli/internal/application/price/ports"
	"github.com/viictormg/product-api-meli/internal/domain/constants"
)

type PriceUsecaseIF interface {
	UploadPriceFile(ctx context.Context, file *multipart.FileHeader) error
}

type PriceUsecase struct {
	event ports.PriceEventyIF
}

func NewPriceUsecase(event ports.PriceEventyIF) PriceUsecaseIF {
	return &PriceUsecase{
		event: event,
	}
}

func (h *PriceUsecase) UploadPriceFile(ctx context.Context, file *multipart.FileHeader) error {
	data, err := extracDataFile(file)

	if err != nil {
		return err
	}

	for _, chunk := range data {
		message := ConverteData(chunk)
		h.event.SendPriceEvent(message)
	}

	h.event.Close()

	return nil
}

func ConverteData(data [][]string) []byte {
	items := []dto.PriceHistory{}

	for _, record := range data {
		priceProduct, err := decimal.NewFromString(record[2])
		if err != nil {
			fmt.Println("Error to convert price to decimal", record[2])
			continue
		}
		price := dto.PriceHistory{
			ProductID: record[0],
			OrderDate: record[1],
			Price:     priceProduct,
		}

		items = append(items, price)
	}

	messageDecode, err := json.Marshal(items)

	if err != nil {
		fmt.Println("Error to convert data to json")
	}
	return messageDecode
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

	return chunkData(records, constants.BatchSize), nil
}
