
/* Registration Code */
$("#register_user").submit(function(e){
    e.preventDefault();
    var username = $("#username").val();
    var password = $("#password").val();
    var email = $("#email").val();
    var mobileNumber = $("#mobileNumber").val();
    var registerData = '{"username":"'+username+'","password":"'+password+'","email":"'+email+'", "mobileNumber":'+mobileNumber+'}'
    console.log("registerData : ", registerData);
    $.ajax({	
        type: "POST",
        url: "registration",
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'application/json'
        },
        data:registerData,
        success: function (response) {
            console.log("Success response is ", response)
       },
       error: (function(response) {
            console.log("error response is ", response)
       })
   });
});


/* Login Code */
$("#login_user").submit(function(e){
    e.preventDefault();
    var email = $("#emailMobile").val();
    var password = $("#password").val();
    var loginData = '{"emailMobile":"'+email+'","password":"'+password+'"}'
    console.log("loginData : ", loginData);
    $.ajax({	
        type: "POST",
        url: "login",
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'application/json'
        },
        data:loginData,
        success: function (response) {
            console.log("Success response is ", response)
       },
       error: (function(response) {
            console.log("error response is ", response)
       })
   });
});

/* Forgot Password Code */
$("#forgot_user").submit(function(e){
    e.preventDefault();
    var mobileNumber = $("#mobileNumber").val();
    var ForgotData = '{"mobileNumber":"'+mobileNumber+'"}'
    console.log("ForgotData : ", ForgotData);
    $.ajax({	
        type: "POST",
        url: "forgotPassword",
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'application/json'
        },
        data:ForgotData,
        success: function (response) {
            console.log("Success response is ", response)
       },
       error: (function(response) {
            console.log("error response is ", response)
       })
   });
});

/* Reset Password Code for Number input Validation */
// $('#password, #repeatPassword').on('keyup', function () {
//     if ($('#password').val() == $('#repeatPassword').val()) {
//       $('#message').html('Password Matching').css('color', 'green');
//     } else {
//       $('#message').html('Password Not Matching').css('color', 'red');
//     }
// });

/* Reset Password Submit Code */
$("#reset_user").submit(function(e){
    e.preventDefault();
    var mobileotp = $("#mobileotp").val();
    var otpreset = '{"otpreset":"'+mobileotp+'"}'
    console.log("otpreset : ", otpreset);
    $.ajax({	
        type: "POST",
        url: "mobileotp",
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'application/json'
        },
        data:otpreset,
        success: function (response) {
            console.log("Success response is ", response)
       },
       error: (function(response) {
            console.log("error response is ", response)
            $('#message').html('Wrong OTP').css('color', 'red');
       })
   });
});
