package main

import (
	"fmt"
    "time"
    "strconv"
   // "strings"
)
func proceso(c1 chan string, p *Proceso) {
	i := uint64(0);
    for {
		f := "Id "
		f += strconv.FormatUint(p.Id, 10)
		f += " "
		g  := strconv.FormatUint(i, 10)
        f += g
        c1 <- f
		i = i + 1
        time.Sleep(time.Millisecond * 500)
    }
}
type Proceso struct{
	Id uint64
	Canal int
}
func (p *Proceso) Start(c1 chan string) {
    go proceso(c1,p)
    go func (){
        for{
            <-c1
        }
    }()
}
func (p *Proceso)Stop() {
    fmt.Println("En Mantenimiento ")
}
type Procesos struct{
	Lista []Proceso
}

func main()  {
    lista := Procesos{}
    c1 := make(chan string)
    c2 := make(chan string)
    var id uint64
    id = 0
    mostrar :=0
    for {
        var c string
        if mostrar == 0{
            fmt.Println("0) Agregar Proceso\n1) Mostrar Procesos\n2) Eliminar Proceso\n3) Salir")
            fmt.Scanln(&c)
        }        
        if c == "1"{
            go func ()  {
                for{
                    select{
                    case msg :=<-c1:
                        fmt.Println(msg)
                    case <-c2:
                        return
                    }
                }
            }()
            var input string
            fmt.Scanln(&input)
            c2 <- input
        }else if c =="2"{
            var valor int
            fmt.Scan(&valor)
            for _,i := range lista.Lista{
                i.Stop()
            }
        }else if c == "3"{
            close(c1)
            break
        }else if c == "0"{            
            p := Proceso{Id:id,Canal:1}
            id +=1
            lista.Lista = append(lista.Lista,p)
            p.Start(c1)
        }
    }
}