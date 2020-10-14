package main

import "fmt"

func main()  {
	var dia int
    var mes int

    fmt.Scanln(&dia)
    fmt.Scanln(&mes)

    if ( dia >= 21 && mes == 3 )||( dia <= 20 && mes == 4 ) {
        fmt.Println("aries")
    } else if ( dia >= 21 && mes == 4 )||( dia <= 20 && mes == 5 ) {
        fmt.Println("tauro")
    } else if ( dia >= 21 && mes == 5 )||( dia <= 21 && mes == 6 ) {
        fmt.Println("geminis")
    } else if ( dia >= 22 && mes == 6 )||( dia <= 22 && mes == 7 ) {
        fmt.Println("cancer")
    } else if ( dia >= 23 && mes == 7 )||( dia <= 22 && mes == 8 ) {
        fmt.Println("leo")
    } else if ( dia >= 23 && mes == 8 )||( dia <= 22 && mes == 9 ) {
        fmt.Println("viergo")
    } else if ( dia >= 23 && mes == 9 )||( dia <= 22 && mes == 10 ) {
        fmt.Println("libra")
    } else if ( dia >= 23 && mes == 10 )||( dia <= 22 && mes == 11 ) {
        fmt.Println("escorpio")
    }  else if ( dia >= 23 && mes == 11 )||( dia <= 21 && mes == 12 ) {
        fmt.Println("sagitario")
    } else if ( dia >= 22 && mes == 12 )||( dia <= 20 && mes == 1 ) {
        fmt.Println("capricornio")
    } else if ( dia >= 21 && mes == 1 )||( dia <= 18 && mes == 2 ) {
        fmt.Println("acuario")
    } else if ( dia >= 19 && mes == 2 )||( dia <= 20 && mes == 3 ) {
        fmt.Println("piscis")
    }
}
