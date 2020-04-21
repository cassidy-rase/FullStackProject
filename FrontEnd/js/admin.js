class userInfo {
    constructor(userId, firstName, lastName, email){
        this.userId = userId,
        this.firstName = firstName,
        this.lastName = lastName,
        this.email = email
    }
}

function viewUser(){

var user = new userInfo

fetch('http://localhost:8080')
.then(res => {
    return res.json()
}).then(user =>{
    console.log(user)
})
.catch(error => {
    console.log(error)
})
}

function createDb(){

    var user = new userInfo

    var firstName = document.getElementById('firstName').value;
    var lastName = document.getElementById('lastName').value;
    var email = document.getElementById('email').value;

    fetch('http://localhost:8080', {
        method: "post",
        body: JSON.stringify({first_name:firstName, last_name:lastName, email:email})
    }).then(res => {
        return res.json()
    }).then(user => {
        console.log(user)
    })
}



// function createDb(){

//     var user = new userInfo

//     var firstName = document.getElementById('firstName').value;
//     var lastName = document.getElementById('lastName').value;
//     var email = document.getElementById('email').value;

//     fetch('http://localhost:8080', {
//         method: "post",
//         body: JSON.stringify({first_name:firstName, last_name:lastName, email:email})
//     }).then(res => {
//         return res.json()
//     }).then(user => {
//         console.log(user)
//     })
// }

// function updateDb() {

//     var firstName = document.getElementById('firstName').value;
//     var lastName = document.getElementById('lastName').value;
//     var email = document.getElementById('email').value;

//     fetch('http://localhost/update', {
//         method: "post",
//         body: JSON.stringify({first_name:firstName, last_name:lastName, email:email})
//     }).then(res => {
//         return res.json()
//     }).then(user => {
//         console.log(user)
//     })
// }

// CONTACT FORM 
function createUser(){

    var user = new userInfo
    
    fetch('http://localhost:8080')
    .then(res => {
        return res.json()
    }).then(user =>{
        console.log(user)
    })
    .catch(error => {
        console.log(error)
    })
    }