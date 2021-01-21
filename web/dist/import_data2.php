<?php
error_reporting(0);
include './koneksi.php';
require('spreadsheet-reader-master/php-excel-reader/excel_reader2.php');
require('spreadsheet-reader-master/SpreadsheetReader.php');

//upload data excel kedalam folder uploads
$target_dir = "file/".basename($_FILES['file']['name']);

move_uploaded_file($_FILES['file']['tmp_name'],$target_dir);

$Reader = new SpreadsheetReader($target_dir);

foreach ($Reader as $Key => $Row)
{
 // import data excel mulai baris ke-2 (karena ada header pada baris 1)
 if ($Key < 1) continue;   
 $query=mysqli_query($conn, "INSERT INTO fakultasprodi(kodeprodi,namafakultas,namaprodi) VALUES ('".$Row[2]."', '".$Row[0]."','".$Row[1]."')");
}

if ($query){
    echo "<script> alert ('Data Mahasiswa Berhasil Diimport');
    window.location='kelolaprodi.html'</script>";
}else
{
    echo "<script> alert ('Data Mahasiswa Gagal Diimport');
    window.location='kelolaprodi.html'</script>";
}
?>