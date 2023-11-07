package models

type Payload struct {
	HTML     string `json:"html"`
	FileName string `json:"fileName"`
	Storage  struct {
		Method           string            `json:"method"`
		URL              string            `json:"url"`
		ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders"`
	} `json:"storage"`
}

type ResponseData struct {
	ResponseId string  `json:"responseId"`
	MbOut      float64 `json:"mbOut"`
	Cost       float64 `json:"cost"`
	Seconds    float64 `json:"seconds"`
	Error      string  `json:"error"`
	Success    bool    `json:"success"`
	FileUrl    string  `json:"fileUrl"`
}

type Address struct {
	Sokak     string
	Sehir     string
	PostaKodu string
}

type Person struct {
	Isim  string
	Yas   int
	Adres Address // Adres tipinde bir başka struct
}
