<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<script>

</script>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
 <link rel="shortcut icon" href="#" type="image/x-icon">
    <title>Cetak Data Pengabdian Masyarakat</title>
        
</head>

<body>
<?php
include '../config/koneksi.php';

?>
<!doctype html>
<html>
  <head>
    <title>Data Pengabdian Masyarakat</title>
  </head>
  <body>
      
<div style=" width:1000px; margin: 0px auto;">
<center><h2> DATA PENGABDIAN MASYARAKAT </h2></center>
      <p>
      <table border="1">
            
          <th><center>No</th>
          <th><center>Nama Mahasiswa</th>
          <th><center>NIM</th>
          <th><center>Nama Program</th>
          <th><center>Tahun Kegiatan</th>
          <th><center>Surat Tugas</th>
          
          </thead>
                                        <tbody>
          <?php
                 $no=1;
                $data_pengabdian = mysqli_query($conn,"SELECT a.nama, a.nim, b.* from datapengabdian b
                LEFT JOIN datamahasiswa a ON b.idmhs = a.idmhs WHERE b.status='setuju' order by id_pengabdian  DESC");
                while($data=mysqli_fetch_array($data_pengabdian))
                {
                ?>

        <tr align="center">
        <td><center><?php echo $no++; ?></center></td>
                                   <td><?php echo $data['nama']; ?></td>
                                    <td><?php echo $data['nim']; ?></td>
                                    <td><?php echo $data['nama_program']; ?></td>
                                    <td><?php echo $data['tahun_kegiatan']; ?></td>
                                    
                                   
                                   
                                 
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