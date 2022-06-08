let delProductData = document.getElementsByClassName("product_input");
if (delProductData.length > 0){
    delProductData[1].oninput = () => {
        if (delProductData[1].value < 0){
            delProductData[1].value = 1;
        }
    }
}
let delProductBtn = document.getElementById("product_btn");
if (delProductBtn){
    delProductBtn.onclick = () => {
        General("PUT", "/api/motionProduct/del", {
            Name: delProductData[0].value,
            Amount: +delProductData[1].value,
        }, (response) => {
            alert(response)
        })
    }
}
