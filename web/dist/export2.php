<?php
// Fungsi header dengan mengirimkan raw data excel
header("Content-type: application/vnd-ms-excel");
 

header("Content-Disposition: attachment; filename=Data Organisasi.xls");
 
// Tambahkan table
include 'print2.php';
?>