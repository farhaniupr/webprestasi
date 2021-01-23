<?php
error_reporting(0);
include './koneksi.php';
 

 $id_prestasi 			= $_POST['idprestasi2'];
 $lokasi_file1 			= $_FILES['unggahsertifikat']['tmp_name'];
 $nama_file1 			= $_FILES['unggahsertifikat']['name'];
 $ukuran_foto1   		= $_FILES['unggahsertifikat']['size'];

		$max = 5000000; //5Mb
		$max2 = 1000000; //1Mb
		$valid_ext = array('pdf');
		$valid_ext2 = array('jpeg','png','jpg');
		$ext1 = strtolower(end(explode('.', $nama_file1)));
        
        
 if(!empty($lokasi_file1))
 {
		if($ukuran_foto1 > $max )
		{
			 echo "<script> alert ('Gagal.. Ukuran file terlalu besar Mak file PDF 5 MB dan Foto Mak 1 MB.'); history.back(); </script>";
		}else{
		if(in_array($ext1, $valid_ext)){
			move_uploaded_file($_FILES['unggahsertifikat']['tmp_name'], "image/".$_FILES['unggahsertifikat']['name']);

		$update=mysqli_query($conn,"UPDATE prestasis SET
		unggahsertifikat='$nama_file1'
		WHERE idprestasi='$id_prestasi'");
		
			}else{
		

					echo "<script> alert ('Maaf... file yang anda pilih bukan file yang seharusnya.'); history.back(); </script>";
			}
			 }}
			 if($update) {
				echo "<script> alert('Anda Telah Berhasil input data');
				window.location='http://localhost/webprestasi/web/dist/editprestasi.html?id=43' </script>";
			} else {
				echo "<script> alert ('Data  Gagal di input');
				window.location='http://localhost/webprestasi/web/dist/editprestasi.html?id=43' </script>";
			}	
            ?>