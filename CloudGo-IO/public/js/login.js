$(document).ready(function () {
    const DOMForm = document.getElementById("register-form");

    DOMForm.addEventListener("submit", function (e) {
        e.preventDefault();
        const form = $(DOMForm).serializeArray();
        const stuId = form[0].value;
        window.location.href = `/api/user?stu_id=${stuId}`;
    });
});