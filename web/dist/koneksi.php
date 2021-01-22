<?php
$database ="prestasi";
$servername = "localhost";
$username = "root";
$password = "";

$conn = mysqli_connect($servername, $username, $password, $database);
// mengecek koneksi
if (!$conn) {
    die("Koneksi gagal: " . mysqli_connect_error());
}


function tgl_baru($tgl){
    $tanggal = substr($tgl,8,2);
    $bulan = substr($tgl,5,2);
    $tahun = substr($tgl,0,4);
    return $tanggal.'-'.$bulan.'-'.$tahun;      
}

function thn_baru($tgl){
    
    $tahun = substr($tgl,0,4);
    return $tahun;      
}


?>