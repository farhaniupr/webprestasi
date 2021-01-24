<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<script>

</script>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
 <link rel="shortcut icon" href="#" type="image/x-icon">
    <title>Cetak Data Prestasi Non Kompetisi</title>
        
</head>

<body>
<?php
include './koneksi.php';

?>
<!doctype html>
<html>
  <head>
    <title>Data Prestasi Non Kompetisi</title>
  </head>
  <body>
      
<div style=" width:1000px; margin: 0px auto;">
<center><h2> DATA PRESTASI NON KOMPETISI </h2></center>
      <p>
      <table border="1">
            
          <th><center>No</th>
          <th><center>Nama Mahasiswa</th>
          <th><center>NIM</th>
          <th><center>Nama Kegiatan</th>
          <th><center>Nama Penyelenggara</th>
          <th><center>Tempat Kegiatan</th>
          <th><center>Tanggal Awal</th>
          <th><center>Tanggal Akhir</th>
          
          </thead>
                                        <tbody>
          <?php
                $no=1;
                $data_prestasinonkompetisi = mysqli_query($conn,"SELECT a.nama, a.nim, b.* from prestasinonkompetisis b
                LEFT JOIN mahasiswas a ON b.idmhs = a.idmhs WHERE b.status='setuju' order by b.idprestasinon DESC");
                while($data=mysqli_fetch_array($data_prestasinonkompetisi))
                {
                ?>

        <tr align="center">
        <td><center><?php echo $no++; ?></center></td>
                                   <td><?php echo $data['nama']; ?></td>
                                    <td><?php echo $data['nim']; ?></td>
                                    <td><?php echo $data['namakegiatan']; ?></td>
                                    <td><?php echo $data['namapenyelenggaraan']; ?></td>
                                    <td><?php echo $data['tempatkegiatan']; ?></td>
                                    <td><?php echo $data['tanggalawal']; ?></td>
                                    <td><?php echo $data['tanggalakhir']; ?></td>
                                   
                                   
                                 
        </tr>
                
              <?php } ?>
                                </tbody>
      </table>
            
<br />
<br />
<br />
<br />
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
</body>
</html>