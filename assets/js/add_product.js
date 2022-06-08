let addProductData = document.getElementsByClassName("product_input");
if (addProductData.length > 0){
    addProductData[1].oninput = () => {
        if (addProductData[1].value < 0){
            addProductData[1].value = 1;
        }
    }
}
let addProductBtn = document.getElementById("product_btn");
if (addProductBtn){
    addProductBtn.onclick = () => {
        General("PUT", "/api/motionProduct/add", {
            Name: addProductData[0].value,
            Amount: +addProductData[1].value,
        }, (response) => {
            if (response){
                alert("Товар добавлен на склад");
            }
        })
    }
}
