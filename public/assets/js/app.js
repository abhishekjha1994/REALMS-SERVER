// ----- Light ON/OFF function starts -----
function light(sw) {
  var pic;
  var light_id = 1234;
  if (sw == 0) {
    
    
    var lightdata = '{"light_id":"'+light_id+'","OnOff":"'+sw+'"}'
    console.log(lightdata)
        $.ajax({	
          type: "POST",
          url: "OnOff",
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          data:lightdata,
          success: function (response) {
              console.log("Success response is ", response)
              pic = "https://cdncontribute.geeksforgeeks.org/wp-content/uploads/ONbulb.jpg"
          },
          error: (function(response) {
              console.log("error response is ", response)
              pic = "https://cdncontribute.geeksforgeeks.org/wp-content/uploads/OFFbulb.jpg"
          })
        });
  } else {
    
    var lightdata = '{"light_id":"'+light_id+'","OnOff":"'+sw+'"}'
    console.log(lightdata)
        $.ajax({	
          type: "POST",
          url: "OnOff",
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          data:sw,
          success: function (response) {
              console.log("Success response is ", response)
              pic = "https://cdncontribute.geeksforgeeks.org/wp-content/uploads/ONbulb.jpg"
          },
          error: (function(response) {
              console.log("error response is ", response)
              pic = "https://cdncontribute.geeksforgeeks.org/wp-content/uploads/OFFbulb.jpg"

          })
        });
  }
  document.getElementById('myImage').src = pic;
}
// ----- Light ON/OFF function Ends -----


// ----- LightSettings slider control code Starts -----
var rangeslider = document.getElementById("sliderRange"); 

rangeslider.oninput = function() { 
  $('#adjustNumber').val(this.value);
} 

var rangeslider2 = document.getElementById("sliderRangeCCT"); 

rangeslider2.oninput = function() { 
  $('#adjustNumberCCT').val(this.value);
} 
// ----- LightSettings slider control code Ends -----


// ----- LightSettings Submit Function Starts -----
$("#lightSettings").click(function(e){
  e.preventDefault();
  var Light_id = 1234;
  var inputAdjustNum = $('#adjustNumber').val();
  if (inputAdjustNum == ""){
    alert("please enter Intensity Value")
    return false;
  }
  var inputAdjustNumCCT = $('#adjustNumberCCT').val();
  if (inputAdjustNumCCT == ""){
    alert("please enter CCT Value")
    return false;
  }
  // alert("Intensity: " +inputAdjustNum+" / "+"CCT: "+inputAdjustNumCCT);
  var lightdata = '{"Light_id":"'+Light_id+'","Intensity":"'+inputAdjustNum+'","CCT":"'+inputAdjustNumCCT+'"}'
    console.log("lightdata : ", lightdata);
    $.ajax({	
        type: "POST",
        url: "registration",
        headers: {
           'Accept': 'application/json',
           'Content-Type': 'application/json'
        },
        data:lightdata,
        success: function (response) {
            console.log("Success response is ", response)
       },
       error: (function(response) {
            console.log("error response is ", response)
       })
   });

});
// ----- LightSettings Submit Function Ends -----
  