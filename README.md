# MaoGo

<div align="center">
  <img src="/logo.jpg" width="30%" alt="Logo Mao"><br>
  <a href="#"><img alt="MaoGo" src="https://img.shields.io/badge/MaoGo-green?colorA=%23ff0000&colorB=%23017e40&style=for-the-badge"></a><br>
  <a href="https://github.com/fckvania/MaoGo/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/fckvania/MaoGo?style=flat-square"></a>
  <a href="https://github.com/fckvania/MaoGo/network/members"><img alt="Forks" src="https://img.shields.io/github/forks/fckvania/MaoGo?style=flat-square"></a>
  <a href="https://github.com/fckvania/MaoGo/watchers"><img alt="Watchers" src="https://img.shields.io/github/watchers/fckvania/MaoGo?style=flat-square"></a>
</div>

<br><br>
<p>MaoGo adalah Bot WhatsApp yang dibuat menggunakan Golang dan package <a href="https://github.com/tulir/whatsmeow" target="_blank">whatsmeow</a>.</p><br>


## Cara Penggunaan Dan Penginstalan

1. **Langkah 1:** Unduh atau clone repositori ini.
3. **Langkah 2:** Instal Golang [disini](https://go.dev/doc/install).
4. **Langkah 3:** Ubah file `.env` dengan informasi yang diperlukan (seperti jid owner bot dan nama bot).
5. **Langkah 4:** Jalankan bot dengan perintah:
```shell
cd MaoGo
go run src/run.go
# atau
go build src/run.go
./run
```
7. **Langkah 5:** Buka WhatsApp dan scan QR yang muncul di log.

## Kontribusi

Jika Anda ingin berkontribusi pada pengembangan bot ini, berikut adalah langkah-langkah yang dapat Anda lakukan:
- Fork repositori ini.
- Buat branch baru: `git checkout -b fitur-baru`.
- Lakukan perubahan yang diperlukan.
- Commit perubahan Anda: `git commit -m 'Menambahkan fitur baru'`.
- Push ke branch yang Anda buat: `git push origin fitur-baru`.
- Buat pull request.

## Lisensi

[GPL-3.0 license](/LICENSE.txt)
