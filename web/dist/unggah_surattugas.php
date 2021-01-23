<?php
error_reporting(0);
include './koneksi.php';
 

 $id_prestasi 			= $_POST['idprestasi'];

 $lokasi_file2 			= $_FILES['unggahsurattugas']['tmp_name'];
 $nama_file2 			= $_FILES['unggahsurattugas']['name'];
 $ukuran_foto2   		= $_FILES['unggahsurattugas']['size'];


		$max = 5000000; //5Mb
		$max2 = 1000000; //1Mb
		$valid_ext = array('pdf');
		$valid_ext2 = array('jpeg','png','jpg');
		$ext2 = strtolower(end(explode('.', $nama_file2)));
    
        
        
 if(!empty($lokasi_file2))
 {
		if($ukuran_foto2 > $max)
		{
			 echo "<script> alert ('Gagal.. Ukuran file terlalu besar Mak file PDF 5 MB dan Foto Mak 1 MB.'); history.back(); </script>";
		}else{
		if(in_array($ext2, $valid_ext)){
			move_uploaded_file($_FILES['unggahsurattugas']['tmp_name'], "image/".$_FILES['unggahsurattugas']['name']);

		$update=mysqli_query($conn,"UPDATE prestasis SET
		unggahsurattugas='$nama_file2'
		WHERE idprestasi='$id_prestasi'");
			}else{
					echo "<script> alert ('Maaf... file yang anda pilih bukan file yang seharusnya.'); history.back(); </script>";
			}
             }
}
if($update) {
	echo "<script> alert('Anda Telah Berhasil input data');
	window.location='http://localhost/webprestasi/web/dist/editprestasi.html?id=43' </script>";
} else {
	echo "<script> alert ('Data  Gagal di input');
	window.location='http://localhost/webprestasi/web/dist/editprestasi.html?id=43' </script>";
}	
?>