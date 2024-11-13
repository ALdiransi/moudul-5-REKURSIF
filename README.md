# moudul-5-REKURSIF
1.Kode ini menggunakan pendekatan memoization untuk menghitung nilai Fibonacci secara efisien. Dengan mendeklarasikan peta memo secara global, hasil perhitungan Fibonacci yang sudah dilakukan dapat disimpan dan digunakan kembali,
sehingga mempercepat proses kalkulasi pada pemanggilan berikutnya. Fungsi fibonacci bekerja secara rekursif: pertama-tama memeriksa apakah nilai n sudah ada di dalam memo; jika ada, fungsi langsung mengembalikannya. Jika belum,
dan jika n lebih kecil atau sama dengan 1, maka n dikembalikan sebagai kasus dasar. Untuk nilai lainnya, fungsi menghitung fibonacci(n-1) + fibonacci(n-2) dan menyimpan hasilnya di memo, kemudian mengembalikan nilai tersebut. 
Pendekatan ini mengurangi waktu komputasi untuk nilai n yang besar

2.Kode ini menggunakan rekursi untuk mencetak pola bintang segitiga ke bawah. Fungsi printStars dipanggil secara rekursif hingga n bernilai 0, yang bertindak sebagai kondisi penghentian (base case). Setiap kali n lebih dari 0,
printStars akan memanggil dirinya sendiri dengan nilai n-1, sehingga menghasilkan baris dengan jumlah bintang yang meningkat dari atas ke bawah.

3.Kode ini menggunakan fungsi rekursif findFactors untuk mencari dan mencetak semua faktor dari sebuah bilangan num yang dimasukkan pengguna. Fungsi ini menerima dua parameter, num sebagai bilangan yang akan dicari faktornya,
dan divisor sebagai pembagi yang dimulai dari 1. Jika divisor melebihi num, fungsi berhenti (base case). Untuk setiap nilai divisor yang habis membagi num, nilai tersebut dicetak sebagai faktor. Fungsi findFactors kemudian 
dipanggil kembali dengan divisor ditambah 1, sehingga setiap pembagi diperiksa hingga mencapai num.

4.Kode ini menggunakan fungsi rekursif printSequence untuk mencetak barisan bilangan dari n ke 1 secara menurun, kemudian mencetaknya kembali secara naik. Fungsi printSequence menerima satu parameter, n, yang mewakili 
bilangan tertinggi dalam urutan. Jika n sama dengan 0, fungsi berhenti (base case). Saat n lebih dari 0, fungsi mencetak nilai n, kemudian memanggil dirinya sendiri dengan n-1, menciptakan bagian penurunan

5.Kode ini menggunakan fungsi rekursif printOddNumbers untuk mencetak bilangan ganjil dari 1 hingga n secara urut. Fungsi printOddNumbers menerima parameter n sebagai batas atas bilangan yang akan dicetak. Jika n kurang dari 1,
fungsi berhenti (base case). Untuk nilai n yang valid, fungsi memanggil dirinya sendiri dengan n - 2 terlebih dahulu untuk mencetak bilangan ganjil dari yang terkecil, lalu mencetak nilai n setelah rekursi selesai, memastikan 
urutan naik dari 1 ke n.


6.Kode ini menggunakan fungsi rekursif pangkat untuk menghitung hasil perpangkatan x dengan eksponen y. Fungsi pangkat menerima dua parameter, x sebagai bilangan pokok dan y sebagai eksponen. Jika y sama dengan 0,
fungsi mengembalikan 1 karena bilangan apapun berpangkat nol hasilnya adalah 1 (base case). Jika y positif, fungsi mengembalikan hasil perkalian x dengan pangkat(x, y-1), sehingga nilai x dikalikan berulang hingga y kali.
Jika y negatif, fungsi mengembalikan hasil kebalikan dari pangkat(x, -y), yang menghitung nilai pangkat untuk eksponen positif dan membaliknya sesuai aturan pangkat negatif. Di fungsi main,
