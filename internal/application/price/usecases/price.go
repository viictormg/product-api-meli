package usecases

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"slices"
	"strconv"

	"github.com/viictormg/product-api-meli/internal/domain/constants"
)

type PriceUsecaseIF interface {
	UploadPriceFile(*multipart.FileHeader) error
}

func NewPriceUsecase() PriceUsecaseIF {
	return &PriceUsecase{}
}

type PriceUsecase struct {
}

type PriceHistory struct {
	ProductID string  `json:"product_id"`
	OrderDate string  `json:"order_date"`
	Price     float64 `json:"price"`
}

func (h *PriceUsecase) UploadPriceFile(file *multipart.FileHeader) error {
	data, err := extracDataFile(file)

	if err != nil {
		return err
	}

	for _, chunk := range data {

		message := ConverteData(chunk)

		PushPrice("price", message)

	}

	return nil
}

func ConverteData(data [][]string) []byte {
	items := []PriceHistory{}

	productIDs := []string{}

	for _, record := range data {
		priceProduct, _ := strconv.ParseFloat(record[2], 32)

		if !slices.Contains(productIDs, record[0]) {
			productIDs = append(productIDs, record[0])
		}

		price := PriceHistory{
			ProductID: record[0],
			OrderDate: record[1],
			Price:     priceProduct,
		}

		items = append(items, price)
	}

	fmt.Println("Product IDs: ", productIDs)
	jsonMessage, _ := json.Marshal(items)
	return jsonMessage
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
