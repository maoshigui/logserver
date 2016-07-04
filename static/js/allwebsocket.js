var socket;

$(document).ready(function () {
    // Create a socket
	wsUrl = 'ws://' + window.location.host + '/' + $('#logtype').attr("value") + '/goctlog';
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
	
	// Create a socket
	wsUrlMs = 'ws://' + window.location.host + '/' + $('#logtype').attr("value") + '/mslog';
    socketMs = new WebSocket(wsUrlMs);
    // Message received on the socket
	socketMs.onopen = function(){  
		$('#logMs').append('<p style="color:red;">connected to : ' + wsUrlMs + '<p>');
		$('#logMs').get(0).scrollTop=$('#logMs').get(0).scrollHeight;
	}  

	socketMs.onclose = function(e) {
		$('#logMs').append('<p style="color:red;">connection closed (' + e.code + ')<p>');
 		$('#logMs').get(0).scrollTop=$('#logMs').get(0).scrollHeight;
	}  
	
    socketMs.onmessage = function (event) {
        $('#logMs').append('<p>' + event.data + '<p>');
        $('#logMs').get(0).scrollTop=$('#logMs').get(0).scrollHeight;
    };
});