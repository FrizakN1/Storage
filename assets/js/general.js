function General(method, uri, data, callback){
    let xhr = new XMLHttpRequest();
    xhr.open(method, uri);
    xhr.onload = function() {
        if (typeof callback === "function"){
            callback(JSON.parse(this.response))
        }
    }
    if (data){
        xhr.send(JSON.stringify(data))
    } else {
        xhr.send()
    }
}

let hamburger = document.getElementById("hamburger");
let close = document.getElementById("close");
let menu = document.getElementById("menu");
if (hamburger){
    hamburger.onclick = () => {
        menu.classList.add("see");
    }
}

if (close){
    close.onclick = () => {
        menu.classList.remove("see");
    }
}