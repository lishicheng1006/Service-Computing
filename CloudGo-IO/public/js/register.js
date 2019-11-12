function reset() {
    $("#register-form").removeClass("error");
    $("#username").removeClass("error");
    $("#stuId").removeClass("error");
    $("#phone").removeClass("error");
    $("#email").removeClass("error");
    $("#errors").empty();
}

$(document).ready(function () {
    const DOMForm = document.getElementById("register-form");
    const regexStuId = /^[1-9]\d{7}$/;
    const regexUsername = /^[a-zA-Z]\w{5,17}$/;
    const regexPhone = /^[1-9]\d{10}$/;
    const regexEmail = /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/;

    DOMForm.addEventListener("reset", function (e) {
        reset();
    });

    DOMForm.addEventListener("submit", function (e) {
        e.preventDefault();
        reset();

        const form = $(DOMForm).serializeArray();
        const errors = ["", "", "", ""];

        if (!regexStuId.test(form[0].value)) errors[0] = "Student ID";
        if (!regexUsername.test(form[1].value)) errors[1] = "Username";
        if (!regexPhone.test(form[2].value)) errors[2] = "Phone";
        if (!regexEmail.test(form[3].value)) errors[3] = "Email";

        if (errors.filter((i) => i !== "").length) {
            const div = $(`<div class="list"></div>`);

            for (let i = 0; i < errors.length; i++) {
                if (errors[i] === "") continue;
                $(`#${form[i].name}`).addClass("error");
                div.append(`<li>${errors[i]} is invalid</li>`);
            }

            $("#errors").append(div);
            $("#register-form").addClass("error");
            return;
        }

        let requestData = {};
        for (let data of form) requestData[data.name] = data.value;

        fetch("/api/user/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(requestData)
        })
            .then((res) => res.json())
            .then((res) => {
                if (res.code === 200) {
                    window.location.href = `/api/user?stu_id=${requestData.stu_id}`;
                    return;
                }

                const err = $(`<div>Error: ${res.msg}</div>`)
                $("#errors").append(err);
                $("#register-form").addClass("error");
            })
            .catch((err) => alert(err));
    });
});