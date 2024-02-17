package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ //todos devem ser do mesmo tipo de dados, dentro dos [] é o tipo das chaves e fora é o tipo dos valores
		"nome":      "joe ",
		"sobrenome": "doe",
	}
	fmt.Println(usuario) // não da para acessar tipo usuario.nome, para acessar tem q ser usuario["nome"]
	fmt.Println(usuario["nome"])

	usuario["bairro"] = "centro"

	fmt.Println(usuario)

	//como aninhar maps, lembrando q todos os dado dentro do map, tem q ser do tipo map, OBS: aninhar muitos fica bem confuso
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "eduardo",
			"ultimo":   "mello",
		},
		"curso": {
			"nome":      "engenharia",
			"faculdade": "UFCAT",
		},
	}
	fmt.Println(usuario2)

	//adicionando outro valor no map
	usuario2["documento"] = map[string]string{
		"cpf": "1892371928",
	}
	fmt.Println(usuario2)

	//deletando valor
	delete(usuario2, "documento") //delete function = deletes the element with the specified key (m[key]) from the map. Primeio passa o mapa e depois o elemento que quer deletar
	fmt.Println(usuario2)
}
