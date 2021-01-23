<?php
error_reporting(0);
 
$lokasi_file3 		= $_FILES['unggahfoto']['tmp_name'];
$nama_file3 		= $_FILES['unggahfoto']['name'];
$ukuran_foto3  		= $_FILES['unggahfoto']['size'];

$lokasi_file2 			= $_FILES['unggahsurattugas']['tmp_name'];
$nama_file2 			= $_FILES['unggahsurattugas']['name'];
$ukuran_foto2   		= $_FILES['unggahsurattugas']['size'];

$lokasi_file1 			= $_FILES['unggahsertifikat']['tmp_name'];
$nama_file1 			= $_FILES['unggahsertifikat']['name'];
$ukuran_foto1   		= $_FILES['unggahsertifikat']['size'];


$max = 5000000; //5Mb
$max2 = 1000000; //1Mb
$valid_ext = array('pdf');
$valid_ext2 = array('jpeg','png','jpg');
$ext1 = strtolower(end(explode('.', $nama_file1)));
$ext3 = strtolower(end(explode('.', $nama_file3)));
$ext2 = strtolower(end(explode('.', $nama_file2)));
        

 		if($ukuran_foto1 > $max || $ukuran_foto2 > $max || $ukuran_foto3 > $max2)
		{
			 echo "<script> alert ('Gagal.. Ukuran file terlalu besar Mak file PDF 5 MB dan Foto Mak 1 MB.'); history.back(); </script>";
		}else{
		if(in_array($ext1, $valid_ext) && in_array($ext2, $valid_ext) && in_array($ext3, $valid_ext2)){
			move_uploaded_file($_FILES['unggahfoto']['tmp_name'], "image/".$_FILES['unggahfoto']['name']);
            
            echo "<script> alert('Anda Telah Berhasil input data'); 
            window.location='http://localhost/webprestasi/web/dist/kelolaprestasikompetisi.html' </script>";
            
			}else{
					echo "<script> alert ('Data Gagal diinput.'); history.back(); </script>";
			}
        }	
?>