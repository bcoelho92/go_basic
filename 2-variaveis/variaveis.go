package main

import "fmt"

func main() {
	const (
		empresa = "Empresa XYZ"
	)
	var (
		placa_carro    string = "ABC-1234"
		modelo_carro   string = "Toyota Corolla"
		ano_fabricacao int    = 2020
	)
	idade := 33
	ano := 1992
	dia := 27
	mes := 8
	nome_completo := "João da Silva"
	rico := false
	email := "joao.silva@email.com"
	salario := 2500.50
	data_admissao := "01/01/2020"

	va, vr, valorVr, valorVa := true, false, 0.00, 500.00

	cargo := "Analista de Sistemas"

	fmt.Println(
		"- Empresa:", empresa,
		"- Cargo:", cargo,
		"- Idade:", idade,
		"- Data de nascimento:", dia, "/", mes, "/", ano,
		"- Nome completo:", nome_completo,
		"- É rico:", rico,
		"- Email:", email,
		"- Data de admissão:", data_admissao,
		"- Salário:", fmt.Sprintf("R$ %.2f", salario),
		"- Placa do carro:", placa_carro,
		"- Modelo do carro:", modelo_carro,
		"- Ano de fabricação:", ano_fabricacao,
		"- Vale refeição:", vr,
		"- Vale refeição:", fmt.Sprintf("R$ %.2f", valorVr),
		"- Vale alimentação:", va,
		"- Vale alimentação:", fmt.Sprintf("R$ %.2f", valorVa),
	)
	// troca de valor de variáveis
	v1 := 10
	v2 := 20

	fmt.Println("Antes da troca: v1 =", v1, "v2 =", v2)

	v1, v2 = v2, v1

	fmt.Println("Depois da troca: v1 =", v1, "v2 =", v2)
}
