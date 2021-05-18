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

//function changeColour(element) {
//    var siblings = element.parentElement.children;
//    for (var i =0; i<siblings.length; i++) {
//	if (siblings[i] !== element) {
//	    siblings[i].style.color = "black";
//    	siblings[i].setAttribute('data-selected', 'false');
//	}
//    }
//    element.style.color = "red";
//    element.setAttribute('data-selected', 'true');
//}

function changeColour(element) {
    //var siblings = document.getElementsByClassName('netbutton');
    var siblings = element.parentNode.children
    for (var i =0 ; i<siblings.length; i++) {
	siblings[i].style.color = 'black';
    	siblings[i].setAttribute('data-selected', 'false')
    }
    element.style.color = 'blue';
    element.setAttribute('data-selected', 'true')
}

function displayTab() {
    var nets, details, net, detail, i, selected;
    //set default for tabname
    var tabname = 'All Networks-Network Details';
    nets = document.getElementsByClassName('netbutton');
    for (i =0; i<nets.length; i++) {
	selected = nets[i].getAttribute('data-selected');
	if (selected == 'true') {
	    net = nets[i].innerText;
	}
    }
    console.log(net)
    details = document.getElementsByClassName('tabbutton');
    for (i =0; i<details.length; i++) {
	selected = details[i].getAttribute('data-selected');
	if (selected == 'true') {
	    detail = details[i].innerText;
	}
    }
    console.log(detail)
    tabname = net+"-"+detail
    console.log(tabname)
    openTab(tabname)
    }

