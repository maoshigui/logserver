var socket;

$(document).ready(function () {
    // Create a socket
	wsUrl = 'ws://' + window.location.host + '/' + $('#logtype').attr("value") + '/log';
    socket = new WebSocket(wsUrl);
    // Message received on the socket
	socket.onopen = function(){  
		$('#log').append('<p style="color:red;">connected to : ' + wsUrl + '<p>');
		$('#log').get(0).scrollTop=$('#log').get(0).scrollHeight;
	}  

	socket.onclose = function(e) {
		$('#log').append('<p style="color:red;">connection closed (' + e.code + ')<p>');
 		$('#log').get(0).scrollTop=$('#log').get(0).scrollHeight;
	}  
	
    socket.onmessage = function (event) {
        $('#log').append('<p>' + event.data + '<p>');
        $('#log').get(0).scrollTop=$('#log').get(0).scrollHeight;
    };
	
	// Send messages.
    var postConecnt = function () {
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val("");
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
});