package inttoptbr

import (
	"fmt"
	"strconv"
	"strings"
)

type Extenso struct {
	Unidade  []string
	DezVinte []string
	Dezena   []string
	Centena  []string
	Milhar   []string
}

func TranscreveValor(numero int64) (string, error) {

	if numero <= 0 || numero > 999999999 {
		return "", fmt.Errorf("Programa limitado a numero maior que 0 e menor que 1 Bilhao")
	}

	var ValorEmExtenso string
	var err error
	numeros := makeExtenso()

	strValue := fmt.Sprintf("%d", numero)

	splitted := splitEach(strValue, 3)
	numCasas := len(splitted)

	var classe1, classe2, classe3 string

	//trancreve o valor enviado separado por classes

	if numCasas == 3 {
		classe3, err = transcreveCentena(splitted[numCasas-1])
		if err != nil {
			return "", err
		}

		if classe3 == numeros.Unidade[1] {
			classe3 = fmt.Sprint(classe3, numeros.Milhar[1])

		} else {

			classe3 = fmt.Sprint(classe3, numeros.Milhar[2])
		}
		numCasas--

	}

	if numCasas == 2 {

		classe2, err = transcreveCentena(splitted[numCasas-1])
		if err != nil {
			return "", err
		}

		if classe2 != "" {

			classe2 = fmt.Sprint(classe2, numeros.Milhar[0])
		}

		numCasas--
	}

	if numCasas == 1 {

		classe1, err = transcreveCentena(splitted[numCasas-1])
		if err != nil {
			return "", err
		}
	}

	//concatena os valores com a devida conjunção

	if classe3 != "" {

		sToInt, err := strconv.Atoi(splitted[1])
		if err != nil {
			return "", err
		}

		//se houver numeros na classe das milhares(2) ou das unidades(1) && na classe das milhares(2) nao possuir uma centena sem dezena ou unidade, ou seja qualquer numero terminado com 00, entao concatena " e "
		if (classe2 != "" || classe1 != "") && (!(sToInt > 99 && !strings.Contains(splitted[1], "00")) && classe1 == "") {

			ValorEmExtenso = fmt.Sprint(ValorEmExtenso, classe3, "e ")
		} else {

			ValorEmExtenso = fmt.Sprint(ValorEmExtenso, classe3)
		}
	}

	if classe2 != "" {

		sToInt, err := strconv.Atoi(splitted[0])
		if err != nil {
			return "", err
		}

		// se na classe da unidade(1) possuir uma centena sem dezena ou unidade, ou seja qualquer numero terminado com 00, entao NAO concatena " e "
		if sToInt > 99 && !strings.Contains(splitted[0], "00") || classe1 == "" {

			ValorEmExtenso = fmt.Sprint(ValorEmExtenso, classe2)
		} else {

			ValorEmExtenso = fmt.Sprint(ValorEmExtenso, classe2, "e ")
		}
	}

	ValorEmExtenso = fmt.Sprint(ValorEmExtenso, classe1)

	return ValorEmExtenso, nil
}

func transcreveCentena(centena string) (string, error) {

	if len(centena) > 3 {
		return "", fmt.Errorf("%v está além de uma centena!", centena)
	}

	centenaInt, err := strconv.Atoi(centena)
	if err != nil {
		return "", fmt.Errorf("permitidos até 3 numero, somente numeros! %v", err)
	}

	centena = strconv.Itoa(centenaInt)

	var sToInt int
	var valorTranscrito string
	numeros := makeExtenso()

	valorInvertido := strings.Split(reverse(centena), "")
	numCasas := len(centena)

	if numCasas == 3 {

		if valorInvertido[1] != "0" || valorInvertido[0] != "0" {

			sToInt, _ = strconv.Atoi(valorInvertido[2])

			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Centena[sToInt], " e ")
		} else if valorInvertido[2] == "1" {

			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Centena[0])
		} else if valorInvertido[2] != "0" {
			sToInt, _ = strconv.Atoi(valorInvertido[2])

			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Centena[sToInt])
		}
		valorInvertido = valorInvertido[:2]
		numCasas--

	}

	if numCasas == 2 {
		if valorInvertido[0] == "0" {
			sToInt, _ = strconv.Atoi(valorInvertido[1])

			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Dezena[sToInt])
		} else if valorInvertido[1] == "1" && valorInvertido[0] != "0" {
			sToInt, _ = strconv.Atoi(valorInvertido[0])

			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.DezVinte[sToInt])
		} else {
			sToInt, _ = strconv.Atoi(valorInvertido[1])

			if sToInt != 0 {

				valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Dezena[sToInt], " e ")
			}
		}

	}

	if numCasas == 1 {
		if valorInvertido[0] != "0" {
			sToInt, _ = strconv.Atoi(valorInvertido[0])
			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Unidade[sToInt])
		}

	} else {

		if valorInvertido[0] != "0" && valorInvertido[1] != "1" {
			sToInt, _ = strconv.Atoi(valorInvertido[0])
			valorTranscrito = fmt.Sprint(valorTranscrito, numeros.Unidade[sToInt])
		}
	}

	return valorTranscrito, nil
}

func makeExtenso() Extenso {
	var numeros Extenso

	numeros.Unidade = append(numeros.Unidade, "")
	numeros.Unidade = append(numeros.Unidade, "um")
	numeros.Unidade = append(numeros.Unidade, "dois")
	numeros.Unidade = append(numeros.Unidade, "três")
	numeros.Unidade = append(numeros.Unidade, "quatro")
	numeros.Unidade = append(numeros.Unidade, "cinco")
	numeros.Unidade = append(numeros.Unidade, "seis")
	numeros.Unidade = append(numeros.Unidade, "sete")
	numeros.Unidade = append(numeros.Unidade, "oito")
	numeros.Unidade = append(numeros.Unidade, "nove")

	numeros.DezVinte = append(numeros.DezVinte, "")
	numeros.DezVinte = append(numeros.DezVinte, "onze")
	numeros.DezVinte = append(numeros.DezVinte, "doze")
	numeros.DezVinte = append(numeros.DezVinte, "treze")
	numeros.DezVinte = append(numeros.DezVinte, "quatorze")
	numeros.DezVinte = append(numeros.DezVinte, "quinze")
	numeros.DezVinte = append(numeros.DezVinte, "dezesseis")
	numeros.DezVinte = append(numeros.DezVinte, "dezesete")
	numeros.DezVinte = append(numeros.DezVinte, "dezoito")
	numeros.DezVinte = append(numeros.DezVinte, "dezenove")

	numeros.Dezena = append(numeros.Dezena, "")
	numeros.Dezena = append(numeros.Dezena, "dez")
	numeros.Dezena = append(numeros.Dezena, "vinte")
	numeros.Dezena = append(numeros.Dezena, "trinta")
	numeros.Dezena = append(numeros.Dezena, "quarenta")
	numeros.Dezena = append(numeros.Dezena, "cinquenta")
	numeros.Dezena = append(numeros.Dezena, "sessenta")
	numeros.Dezena = append(numeros.Dezena, "setenta")
	numeros.Dezena = append(numeros.Dezena, "oitenta")
	numeros.Dezena = append(numeros.Dezena, "noventa")

	numeros.Centena = append(numeros.Centena, "cem")
	numeros.Centena = append(numeros.Centena, "cento")
	numeros.Centena = append(numeros.Centena, "duzentos")
	numeros.Centena = append(numeros.Centena, "trezentos")
	numeros.Centena = append(numeros.Centena, "quatrocentos")
	numeros.Centena = append(numeros.Centena, "quinhentos")
	numeros.Centena = append(numeros.Centena, "seiscentos")
	numeros.Centena = append(numeros.Centena, "setecentos")
	numeros.Centena = append(numeros.Centena, "oitocentos")
	numeros.Centena = append(numeros.Centena, "novecentos")

	numeros.Milhar = append(numeros.Milhar, " mil ")
	numeros.Milhar = append(numeros.Milhar, " milhão ")
	numeros.Milhar = append(numeros.Milhar, " milhões ")

	return numeros
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func splitEach(s string, n int) []string {
	var min int
	var a []string

	for i := len(s); i >= 0; i -= n {
		min = i - n
		if min < 0 {
			min = 0
		}
		if s[min:i] != "" {

			a = append(a, s[min:i])
		}
	}

	return a
}
