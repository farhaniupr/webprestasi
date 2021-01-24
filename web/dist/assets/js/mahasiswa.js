// create a new mhs
const AddMahasiswa = (mhs) => {
    axios.post('http://localhost:8085/addmahasiswa', mhs)
        .then(response => {
            const addmhs = response.data;
            console.log(`POST: mhs is added`, addmhs);
            
        })
        .catch(error => console.error(error));
};

// event listener for form submission
const form = document.querySelector('form');

const formEvent = form.addEventListener('submit', event => {
    event.preventDefault();

    const nama = document.querySelector('#nama').value;
    const nim = document.querySelector('#nim').value;
    const fakultas = document.querySelector('#fakultas').value;
    const programstudi = document.querySelector('#programstudi').value;
    const nohp = document.querySelector('#nohp').value;
    const email = document.querySelector('#email').value;
    const password = document.querySelector('#password').value;
    const alamatktp = document.querySelector('#alamatktp').value;
    const alamatdomisili = document.querySelector('#alamatdomisili').value;
    const hobi = document.querySelector('#hobi').value;
    const mhs = { 
        nama : nama, 
        nim : nim, 
        fakultas : fakultas, 
        programstudi : programstudi, 
        nohp : nohp, 
        email : email,
        password : password,
        alamatktp : alamatktp,
        alamatdomisili : alamatdomisili,
        hobi : hobi 
    };
    AddMahasiswa(mhs);
});


