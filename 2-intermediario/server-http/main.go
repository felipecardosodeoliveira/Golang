package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func fetchCep(cep string) (CEP, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao buscar endereço %v \n", err)
		return CEP{}, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler corpo da requisição %v \n", err)
		return CEP{}, err
	}

	var data CEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v \n", err)
	}
	return data, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusFound)
			return
		}
		cepParam := r.URL.Query().Get("cep")
		if cepParam == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cep, err := fetchCep(cepParam)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		result, err := json.Marshal(cep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(result)

		// json.NewEncoder(w).Encode(cep)
	})
	fmt.Println("Server running")
	http.ListenAndServe(":8000", nil)
}
