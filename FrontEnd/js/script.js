class productInfo {
    constructor(productId, productTypeId, productName, productPrice){
        this.productId = productId,
        this.productTypeId = productTypeId,
        this.productName = productName,
        this.productPrice = productPrice
    }
}

// SLIDESHOW 
var imgArray = [
    'img/image1.jpg',
    'img/image2.jpg',
    'img/image3.jpg'],
    curIndex = 0;
    imgDuration = 5000;

function slideShow() {
    document.getElementById('image1').src = imgArray[curIndex];
    curIndex++;
    if (curIndex == imgArray.length) { curIndex = 0; }
    setTimeout("slideShow()", imgDuration);
}
// slideShow();

// FORM VALIDATION 
function validateForm() {
  var name = document.myForm.name.value;
  var email = document.myForm.email.value;

  if (name == null || name == ""){
      alert("Name can't be blank!");
      return false;
  } 
}

function validateEmail(){
  var x = document.myForm.email.value;
  var atposition = x.indexOf("@");
  var dotposition = x.indexOf(".");

  if (atposition<1 || dotposition<atposition+2 || dotposition+2>=x.length){
      alert("Please enter a valid email address!");
      return false;
  }
}

//  loads products and slideshow
window.onload = function(){
  displayProducts();
  slideShow();
}

// function to display products
function displayProducts() {
  fetch('http://localhost:8080/createproduct')
  .then(res => {
      return res.json()
  })
  .then(prod =>{
      
      for( let i=0; i < prod.length; i++){
        var product = `<figure class="individual_row" id="${prod[i].product_typeid}">
                          <a href="productdetail.html">
                              <img src="img/image${i}.jpg" alt="${prod[i].product_name}" />
                              <figcaption>${prod[i].product_name}<br/>${prod[i].product_price}</figcaption>
                          </a>
                       </figure>
        `
        document.getElementById("arow").innerHTML += product
      }
      console.log(prod);
  }).catch(error => {
      console.log(error);
  })
}

// function to filter products
function filterProducts(productTypeId) {
  let rows = document.getElementById('arow').children;
  if (productTypeId == "-1") {
    for (let z = 0;z<rows.length;z++) {
      rows[z].setAttribute('style',"width: 15%, display: inline-block")
      }
    } else {
    for (let z = 0;z<rows.length;z++) {
      if(rows[z].id != productTypeId) {
        rows[z].setAttribute('style', "display: none")
      } else {
        rows[z].setAttribute('style',"width: 15%, display: inline-block")
      }
    }
  }
}

// nav bar
function topNav() {
  var x = document.getElementById("myTopnav");
  if (x.className === "topnav") {
    x.className += " responsive";
  } else {
    x.className = "topnav";
  }
}
