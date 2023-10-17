package germany

import (
	"encoding/json"
	"fmt"
)

func (f *Form) GetCountryName() string {
	return CountryName
}

func (f *Form) GetProcessingDays() int {
	return processingDays
}

func (f *Form) PrintForm() {
	res, _ := json.Marshal(f)
	fmt.Printf("Ваше заявление: " + string(res))
}

func (f *Form) ScanCountryData() {
	f.ScanBaseData()
	fmt.Println("Уровень немецкого языка:")
	fmt.Scan(&f.GermanLanguageLevel)
}
