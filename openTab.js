function openTab(evt, tabname) { 
	var i; 
	var x = document.getElementsByClassName('tab'); 
	for (i=0; i<x.length; i++) { 
		x[i].style.display='none'; 
	}	 
	document.getElementById(tabname).style.display = 'block'; 
    	var x = document.getElementsByClassName('tabbuttons')
    	for (i =0 ; i <x.length; i++) {
	    x[i].className= x[i].className.replace('active', '');
	}
    	evt.currentTarget.className += ' active';

}

function openForm(formname) {
    document.getElementById(formname).style.display = 'block';
}

function closeForm(formname) {
    document.getElementById(formname).style.display = 'none';
}

