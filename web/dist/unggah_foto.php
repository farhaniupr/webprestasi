<?php
error_reporting(0);
include './koneksi.php';
 
$id_prestasi 				= $_POST['id_prestasi'];
$lokasi_file3 		= $_FILES['unggah_foto']['tmp_name'];
$nama_file3 		= $_FILES['unggah_foto']['name'];
$ukuran_foto3  		= $_FILES['unggah_foto']['size'];

$max = 5000000; //5Mb
$max2 = 1000000; //1Mb
$valid_ext = array('pdf');
$valid_ext2 = array('jpeg','png','jpg');
$ext1 = strtolower(end(explode('.', $nama_file1)));
$ext2 = strtolower(end(explode('.', $nama_file2)));
$ext3 = strtolower(end(explode('.', $nama_file3)));
        
        
 if(!empty($lokasi_file1) && !empty($lokasi_file2) && !empty($lokasi_file3))
 {
		if($ukuran_foto > $max || $ukuran_foto2 > $max || $ukuran_foto3 > $max2 )
		{
			 echo "<script> alert ('Gagal.. Ukuran file terlalu besar Mak file PDF 5 MB dan Foto Mak 1 MB.'); history.back(); </script>";
		}else{
		if(in_array($ext1, $valid_ext) && in_array($ext2, $valid_ext) && in_array($ext3, $valid_ext2)){
			move_uploaded_file($_FILES['unggah_foto']['tmp_name'], "image/".$_FILES['unggah_foto']['name']);

		$update=mysqli_query($conn,"UPDATE prestasis SET
		unggah_foto='$nama_file3'
		WHERE idprestasi='$id_prestasi'");
		
			}else{
					echo "<script> alert ('Maaf... file yang anda pilih bukan file yang seharusnya.'); history.back(); </script>";
			}
             }
    }

	if($update) {
		echo "<script> alert('Anda Telah Berhasil input data');
		window.location = 'editprestasi.html'</script>";
	} else {
		echo "<script> alert ('Data  Gagal di input');
		window.location = 'editprestasi.html'</script>";
	}	
?>