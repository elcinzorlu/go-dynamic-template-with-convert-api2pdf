package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/elcinzorlu/dynamic-template-with-convert-api2pdf/models"
)

func main() {
	var buf bytes.Buffer
	var resultString string
	//create a person struct
	persons := []models.Person{
		{Isim: "John", Yas: 30, Adres: models.Address{Sokak: "123 Main St", Sehir: "New York", PostaKodu: "10001"}},
		{Isim: "Jane", Yas: 25, Adres: models.Address{Sokak: "456 Elm St", Sehir: "Los Angeles", PostaKodu: "90001"}},
		{Isim: "Alice", Yas: 35, Adres: models.Address{Sokak: "789 Oak St", Sehir: "Chicago", PostaKodu: "60601"}},
	}
	//parse template
	tmpl, _ := template.ParseFiles("templates/dynamictemplate.html")
	err := tmpl.Execute(&buf, persons)
	if err != nil {
	} else {
		resultString = buf.String()
	}

	apiUrl := "https://v2.api2pdf.com/chrome/pdf/html"

	// Send POST request
	data := models.Payload{
		HTML:     resultString,
		FileName: "test.pdf",
	}
	data.Storage.Method = "PUT"
	data.Storage.URL = "https://presignedurl"
	data.Storage.ExtraHTTPHeaders = make(map[string]string)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON convert error:", err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Create request error:", err)
		return
	}

	//Set Authorization and YOUR-API-KEY
	req.Header.Set("Authorization", "YOUR-API-KEY")

	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Request Succes.")
	} else {
		fmt.Printf("Error code: %d\n", response.StatusCode)
		fmt.Printf("Error message: %s\n", response.Status)
	}

	var responseData models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		fmt.Println("JSON response error:", err)
		return
	}
	resp := models.ResponseData{
		ResponseId: responseData.ResponseId,
		MbOut:      responseData.MbOut,
		Cost:       responseData.Cost,
		Seconds:    responseData.Seconds,
		Error:      responseData.Error,
		Success:    responseData.Success,
		FileUrl:    responseData.FileUrl,
	}
	fmt.Println(resp)
}
