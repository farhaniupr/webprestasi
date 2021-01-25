## Instruksi
1. Download Golang for windows dan install
    ```https://golang.org/dl/go1.15.7.windows-amd64.msi```

2. Cara Check apabila golang sudah terinstall, buka CMD dan ketika 
    ```go version```

3. Download extension CORS Unblock dan enable extension
    ```https://chrome.google.com/webstore/detail/cors-unblock/lfhmikememgdcahcdlaciloancbhjino?hl=en```

4. Download repository ini secara manual >> klik tombol download(zip)  (tidak recomended)

   ```atau```

   clone repository pake git (download git) >>> ```https://git-scm.com/download/win```

   cara clone 
    ```bash
        git clone https://github.com farhaniupr/webprestasi.git
    ```

    bila ada update-an website yang dikerjakan tinggal ketik di folder webprestasi:
    ```bash
        git pull 
    ```
5. IMPORT terlebih dahulu prestasi.sql yang ada di repository karena struktur type data table BEDA

## Demo
6. Running  golang buka folder webprestasi hasil download / cloning
    ```bash
        go run main.go
    ```
    Contoh hasil running bisa dilihat di folder screenshoot

7. Copy folder web ke xampp dan running Website prestasi buka di browser 

    ```http://localhost/?<folder lain atau tidak>?/folder>/web/dist/index.html```

### Note
- file yang bisa diupload di tambah dan edit prestasi dan lain-lain >>> pdf dan image(hanya png) < 1MB 
- untuk import prodi tetap sama seperti sebelumnya memakai file excel