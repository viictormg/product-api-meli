package dto

type ProductResponseDTO struct {
	ItemID   string      `json:"item_id"`
	Price    float64     `json:"price"`
	Anomaly  bool        `json:"anomaly"`
	Metadata interface{} `json:"metadata"`
}
