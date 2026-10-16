package main 

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main () {
	if len(os.Args) < 2 {
		fmt.Println("veuillez entrer un fichier")
		return
	}
	nomDuFichier := os.Args[1]
	donnees, err := os.ReadFile(nomDuFichier)
    if err != nil {
        fmt.Println("Erreur lors de la lecture :", err)
        return
	}
	texte := string(donnees)
	arguments := strings.Split(texte, " ")
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "(hex)" && i > 0 {
			nombre, err := strconv.ParseInt(arguments[i-1] ,16 ,64)
			if err != nil {
        		fmt.Println("Erreur")
        		return
			}
			txtDecimalHex := strconv.FormatInt(nombre, 10)
			arguments[i-1] = txtDecimalHex
			arguments = append(arguments[:i], arguments[i+1:]...)
			i--
		}
		if arguments[i] == "(bin)" && i > 0 {
			nombre, err := strconv.ParseInt(arguments[i-1] ,2 ,64)
			if err != nil {
        		fmt.Println("Erreur")
        		return
			}
			txtDecimalBin := strconv.FormatInt(nombre, 10)
			arguments[i-1] = txtDecimalBin
			arguments = append(arguments[:i], arguments[i+1:]...)
			i--
		}
		if arguments[i] == "(up)" && i > 0 {
			txtMaj := strings.ToUpper(arguments[i-1])
			arguments[i-1] = txtMaj
			arguments = append(arguments[:i], arguments[i+1:]...)
			i--
		}
		if arguments[i] == "(up," {
			txtNbr := arguments[i+1] 
			txtNbr = strings.TrimSuffix(txtNbr, ")")
			nombre, err := strconv.Atoi(txtNbr)
			if err != nil {
        		fmt.Println("Erreur")
        		return
			}
			if nombre > i {
				nombre = i
			}
			for j := 1 ; j <= nombre ; j++ {
				arguments[i-j] = strings.ToUpper(arguments[i-j])
			}
			arguments = append(arguments[:i], arguments[i+1:]...)
			i--
		}
		if arguments[i] == "(low)" && i > 0 {
			txtMin := strings.ToLower(arguments[i-1])
			arguments[i-1] = txtMin
			arguments = append(arguments[:i], arguments[i+2:]...)
			i--
		}
		if arguments[i] == "(low," {
			txtNbr := arguments[i+1] 
			txtNbr = strings.TrimSuffix(txtNbr, ")")
			nombre, err := strconv.Atoi(txtNbr)
			if err != nil {
        		fmt.Println("Erreur")
        		return
			}
			if nombre > i {
				nombre = i
			}
			for j := 1 ; j <= nombre ; j++ {
				arguments[i-j] = strings.ToLower(arguments[i-j])
			}
			arguments = append(arguments[:i], arguments[i+2:]...)
			i--
		}
		if arguments[i] == "(cap)" && i > 0 {
			capMaj := strings.ToUpper(arguments[i-1][:1])
			arguments[i-1] = capMaj + arguments[i-1][1:]
			arguments = append(arguments[:i], arguments[i+1:]...)
			i--
		}
		if arguments[i] == "(cap," && i > 0 {
			txtNbr := arguments[i+1] 
			txtNbr = strings.TrimSuffix(txtNbr, ")")
			nombre, err := strconv.Atoi(txtNbr)
			if err != nil {
        		fmt.Println("Erreur")
        		return
			}
			if nombre > i {
			nombre = i
			}
			for j := 1 ; j <= nombre ; j++ {
				arguments[i-j] = strings.ToUpper(arguments[i-j][:1]) + arguments[i-j][1:]
			}
			arguments = append(arguments[:i], arguments[i+2:]...)
			i--
		}

