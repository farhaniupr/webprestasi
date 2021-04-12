<?php
// Fungsi header dengan mengirimkan raw data excel
header("Content-type: application/vnd-ms-excel");
 

header("Content-Disposition: attachment; filename=Data Mahasiswa.xls");
 
// Tambahkan table
include 'print5.php';
?>