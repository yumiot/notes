function userLogin() {
    let xmlhttp;
    let loginmail;
    let loginpassword;
    let usertishi;

    if (window.XMLHttpRequest) {
        xmlhttp = new XMLHttpRequest();
    }
    else {
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
            usertishi = xmlhttp.responseText;
            if (usertishi === "登录成功") {
                window.location = "/m/userown";
            }
            else {
                document.getElementById("myDiv1").innerHTML = usertishi;
            }
        }
    };

    loginmail = document.getElementById("mailbox").value;
    loginpassword = document.getElementById("password").value;

    xmlhttp.open("POST", "http://localhost:8080/login", true);
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xmlhttp.send("mailbox=" + loginmail + "&password=" + loginpassword);
}