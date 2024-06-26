// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// type Address struct {
// 	Cep         string `json:"cep"`
// 	Logradouro  string `json:"logradouro"`
// 	Complemento string `json:"complemento"`
// 	Bairro      string `json:"bairro"`
// 	Localidade  string `json:"localidade"`
// 	Uf          string `json:"uf"`
// 	Ibge        string `json:"ibge"`
// 	Gia         string `json:"gia"`
// 	Ddd         string `json:"ddd"`
// 	Siafi       string `json:"siafi"`
// }

// func main() {
// 	for _, cep := range os.Args[1:] {
// 		req, err := http.Get(" http://viacep.com.br/ws/" + cep + "/json")
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao fazer requisições %v", err)
// 		}
// 		defer req.Body.Close()

// 		res, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao ler respostas: %v \n", err)
// 		}

// 		var data Address
// 		err = json.Unmarshal(res, &data)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao ler respostas: %v \n", err)
// 		}

// 		file, err := os.Create("cidade.txt")
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo %v \n", err)
// 		}
// 		defer file.Close()

//			_, err = file.WriteString(fmt.Sprintf("CEP %s, Localidade %s UF %s", data.Cep, data.Localidade, data.Uf))
//		}
//	}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
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

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar endereço %v \n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler corpo da requisição %v \n", err)
		}

		var data Address
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v \n", err)
		}

		// CEP 81490-420, Localidade Curitiba UF PR
		f, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v \n", err)
		}
		defer f.Close()

		n3, err := f.WriteString(fmt.Sprintf("CEP %s, Localidade %s UF %s", data.Cep, data.Localidade, data.Uf))
		if err != nil {
			panic(err)
		}
		// fmt.Sprintf("CEP %s, Localidade %s UF %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("wrote bytes\n", n3)
	}
}
