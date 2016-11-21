package main

import "fmt"
import "regexp"
import "io/ioutil"
import "strings"

//import "os"

func check(e error) {
if e != nil {
panic(e)
}
}

func main() {



var c string
var elegir uint8

elegir = 0

fmt.Println("\n\t Eliga una opcion: \n")
fmt.Println("\t1 - Una Cadena \n ")
fmt.Print("\t2 - Un Archivo de Texto \n")
fmt.Scan(&elegir)

fmt.Println()
fmt.Println()

elecion1 := ("reconocer si es un identificador ,un numero,correo o un URL \n")

elecion2 := ("desde un archivo reconocer una cadena \n")

switch elegir {
case 1:
fmt.Println(elecion1)
fmt.Print("\tIngrese una cadena: \n")
fmt.Scan(&c)

pp1 := regexp.MustCompile("((^[a-zA-Z-0-9]+([_.-][a-zA-Z0-9]+)*)@[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*(.[a-z])$)")
p1 := pp1.MatchString(c)
pp2 := regexp.MustCompile("(^[0-9]*$)")
p2 := pp2.MatchString(c)
pp3 := regexp.MustCompile("(^[a-zA-Z]([a-zA-Z]|[0-9])*$)")
p3 := pp3.MatchString(c)
pp4 := regexp.MustCompile("(^(https?://)?([0-9a-zA-Z.-]+).([a-zA-Z]{2,6})([//a-zA-Z=.-]*)*/?$)")
p4 := pp4.MatchString(c)

if p1 == true {
fmt.Print("\n Esto es un correo.")
} else if p2 == true {
fmt.Print("\n Esto es una cadena numerica.")
} else if p3 == true {
fmt.Print("\n Esto es un identificador.")
} else if p4 == true {
fmt.Print("\n Esto es un URL.")
} else {
fmt.Print("\n No existe. ")
}

case 2:
fmt.Println(elecion2)
fmt.Print("contenido de archivo: \n")

texto, er := ioutil.ReadFile("C:\\ana\\ha.txt")
check(er)
fmt.Print(string(texto))

b := string(texto[:])
check(er)



v := strings.NewReplacer("{", " { ", "}", " } ", ";", " ; ", "(", " ( ", ")", " ) ", "+", " + ", "-", " - ", "*", " * ", "/", " / ", "%", " % ", "[", " [ ", "]", " ]", "==", " == ", "<=", " <= ", ">=", ">=", "!=", " != ")
m := v.Replace(b)
nuStrings := strings.Fields(m)

fmt.Println()
fmt.Println()

for a := range nuStrings {                                                       

igualacion, _ := regexp.MatchString("(^\\=$)", nuStrings[a])

opera   := regexp.MustCompile("^[\\+|\\-|\\*|\\/|\\%]$")
operadoresAritmeticos := opera.MatchString(nuStrings[a])

opera2    := regexp.MustCompile("((^\\<$)|(^\\>$)|(^\\=\\=$)|(^\\!\\=$)|(^\\>\\=$)|(^\\<\\=$)|(^\\^$))")
operadoresRelacio:= opera2.MatchString(nuStrings[a])

opera3     := regexp.MustCompile("(^\\&\\&$)|(^\\|\\|$)|(^\\!$)")
operadoreslogicos  := opera3.MatchString(nuStrings[a])

opera4    := regexp.MustCompile("(^\\&$)|(^nott$)|(^\\<\\<$)|(^\\>\\>$)|(^orr$)")
operadoresbits := opera4.MatchString(nuStrings[a])

tipos       := regexp.MustCompile("(^tipobyte$)|(^tiposbyte$)|(^tiposhort$)|(^tipoint$)|(^tipouint$)|(^tipolong$)|(^tipoulong$)|(^tipoushort$)")
datoentero  := tipos.MatchString(nuStrings[a])

tipos2      := regexp.MustCompile("(^tipofloat$)|(^tipodouble$)|(^tipodecimal$)")
datoReal    := tipos2.MatchString(nuStrings[a])

tipos3      := regexp.MustCompile("(^tipotrue$)|(^tipofalse$)")
datoBole    := tipos3.MatchString(nuStrings[a])

tipos4      := regexp.MustCompile("(^tipochar$)|(^tipostring$)")
datoAlfa    := tipos4.MatchString(nuStrings[a])

ciclos     := regexp.MustCompile("(^ciclodo$)|(^ciclofor$)|(^ciclowhi$)|(^ciclodowhi$)")
tipociclos  := ciclos.MatchString(nuStrings[a])

preguntas := regexp.MustCompile("(^dif$)|(^delse$)|(^delseif$)|(^dswitch$)")
selecion  := preguntas.MatchString(nuStrings[a])

instruci  := regexp.MustCompile("(^ibreak$)|(^igoto$)|(^icontinue$)|(^idefault$)|(^ireturn$)|(^iyield$)|(^ifixed$)|(^ilock$)")
ins   := instruci.MatchString(nuStrings[a])

atrapador := regexp.MustCompile("(^etry$)|(^ecatch$)|(^efinally$)")
errores    := atrapador.MatchString(nuStrings[a])

agrupa:= regexp.MustCompile("^[(\\[)|(\\])|(\\()|(\\))|(\\{)|(\\})]$")
agrupadores    := agrupa.MatchString(nuStrings[a])

puntuadu:= regexp.MustCompile("^[\\.|\\,|\\:]$")
puntuadores    := puntuadu.MatchString(nuStrings[a])

ope5:= regexp.MustCompile("(^asignasuma$)|(^asignaresta$)|(^asignamulti$)|(^asignadivisi$)|(^asignaporce$)|(^asignaand$)|(^asignaor$)|(^asignaxor$)|(^asidnadesplaizq$)|(^asignadespladere$)")
operadoAsi   := ope5.MatchString(nuStrings[a])

coment:= regexp.MustCompile("(^comentadiagonls)|(^comentagato$)|(^comentdiagonalaster$)")
com   := coment.MatchString(nuStrings[a])

imp:= regexp.MustCompile("(^iprint$)|(^iprintn$)|(^iprintl$)|(^iscan$)|(^iscanf$)|(^iscanln$)")
prim   := imp.MatchString(nuStrings[a])

ck:= regexp.MustCompile("(^dcheck$)|(^duncheck$)")
check   := ck.MatchString(nuStrings[a])

mod:= regexp.MustCompile("(^mpublic$)|(^mprivate$)|(^minternal$)|(^mprotected$)")
aceso  := mod.MatchString(nuStrings[a])

mod2:= regexp.MustCompile("(^devento$)|(^dreadonly$)|(^dstatic$)|(^dunsafe$)|(^dvirtual$)")
declara  := mod2.MatchString(nuStrings[a])

mod3:= regexp.MustCompile("(^iabstract$)|(^iclass$)|(^iasync$)|(^iextern$)|(^ivolatile$)|(^iextends$)|(^ifinal$)")
indica  := mod3.MatchString(nuStrings[a])

mod4:= regexp.MustCompile("(^econst$)|(^esealed$)")
especifica := mod4.MatchString(nuStrings[a])

mod5:= regexp.MustCompile("(^rnew$)|(^roverride$)")
rese := mod5.MatchString(nuStrings[a])

deci:= regexp.MustCompile("(^das$)|(^din$)|(^dis$)")
de := deci.MatchString(nuStrings[a])

decre:= regexp.MustCompile("(^\\-\\-$)|(^\\+\\+$)")
decremento := decre.MatchString(nuStrings[a])

pl:= regexp.MustCompile("(^pexplicit$)|(^pimplicit$)")
palab := pl.MatchString(nuStrings[a])

para:= regexp.MustCompile("(^pparams$)|(^pref$)|(^pout$)")
parametros := para.MatchString(nuStrings[a])

paraaceso:= regexp.MustCompile("(^abase$)|(^athis$)")
paraace := paraaceso.MatchString(nuStrings[a])

paracont:= regexp.MustCompile("(^padd$)|(^pasync$)|(^pdynamic$)|(^pglobal$)|(^pjoin$)|(^ppartial$)|(^pselect$)|(^pvar$)|(^palias$)|(^pawait$)|(^pfrom$)|(^pgroup$)|(^plet$)|(^ppartialmethod$)|(^pset$)|(^pwhereres$)|(^pwhereclaula$)|(^pasceding$)|(^pdescen$)|(^pget$)|(^pinto$)|(^porderby$)|(^premove$)|(^pvalue$)")
cont := paracont.MatchString(nuStrings[a])

palabra:= regexp.MustCompile("(^pfexed$)|(^pdelect$)|(^pnamespace$)|(^pobject$)|(^pstruc$)|(^pthrow$)|(^pthrows$)|(^ptypeof$)|(^pnull$)|(^psizeof$)|(^penum$)|(^pauto$)|(^pvolatile$)|(^psigned$)|(^pregister$)|(^pgets$)|(^ptypedef$)|(^punsigned$)|(^pvoid$)|(^punion$)|(^pinline$)|(^penurn$)|(^ptemplate$)|(^poperator$)|(^passert$)|(^ppass$)|(^pexcept$)|(^pimport$)|(^pexec$)|(^praise$)|(^pdef$)|(^pdel$)|(^plambda$)|(^pnone$)|(^pimplements$)|(^pinterface$)|(^pinstanceof$)|(^pnative$)|(^ppackage$)|(^pstrictfp$)|(^psuper$)|(^psynchronized$)|(^ptransient$)|(^pfuture$)|(^prest$)|(^pthreadsafe$)|(^pgeneric$)|(^pinner$)")
reservada := palabra.MatchString(nuStrings[a])

fmatematicas:= regexp.MustCompile("(^pabs$)|(^patn$)|(^pcos$)|(^pexp$)|(^plong$)|(^pround$)|(^psin$)|(^psqr$)|(^ptan$)")
fun := fmatematicas.MatchString(nuStrings[a])

fangulo:= regexp.MustCompile("(^pabsdecimal$)|(^pabsdouble$)|(^pabsint$)|(^pabsintd$)|(^pabsintt$)|(^pabsbyte$)|(^pabssingle$)|(^pacosdouble$)|(^pasindouble$)|(^atandouble$)")
funangulo:= fangulo.MatchString(nuStrings[a])

pp2 := regexp.MustCompile("(^[a-zA-Z]([a-zA-Z]|[0-9])*$)")
p2 := pp2.MatchString((nuStrings[a]))

pp3 := regexp.MustCompile("(^[0-9]*$)")
p3 := pp3.MatchString((nuStrings[a]))

pm := regexp.MustCompile("^\\;$")
pm2 := pm.MatchString((nuStrings[a]))



if operadoresAritmeticos  { 
fmt.Println("  |  Operadores Aritmeticos                  |", nuStrings[a])
} else if operadoresRelacio   {
fmt.Println("  |  Opera Relacionales                      | ", nuStrings[a])
} else if operadoreslogicos  {
fmt.Println("  |  Opera Logicos                           | ", nuStrings[a])
} else if operadoresbits {
fmt.Println("  |  Operadores de bits                      | ", nuStrings[a])
} else if datoentero {
fmt.Println("  |  Dato entero                             | ", nuStrings[a])
} else if datoReal {
fmt.Println("  |  Tipo de dato Real                       | ", nuStrings[a])
} else if datoBole {
fmt.Println("  |  Dato Boleano                            | ", nuStrings[a])
} else if datoAlfa {
fmt.Println("  |  Dato Alfanumerico                       | ", nuStrings[a])
} else if tipociclos {  
fmt.Println("  |  Ciclos                                  | ", nuStrings[a])
} else if selecion  {
fmt.Println("  |  Preguntas de Selecion                   | ", nuStrings[a])
} else if ins  { 
fmt.Println("  |  Instrucciones                           | ", nuStrings[a])
} else if errores  {
fmt.Println("  |  Atrapador de errores                    | ", nuStrings[a])
} else if agrupadores {
fmt.Println("  |  Agrupadores                             | ", nuStrings[a])
}  else if puntuadores {
fmt.Println("  |  Puntuadores                             | ", nuStrings[a])
}   else if operadoAsi   {
fmt.Println("  |  Operadores de Asignacion                | ", nuStrings[a])
}	else if com {
fmt.Println("  |  Coment                                  | ", nuStrings[a])
}   else if prim {
fmt.Println("  |  Imprim                                  | ", nuStrings[a])
}   else if check  {
fmt.Println("  |  Check                                   | ", nuStrings[a])
}   else if aceso  {
fmt.Println("  |  Modificadores de Aceso                  | ", nuStrings[a])
}    else if declara  {
fmt.Println("  |  Modifica Declara                        | ", nuStrings[a])
}    else if indica  {
fmt.Println("  |  Modifi Indica                           | ", nuStrings[a])
}    else if especifica  {
fmt.Println("  |  Especifica                              | ", nuStrings[a])
}    else if rese {
fmt.Println("  |  Modificadores Rese                      | ", nuStrings[a])
}     else if de {
fmt.Println("  |  Decisiones                              | ", nuStrings[a])
}     else if decremento {
fmt.Println("  |  Decremento                              | ", nuStrings[a])
}     else if  palab {
fmt.Println("  |  Palabrac Cl                             | ", nuStrings[a])
}     else if  parametros {
fmt.Println("  |  Parametros de Metodos                   | ", nuStrings[a])
}     else if paraace {
fmt.Println("  |  Aceso                                   | ", nuStrings[a])
}     else if cont  {
fmt.Println("  |  Pcontextuales                           | ", nuStrings[a])
}     else if  reservada  {
fmt.Println("  |  Palabras Reservadas                     | ", nuStrings[a])
}     else if  fun {
fmt.Println("  |  Funciones Matema                        | ", nuStrings[a])
}     else if funangulo {
fmt.Println("  |  Funcion Angulo                          | ", nuStrings[a])
} else if p2 {
fmt.Println("  |  Identificador                           | ", nuStrings[a])
} else if p3{
fmt.Println("  |  Numerico                                | ", nuStrings[a])
} else if pm2{
fmt.Println("  |  punto y coma                            | ", nuStrings[a])
} else if igualacion {
fmt.Println("  |  igualacion                              |", nuStrings[a])
} else{
fmt.Println("  |  No se encontro                          | ", nuStrings[a])
} 


}

}

}
