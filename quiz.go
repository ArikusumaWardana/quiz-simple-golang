package main

import (
	"bufio"
	"fmt"
	"os"
)


type question struct {
	soal string
	option []string
	jawaban int
}

func tampilkanSoal(q question) {
	fmt.Println(q.soal)
	for i, option := range q.option {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func cekJawaban(jawabanUser int, JawabanBenar int) bool {
	return jawabanUser == JawabanBenar
}


func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Nama:")

	scanner.Scan()
	nama := scanner.Text()

	questions := []question {
		{
			soal: "Pada tahun berapa Perang Dunia I dimulai?",
			option: []string{"1914", "1920", "1939", "1945"},
			jawaban: 1,
		},
		{
			soal: "Siapakah tokoh sejarah yang terkenal sebagai penjelajah Portugis yang pertama kali mengelilingi Tanjung Harapan (Cape of Good Hope)?",
			option: []string{"Christopher Columbus", "Ferdinand Magellan", "Vasco da Gama", "Marco Polo"},
			jawaban: 3,
		},
		{
			soal: "Kaisar Romawi yang terkenal karena membuat \"Kode Hukum Romawi\" adalah...",
			option: []string{"Julius Caesar", "Augustus", "Nero", "Justinian I"},
			jawaban: 4,
		},
		{
			soal: "Pada tahun berapa Amerika Serikat merdeka dari Inggris?",
			option: []string{"1676", "1776", "1876", "1976"},
			jawaban: 2,
		},
		{
			soal: "Siapa pemimpin Perang Salib Pertama yang terkenal?",
			option: []string{"Richard the Lionheart", "Saladin", "Charlemagne", "Genghis Khan"},
			jawaban: 1,
		},
		{
			soal: "Apa nama peristiwa yang menandai berakhirnya Perang Dunia II di Eropa?",
			option: []string{"D-Day", "V-E Day", "Pearl Harbor Day", "MHiroshima Day"},
			jawaban: 2,
		},
		{
			soal: "Siapakah penulis \"The Communist Manifesto\" bersama Karl Marx?",
			option: []string{"Friedrich Engels", "Vladimir Lenin", "Leon Trotsky", "Joseph Stalin"},
			jawaban: 1,
		},
		{
			soal: "Perang Dingin terjadi antara blok politik apa?",
			option: []string{"NATO dan Pakta Warsawa", "Axis dan Sekutu", "Central Powers dan Allied Powers", "Entente dan Triple Alliance"},
			jawaban: 1,
		},
		{
			soal: "Di mana letak Zona Demiliterisasi Korea (DMZ) yang memisahkan Korea Utara dan Korea Selatan?",
			option: []string{"38th Parallel", "17th Parallel", "Equator", "Tropic of Cancer"},
			jawaban: 1,
		},
		{
			soal: "Apa peristiwa sejarah yang menandai akhir Kekaisaran Romawi Barat pada tahun 476 M?",
			option: []string{"Pertempuran Adrianople", "Invasi Viking", "Penaklukan Britania oleh Romawi", "Penyatuan Italia"},
			jawaban: 4	,
		},
	} 

	jawabanBenar := 0
	jawabanSalah := 0

	for _, question := range questions {
		tampilkanSoal(question)

		var jawabanUser int
		fmt.Print("Jawaban anda (1/2/3/4):")
		fmt.Scan(&jawabanUser)

		if cekJawaban(jawabanUser, question.jawaban) {
			fmt.Println("Jawaban anda benar!")
			jawabanBenar++
		} else {
			fmt.Println("Jawaban anda salah!")
			jawabanSalah++
		}
		fmt.Printf("\n")
	}

	score := jawabanBenar * 10;

	fmt.Printf("Nama Peserta: %s \n", nama)
	fmt.Printf("Total Score: %d \n", score)
	fmt.Printf("Total Jawaban Benar = %d \n", jawabanBenar)
	fmt.Printf("Total Jawaban Salah = %d", jawabanSalah)


}