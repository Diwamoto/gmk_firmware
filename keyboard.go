package tgk

import (
	"fmt"
	"io"
	"os"
)

type Keyboard struct {
	Name string `json:"name"`
	VendorID int `json:"vendor_id"`
	ProductID int `json:"product_id"`
	ColSize int `json:"col_size"`
	RowSize int `json:"row_size"`
	Keys [][]string `json:"keys"`
	KeyMap map[string][]string `json:"key_map"`
}


func (&kb Keyboard) loadKBDef(json string) error {

	// ローカルのjsonファイルを構造体に起こす
	// ファイルを開く
	f, err := os.Open(json)
	if err != nil {
		return fmt.Errorf("failed to open keyboard definition file: %w", err)
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read keyboard definition file: %w", err)
	}

	json.Unmarshal(d, &kbDef)

	return nil
}
