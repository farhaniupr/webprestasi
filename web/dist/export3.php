<?php
// Fungsi header dengan mengirimkan raw data excel
header("Content-type: application/vnd-ms-excel");
 

header("Content-Disposition: attachment; filename=Data Prestasi Non Kompetisi.xls");
 
// Tambahkan table
include 'print3.php';
?>