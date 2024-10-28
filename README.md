# IndoScript

Bahasa pemrograman interpreted dengan menggunakan sintak Bahasa Indonesia. Interpreternya ditulis dengan menggunakan bahasa GO.

## Contoh Penulisan

```js
fungsi statusSaatIni(umur){

    jika ( umur >= 40 ) {
        balikan "Tua";
    } lain jika (umur >= 30) {
        balikan "Dewasa";
    } lain {
        balikan "Muda";
    }

}

fungsi ambilMasukan(label){
    cetak(label);
    var nama = masukan();

    jika (nama == ""){
        balikan ambilMasukan(label);
    }

    balikan nama;
}

var nama = ambilMasukan("Nama Anda: ");
var lahir = keBilangan(ambilMasukan("Tahun Lahir: "));
var sekarang = keBilangan(ambilMasukan("Tahun Saat Ini: "));

var umurSaya = sekarang - lahir;

cetakBr("===============");
cetakBr("Nama Saya:");
cetakBr(nama);
cetakBr("Umur Saya:");
cetakBr(umurSaya);
cetakBr("Status:");
cetakBr(statusSaatIni(umurSaya));
```

## Dokumentasi

Untuk sementara, dokumentasi mengenai sintak dan referensi bahasa ada di [halaman Wiki](https://github.com/hsnfirdaus/indoscript/wiki) repositori ini.

## Menjalankan

Untuk mulai menjalankan IndoScript, silahkan unduh interpreter di [halaman rilis](https://github.com/hsnfirdaus/indoscript/releases) sesuai dengan sistem operasi anda.

Lalu simpan kode yang ingin anda jalankan dalam sebuah file (ekstensi yang disarankan .ids).

Jalankan interpreter:

```bash
./indoscript nama_file.ids
```

## Pengembangan

Untuk melakukan pengembangan interpreter ini di perangkat lokal anda, gunakan GO dan clone repositori ini.

Untuk menjalankan melalui GO:

```bash
go run . nama_file.ids
```

Didalam repositori ini terdapat beberapa folder:

- lekser: Melakukan _lexing_, yaitu mengubah dari teks biasa menjadi token-token.
- pengurai: Melakukan _parsing_, yaitu mengubah dari token-token menjadi node-node.
- penerjemah: Melakukan _interpreting_, yaitu menerjemahkan dan menjalankan node-node menjadi perintah komputer.
