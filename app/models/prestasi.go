package models

const (
	layoutDate = "2006-01-02"
	layoutYear = "2006"
)

//gorm:"primary_key";auto_increment;not_null
type Mahasiswa struct {
	Idmhs          int    `gorm:"primary_key";auto_increment;not_null json:"idmhs"`
	Nama           string `json:"nama"`
	Nim            string `json:"nim"`
	Password       string `json:"password"`
	Fakultas       string `json:"fakultas"`
	Programstudi   string `json:"programstudi"`
	Nohp           string `json:"nohp"`
	Email          string `json:"email"`
	Alamatktp      string `json:"alamatktp"`
	Alamatdomisili string `json:"alamatdomisili"`
	Hobi           string `json:"hobi"`
}

type Fakultasprodi struct {
	Kodeprodi    string `json:"kode_prodi"`
	Namafakultas string `json:"nama_fakultas"`
	Namaprodi    string `json:"nama_prodi"`
}

type Admin struct {
	Idadmin  int    `json:"id_admin"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Prestasinonkompetisi struct {
	Idprestasinon       int    `gorm:"primary_key";auto_increment;not_null json:"id_prestasi_non"`
	Idmhs               int    `json:"id_mhs"`
	Namakegiatan        string `json:"nama_kegiatan"`
	Namapenyelenggaraan string `json:"nama_penyelenggaraan"`
	Tempatkegiatan      string `json:"tempat_kegiatan"`
	Tanggalawal         string `json:"tanggal_awal"`
	Tanggalakhir        string `json:"tanggal_akhir"`
	Unggahsertifikat    string `json:"unggah_sertifikat"`
	Unggahsurattugas    string `json:"unggah_surat_tugas"`
	Unggahfoto          string `json:"unggah_foto"`
	Status              string `json:"status"`
}

type Prestasi struct {
	Idprestasi          int    `gorm:"primary_key";auto_increment;not_null json:"id_prestasi"`
	Idmhs               int    `json:"id_mhs"`
	Namakegiatan        string `json:"nama_kegiatan"`
	Namapenyelenggaraan string `json:"nama_penyelenggaraan"`
	URL                 string `json:"url"`
	Jumlah              int    `json:"jumlah"`
	Kategorikegiatan    string `json:"kategori_kegiatan"`
	Tingkatkegiatan     string `json:"tingkat_kegiatan"`
	Hasilkegiatan       string `json:"hasil_kegiatan"`
	Tempatkegiatan      string `json:"tempat_kegiatan"`
	Tanggalawal         string `json:"tanggal_awal"`
	Tanggalakhir        string `json:"tanggal_akhir"`
	Unggahsertifikat    string `json:"unggah_sertifikat"`
	Unggahsurattugas    string `json:"unggah_surat_tugas"`
	Unggahfoto          string `json:"unggah_foto"`
	Status              string `json:"status"`
}

type Foto struct {
	Username string `json:"username"`
	Value    string `json:"value"`
}

type Pengabdian struct {
	Idpengabdian  int    `gorm:"primary_key";auto_increment;not_null json:"id_pengabdian"`
	Idmhs         int    `json:"id_mhs"`
	Namaprogram   string `json:"nama_program"`
	Tahunkegiatan string `json:"tahun_kegiatan"`
	Unggahsurat   string `json:"unggah_surat"`
	Status        string `json:"status"`
}

type Organisasi struct {
	Idorganisasi   int    `gorm:"primary_key";auto_increment;not_null json:"id_organisasi"`
	Idmhs          int    `json:"id_mhs"`
	Namaorganisasi string `json:"nama_organisasi"`
	Jabatanmhs     string `json:"jabatan_mhs"`
	Periode        string `json:"periode"`
	Unggahsk       string `json:"unggahsk"`
	Status         string `json:"status"`
}
