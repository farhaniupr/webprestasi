<?php 
//file ini merupakan file login untuk mahasiswa
session_start();
include 'config/koneksi.php';

//cek login
if (isset($_POST['login'])) { //apabila tombol login dijalankan

//mendapatkan input dari pengguna
$username=$_POST['username'];
$password=$_POST['password'];


//mencek di database tabel admin adakah username dan passwordnya 
$login="SELECT * FROM admin WHERE username='$username'";
$cek=mysqli_query($conn,$login);
$ketemu=mysqli_num_rows($cek);



$login2="SELECT * FROM admin WHERE password='$password'";
$cek2=mysqli_query($conn,$login2);
$ketemu2=mysqli_num_rows($cek2);

$login3="SELECT * FROM admin WHERE username='$username' AND password='$password'";
$cek3=mysqli_query($conn,$login3);
$ketemu3=mysqli_num_rows($cek3);


//mencek di database tabel admin adakah username dan passwordnya 
$login4="SELECT * FROM datamahasiswa WHERE nim='$username'";
$cek4=mysqli_query($conn,$login4);
$ketemu4=mysqli_num_rows($cek4);

$login5="SELECT * FROM datamahasiswa WHERE password='$password'";
$cek5=mysqli_query($conn,$login5);
$ketemu5=mysqli_num_rows($cek5);

$login6="SELECT * FROM datamahasiswa WHERE nim='$username' AND password='$password'";
$cek6=mysqli_query($conn,$login6);
$ketemu3=mysqli_num_rows($cek6);

$r=mysqli_fetch_array($cek3);
$r2=mysqli_fetch_array($cek6);

if($username == 'admin'){
// menampilkan pesan gagal jika belum memasukan username dan password
if ($ketemu == 0 & $ketemu2 == 0){
echo "<script> alert ('Username dan Password anda tidak valid ! Mohon periksa kembali.');
window.location = 'loginp.php'</script>";
}
// menampilkan pesan gagal jika username salah
else if ($ketemu == 0){
echo "<script> alert ('Username anda tidak valid ! Mohon periksa kembali.');
window.location = 'loginp.php'</script>";
}

// menampilkan pesan gagal jika password salah
else if ($ketemu2 == 0){
echo "<script> alert ('Password anda tidak valid ! Mohon periksa kembali.');
window.location = 'loginp.php'</script>";
}

//jika ketemu dibuat sesi loginnya
else {
//detail sesi loginnnya
  $_SESSION['id_admin'] = $r['id_admin'];
  $_SESSION['username'] = $r['username'];
  $_SESSION['password'] = $r['password'];
  
  

                            
									
//login berhasil
echo "<script> alert ('Login Berhasil');
window.location = 'index.html'</script>";
}
}else{
// menampilkan pesan gagal jika belum memasukan username dan password
if ($ketemu4 == 0 & $ketemu5 == 0){
	echo "<script> alert ('Username dan Password anda tidak valid ! Mohon periksa kembali.');
	window.location = 'login1.php'</script>";
	}
	// menampilkan pesan gagal jika username salah
	else if ($ketemu4 == 0){
	echo "<script> alert ('Username anda tidak valid ! Mohon periksa kembali.');
	window.location = 'login1.php'</script>";
	}
	
	// menampilkan pesan gagal jika password salah
	else if ($ketemu5 == 0){
	echo "<script> alert ('Password anda tidak valid ! Mohon periksa kembali.');
	window.location = 'login1.php'</script>";
	}
	
	//jika ketemu dibuat sesi loginnya
	else {
	//detail sesi loginnnya
	  $_SESSION['idmhs'] = $r2['idmhs'];
	  $_SESSION['nim'] = $r2['nim'];
	  $_SESSION['password'] = $r2['password'];
	  
	  
	
								
										
	//login berhasil
	echo "<script> alert ('Login Berhasil');
	window.location = 'mhs/index.php'</script>";
	}
	}
}	
?>

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sign in - Voler Admin Dashboard</title>
    <link rel="stylesheet" href="assets/css/bootstrap.css">
    
    <link rel="shortcut icon" href="assets/images/logoupr.png" type="image/x-icon">
    <link rel="stylesheet" href="assets/css/app.css">
</head>

<body>
    <div id="auth">
        
<div class="container">
    <div class="row">
        <div class="col-md-5 col-sm-12 mx-auto">
            <div class="card pt-4">
                <div class="card-body">
                    <div class="text-center mb-5">
                        <img src="assets/images/logoside.png" height="100" class='mb-4'>
                        <h3>Sign In</h3>
                        <p>Please sign in to continue.</p>
                    </div>
                    <form action="login.php" method="post">
                        <div class="form-group position-relative has-icon-left">
                            <label for="username">Username</label>
                            <div class="position-relative">
                                <input type="text" class="form-control" id="username">
                                <div class="form-control-icon">
                                    <i data-feather="user"></i>
                                </div>
                            </div>
                        </div>
                        <div class="form-group position-relative has-icon-left">
                            <div class="clearfix">
                                <label for="password">Password</label>
                                <a href="auth-forgot-password.html" class='float-right'>
                                    <small>Forgot password?</small>
                                </a>
                            </div>
                            <div class="position-relative">
                                <input type="password" class="form-control" id="password">
                                <div class="form-control-icon">
                                    <i data-feather="lock"></i>
                                </div>
                            </div>
                        </div>

                        <div class='form-check clearfix my-4'>
                            <div class="checkbox float-left">
                                
                            </div>
                            <div class="float-right">
                                
                            </div>
                        </div>
                        <div class="clearfix">
                            <button class="btn btn-primary float-right">Login</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

    </div>
    <script src="assets/js/feather-icons/feather.min.js"></script>
    <script src="assets/js/app.js"></script>    
    <script src="assets/js/main.js"></script>
</body>

</html>
