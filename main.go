package main

import "fmt"

func main()  {
	var dia int
    var mes int

    fmt.Scanln(&dia)
    fmt.Scanln(&mes)

    if ( dia >= 21 && mes == 3 )||( dia <= 20 && mes == 4 ) {
        fmt.Println("Aries")
    } else if ( dia >= 21 && mes == 4 )||( dia <= 20 && mes == 5 ) {
        fmt.Println("Tauro")
    } else if ( dia >= 21 && mes == 5 )||( dia <= 21 && mes == 6 ) {
        fmt.Println("Geminis")
    } else if ( dia >= 22 && mes == 6 )||( dia <= 22 && mes == 7 ) {
        fmt.Println("Cancer")
    } else if ( dia >= 23 && mes == 7 )||( dia <= 22 && mes == 8 ) {
        fmt.Println("Leo")
    } else if ( dia >= 23 && mes == 8 )||( dia <= 22 && mes == 9 ) {
        fmt.Println("Viergo")
    } else if ( dia >= 23 && mes == 9 )||( dia <= 22 && mes == 10 ) {
        fmt.Println("Libra")
    } else if ( dia >= 23 && mes == 10 )||( dia <= 22 && mes == 11 ) {
        fmt.Println("Escorpio")
    }  else if ( dia >= 23 && mes == 11 )||( dia <= 21 && mes == 12 ) {
        fmt.Println("Sagitario")
    } else if ( dia >= 22 && mes == 12 )||( dia <= 20 && mes == 1 ) {
        fmt.Println("Capricornio")
    } else if ( dia >= 21 && mes == 1 )||( dia <= 18 && mes == 2 ) {
        fmt.Println("Acuario")
    } else if ( dia >= 19 && mes == 2 )||( dia <= 20 && mes == 3 ) {
        fmt.Println("Piscis")
    }
}
