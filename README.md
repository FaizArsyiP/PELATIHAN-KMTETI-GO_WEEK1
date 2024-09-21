<h1> Tugas 1 Pelatihan WEBDEV 1 KMTETI </h1>
<b> Faiz Arsyi Pragata </b>

membuat program untuk mengonversi suhu dari celcius ke suhu lainnya dengan memanfaatkan switch-case untuk menentukan suhu tujuan untuk di konversi.

<h3>Alur Program</h3>
<h5>1. Input suhu awal</h5>

    fmt.Println("Masukkan suhu awal (dalam derajat Celcius):")
    fmt.Scanln(&suhuAwal)
    fmt.Println("Suhu awal:", suhuAwal, "°C")
    
<p>User akan diminta untuk memasukkan nilai awal dari suhu yang ingin dikonversi. Hasil dari input tersebut akan dimasukkan kedalam variabel suhuAwal</p>

<h5>2. Memilih suhu akhir</h5>

    fmt.Println("Pilih konversi suhu:")
    fmt.Println("1. Reamur")
    fmt.Println("2. Fahrenheit")
    fmt.Println("3. Kelvin")
    fmt.Scanln(&pilihan)

  <p>User akan memilih satuan suhu akhir dengan memasukkan angka 1-3, kemudian angka tersebut akan dimasukkan kedalam variabel pilihan</p>
  
<h5>3. Program akan menampilkan hasil dari konversi suhu sesuai dengan keinginan user</h5>

    switch pilihan {
      case 1:
		  fmt.Println("Konversi ke Reamur :", suhuAwal*0.8, "°R")
	  case 2:
		  fmt.Println("Konversi ke Fahrenheit :", suhuAwal*1.8+32, "°F") 
	  case 3:
		  fmt.Println("Konversi ke Kelvin :", suhuAwal+273.15, " K")
	  default:
		  fmt.Println("Pilihan tidak valid")
	  }

<p>Dengan menggunakan metode switch-case, program akan dijalankan sesuai dengan pilihan dari user.
  <ul> 
    <li>Jika user memasukkan angka 1, maka program akan mengubah suhu awal kedalam bentuk Reamur</li>
    <li>Jika user memasukkan angka 2, maka program akan mengubah suhu awal kedalam bentuk Farenheit</li>
    <li>Jika user memasukkan angka 3, maka program akan mengubah suhu awal kedalam bentuk Kelvin</li>
    <li>Jika selain angka tersebut maka akan dikeluarkan teks "pilihan tidak valid"</li>
  </ul>
</p>
