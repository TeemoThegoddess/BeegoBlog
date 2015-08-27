function checkInput(){
	var account = document.getElementById("account");
	if (account.value.length == 0){
		alert("Please enter your account name");
		return false
	}

	var pwd = document.getElementById("pwd");
	if (pwd.value.length == 0){
		alert("Please enter your password");
		return false
	} 

	return true
}

function backToHome(){
	window.location.href = "/";
	return false
}

function checkCategories(){
	var category = document.getElementById("category");
	if (category.value.length == 0){
		alert("Please enter your category name.If there is no any category,please go to the category page to add one");
		return false
	}

	return true
}

function delConfirm(operation){
	switch(operation)
	{
		case "delTopic":
			if(!confirm("Are you sure to delete this topic?")){
				window.event.returnValue = false;
			}
			break;
		case "delCategory":
			 if(!confirm("Are you sure to delete this category?(All topic under this category will be deleted )")){
				window.event.returnValue = false;
			}
			break;
		default:
			if(!confirm("Are you sure to delete?")){
				window.event.returnValue = false;
			}
	}
	
}