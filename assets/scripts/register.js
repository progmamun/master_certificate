$(document).ready(function () {
    // mdc ripple effect
    const textFieldElements = [].slice.call(document.querySelectorAll('.mdc-text-field'));
    textFieldElements.forEach((textFieldEl) => {
        mdc.textField.MDCTextField.attachTo(textFieldEl);
    });


    $('#regForm').submit(function () {
        // checking if any field is empty(only space) or not
        let flag = 1;
        if ($('#username').val().trim() == "") {
            $('#username').val("");
            alert("Please give a username!");
            flag = 0;
        }
        if (flag == 1 && ($('#password').val() != $('#confirmPassword').val())) {
            alert("Password didn't match!");
            flag = 0;
        }

        if (flag) {
            console.log($('#regForm').serialize())
            //displaying loading gif
            $('#loadingGif').css("display", "block");
            //sending ajax post request
            let request = $.ajax({
                async: true,
                type: "POST",
                url: "/register",
                data: $('#regForm').serialize(),
                // error: function (err, statusCode) {
                //     alert(err, statusCode);
                // }
            });
            request.done(function (response) {
                //hiding loading gif
                $('#loadingGif').css("display", "none");

                if (response.trim() == "username") {
                    alert("Username already taken. Please choose a different one.");
                } else if (response.trim() == "email") {
                    alert("Email already exsist. Please choose a different one.");
                } else if (response.trim() == "Registration Done") {
                    alert("Registration successful. Email verification link was sent to your provided email.");
                    window.location.href = "/login";
                    //resetting form field
                    $('#username').val("");
                    $('#email').val("");
                    $('#password').val("");
                    $('#confirmPassword').val("");
                } else {
                    alert("Registration unsuccessful. Something went wrong. Please try again!");
                }
            });
            request.fail(function (response) {
                //hiding loading gif
                $('#loadingGif').css("display", "none");

                alert(response);
            });
        }
        return false;
    });
});