package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"daud", "hidayat", "ramadhan"})
	_ = writer.Write([]string{"22", "20", "21"})

	writer.Flush()

}
