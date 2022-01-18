package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "128.199.1.222"
	port     = 5432
	user     = "modulo4"
	password = "modulo4"
	dbname   = "giinwedb"
)

const (
	host1     = "castor.db.elephantsql.com"
	port1     = 5432
	user1     = "imprkjzn"
	password1 = "zjRI3c4nmVoYRRclBfrWiJKCMcuQFa0Z"
	dbname1   = "imprkjzn"
)

type ubigeo struct {
	inei  string
	lugar string
}
type lugarDelta struct {
	id           string
	lugarDeltaIT string
}

func main() {

	// connection string
	/*psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	rows, err := db.Query(`
							SELECT "lugar"
							FROM "oferta"
							`)
	CheckError(err)

	defer rows.Close()
	//count := 0
	for rows.Next() {
		var lugar string
		lugares := make([]string, 0)
		lugares = append(lugares, lugar)
		err = rows.Scan(&lugar)
		CheckError(err)
	}*/

	//fmt.Println(count)

	// connection string
	psqlconn1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host1, port1, user1, password1, dbname1)

	// open database
	db1, err := sql.Open("postgres", psqlconn1)
	CheckError(err)

	// close database
	defer db1.Close()

	// check db
	err = db1.Ping()
	CheckError(err)

	fmt.Println("Conexion exitosa a la BD de prueba")

	rows1, err := db1.Query(`
								SELECT id_oferta, lugar
								FROM lugar
								`)
	CheckError(err)

	defer rows1.Close()
	count1 := 0
	lugaresDelaIT := make([]lugarDelta, 0)
	for rows1.Next() {
		var lugar lugarDelta
		err = rows1.Scan(&lugar.id, &lugaresDelaIT)
		lugaresDelaIT = append(lugaresDelaIT, lugar)
		//var str1 = []rune(departamento)
		//var str2 = []rune("PERU")
		//rst := levenshtein(str1, str2)
		CheckError(err)
		count1++
		//fmt.Println(departamento)
	}
	fmt.Println(count1)
	rows2, err := db1.Query(`
								SELECT *
								FROM agrupado
								`)
	CheckError(err)

	defer rows2.Close()
	count2 := 0
	lugaresUbigeo := make([]ubigeo, 0)
	for rows2.Next() {
		var lugar ubigeo
		err = rows2.Scan(&lugar.inei, &lugar.lugar)
		lugaresUbigeo = append(lugaresUbigeo, lugar)
		//var str1 = []rune(departamento)
		//var str2 = []rune("PERU")
		//rst := levenshtein(str1, str2)
		CheckError(err)
		count2++
		//fmt.Println(departamento)
	}
	fmt.Println(count2)
	count3 := 0
	for _, lugarUbigeo := range lugaresUbigeo {
		for _, lugarDeltaIT := range lugaresDelaIT {
			var str1 = []rune(lugarDeltaIT.lugarDeltaIT)
			var str2 = []rune(lugarUbigeo.lugar)
			rst := levenshtein(str1, str2)

			if rst < 2 {
				/*insertStatement := `
									UPDATE lugar
									SET idubigeo = $2
									WHERE id= $1`
				_, err = db1.Exec(insertStatement, lugarDeltaIT.id, lugarUbigeo.inei)
				if err != nil {
					panic(err)
				}*/
				count3++
				fmt.Println("Lugar de delta IT ", lugarDeltaIT.lugarDeltaIT, " es similar a ", lugarUbigeo.lugar, " con similitud: ", rst)
			}
		}
	}
	fmt.Println("Registros similares con similitud menor a 2: ", count3)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
