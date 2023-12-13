package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Menerima Input
	var angkaDepan, angkaBelakang int	
	var inputSerangan, textOutput string
	fmt.Println(`Hanya ada satu cara untuk mengalahkan pasukan ini, yaitu apabila kamu menancapkan pedang kamu tepat di tengah-tengah jantung pasukan ini.
	Serangan kamu pas di jantung musuh apabila: 
	'Angka Pertama' = 'Genap' dan 'Angka Terakhir' = 'Ganjil'.`)
	fmt.Print("Serangan Anda (Angka 2 Bilangan) : ")
	fmt.Scanln(&inputSerangan)	

	serangan, errCheck := strconv.Atoi(inputSerangan)
	if errCheck != nil {
		textOutput = "Mohon untuk hanya menginput angka 2 bilangan positif"
	} else {
		//Perhitungan dan pengecekan input
		if serangan > 100 {
			//Angka1 & Angka2 Lebih Dari 2 Digit Atau Negatif
			textOutput = "Angka Tidak Valid!"
		} else {
			if strings.Contains(string(inputSerangan[0]),"0") && serangan < 10 {
				angkaBelakang = serangan
				angkaDepan = 0
			} else  {
				angkaDepan = serangan/10
				angkaBelakang = serangan - angkaDepan*10
			}	
			if angkaDepan%2 == 0 {
				if angkaBelakang%2 == 1 {
					//Angka1 Genap & Angka2 Ganjil
					textOutput = "One Zombie Down!"
				} else if angkaBelakang%2 == 0 {
					//Angka1 Genap & Angka2 Genap
					textOutput = "Try Again To Attack!"
				}
			} else if angkaDepan%2 == 1 {
				if angkaBelakang%2 == 1 {
					//Angka1 Ganjil & Angka2 Ganjil
					textOutput = "Try Again To Attack!"
				} else if angkaBelakang%2 == 0 {
					//Angka1 Ganjil & Angka2 Genap
					textOutput = "You Are Dead!"
				}
			}
		}
		//fmt.Printf("%d%d\n", angkaDepan, angkaBelakang)		
	}
	//Output
	fmt.Printf("%s", textOutput)	
}