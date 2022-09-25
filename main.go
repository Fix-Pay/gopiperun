package gopiperun

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/armando-couto/goutils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func FindAllOportunities(token string) {
	url := "https://api.pipe.run/v1/deals"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllOportunities", Error: err.Error()})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllOportunities", Error: err.Error()})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllOportunities", Error: err.Error()})
		return
	}

	var response = OportunityResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllOportunities", Error: err.Error()})
		return
	}
}

func FindAllCompanies(token string) {
	url := "https://api.pipe.run/v1/companies"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllCompanies", Error: err.Error()})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllCompanies", Error: err.Error()})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllCompanies", Error: err.Error()})
		return
	}

	var response = CompanyResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllCompanies", Error: err.Error()})
		return
	}

	//fmt.Println(string(body)) //TODO
}

func FindCompanie(token, cpfCnpj string) Company {
	response := Company{}
	url := `https://api.pipe.run/v1/companies?show=123&cnpj=` + cpfCnpj
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindCompanie", Error: err.Error()})
		return response
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindCompanie", Error: err.Error()})
		return response
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindCompanie", Error: err.Error()})
		return response
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindCompanie", Error: err.Error()})
		return response
	}
	return response
}

func CreateCompany(token, cpfCnpj, socialReason, cep, street, number, complement, neighborhood, email, phone string) Company {
	var response = Company{}

	url := "https://api.pipe.run/v1/companies"
	method := "POST"

	payload := strings.NewReader(`{"name": "` + socialReason + `",` +
		`"cnpj": "` + cpfCnpj + `",` +
		`"address_postal_code": "` + cep + `",` +
		`"address": "` + street + `",` +
		`"address_number": "` + number + `",` +
		`"address_complement": "` + complement + `",` +
		`"district": "` + neighborhood + `",` +
		`"country": "Brasil",` +
		`"cep": "` + cep + `",` +
		`"city_id": null,` +
		`"contactEmails":[` +
		`{` +
		`"email":"` + email + `"` +
		`}` +
		`],` +
		`"contactPhones":[` +
		`{` +
		`"phone":"` + phone + `"` +
		`}` +
		`]` +
		`}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateCompany", Error: err.Error()})
		return response
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateCompany", Error: err.Error()})
		return response
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateCompany", Error: err.Error()})
		return response
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateCompany", Error: err.Error()})
		return response
	}
	return response
}

func FindAllPeoples(token string) {
	url := "https://api.pipe.run/v1/persons"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllPeoples", Error: err.Error()})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllPeoples", Error: err.Error()})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllPeoples", Error: err.Error()})
		return
	}

	var response = PersonResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindAllPeoples", Error: err.Error()})
		return
	}

	//fmt.Println(string(body)) //TODO
}

func FindUser(token, usuarioNome string) User {
	url := `https://api.pipe.run/v1/users?show=123&name=` + url.QueryEscape(usuarioNome)
	method := "GET"

	var u User

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindUser", Error: err.Error()})
		return u
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindUser", Error: err.Error()})
		return u
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindUser", Error: err.Error()})
		return u
	}

	var response = UserResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "FindUser", Error: err.Error()})
		return u
	}

	if len(response.Data) > 0 {
		u = response.Data[0]
	}
	return u
}

func CreateFile(token, url, dealId string) error {
	aux := strings.Split(url, "/")
	fileName := aux[len(aux)-1]

	err := DownloadFile(url, fileName)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}

	defer os.Remove("tmp/" + fileName)

	urlReq := "https://api.pipe.run/v1/files"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("tmp/" + fileName)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base("tmp/"+fileName))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return errFile1
	}
	_ = writer.WriteField("deal_id", dealId)
	err = writer.Close()
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, payload)

	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Token", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}
	//var response = FileCreateResponse{}
	var response = FileCreateResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		goutils.CreateFileDay(goutils.Message{File: "CreateFile", Error: err.Error()})
		return err
	}

	if !response.Success {
		return errors.New("CreateFile: error inserir arquivo no piperun")
	}
	return err
}

func DownloadFile(url string, fileName string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	filepath := "tmp/" + fileName

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
