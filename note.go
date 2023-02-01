package belajargolangcontext

/**
- Pengenalan Context
        1. Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline
        2. Context biasanya dibuat per request(misal setiap ada request masuk ke server web melalui http request)
        3. Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses

- Kenapa Context perlu dipelajari?
    1. context di golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain
    2. dengan menggunakan context, ketika kita ingin membatalkan semua proses, kita cukup mengirim sinyal ke context, maka secara otomatis semua proses akan dibatalkan
    3. hampir semua bagian di golang memanfaatkan context, seperti database, http server, http client, dan lain-lain
    4. bahkan di google sendiri, ketika menggunakan golang, context wajib digunakan dan selalu dikirim ke setiap function yang dikirim

- package context
    1. context direpresentasikan di dalam sebuah interface Context
    2. interface Context terdapat dalam package context
    3. https://golang.org/pkg/context/

- Membuat Context
    1. Karena Context adalah sebuah interface, untuk membuat context kita butuh sebuah struct yang sesuai dengan kontrak interface Context
    2. Namun kita tidak perlu membuatnya secara manual
    3. di golang package context terdapat function yang bisa kita gunakan untuk membuat Context

- Function Membuat Context
    1. context.Background() : Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun. Biasanya digunakan di Main function atau dalam unit test, atau dalam awal proses request terjadi.
    2. context.TODO() : Membuat context kosong seperti Background(), namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan.

- Parent dan Child Context
    1. Context menganut konsep parent dan child
    2. artinya, saat kita membuat context, kita bisa membuat child context dari context yang sudah ada
    3. parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context
    4. konsep ini mirip dengan pewarisan di pemrograman berorientasi object

- hubungan antara Parent dan Child Context
    1. Parent dan Child context akan selalu terhubung
    2. saat nanti kita melakukan, misal : pembatalan context A, maka semua child dan sub child dari context A akan ikut dibatalkan
    3. namun jika misal kita membatalkan context B, hanya context B dan semua child dan sub child nya yang dibatalkan, parent context B tidak akan ikut dibatalkan
    4. begitu juga nanti saat kita menyisipkan data ke dalam context A, semua child dan sub child nya bisa mendapatkan data tersebut
    5. namun jika kita menyisipkan data di context B, hanya context B dan semua child dan sub child nya yang mendapatkan data, parent B tidak akan mendapatkan data

- Immutable
    1. Context merupakan object yang immutable, artinya setelah context dibuat, dia tidak bisa diubah lagi
    2. ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dan yang lainnya, secara otomatis akan membentuk child context baru, bukan merubah context tersebut

- cara membuat child context
    1. cara membuat child context ada banyak caranya, yang akan kita bahas di materi-materi selanjutnya.

- Context With Value
    1. pada saat awal membuat context, context tidak memiliki value
    2. kita bisa menambahkan sebuah value dengan data Pair(key - value) ke dalam context
    3. saat kita menambahkan value ke context, secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah sama sekali
    4. untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)

- Context With Cancel
    1. selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
    2. kapan sinyal cancel diperlukan dalam context?
    3. biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
    4. biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke context nya
    5. namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap context nya, jika tidak, tidak ada gunanya
    6. untuk membuat context dengan cancel signal, kita bisa menggunakan function context.WithCancel(parent)

- Context With Timeout
    1. selain menambahkan value ke contet, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis dengan menggunakan pengaturan timeout
    2. dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
    3. penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api, namun ingin menentukan batas maksimal timeout nya
    4. untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan function context.WithTimeout(parent, duration)

- Context With Deadline
    1. selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline
    2. pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang, kalau deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini
    3. untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, kita bisa menggunakan function context.WithDeadline(parent, time)
*/