function openTab(tabname) { 
//function openTab(element, tabname) {
	var i; 
	var x = document.getElementsByClassName('tab'); 
	for (i=0; i<x.length; i++) { 
		x[i].style.display='none'; 
	}	 
	document.getElementById(tabname).style.display = 'block'; 
//    element.style.color = "red";
//    	var x = document.getElementsByClassName('tabbuttons')
//    	for (i =0 ; i <x.length; i++) {
//	    x[i].className= x[i].className.replace('active', '');
//	}
//    	evt.currentTarget.className += ' active';

}

function openForm(formname) {
    document.getElementById(formname).style.display = 'block';
}

function closeForm(formname) {
    document.getElementById(formname).style.display = 'none';
}

function changeColour(element) {
    element.style.color = "red";
    var siblings = element.parentElement.children;
    for (var i =0; i<siblings.length; i++) {
	if (siblings[i] !== element) {
	    siblings[i].style.color = "blue";
	}
    }
}
