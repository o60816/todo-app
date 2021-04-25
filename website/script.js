// getting all required elements
const inputBox = document.querySelector(".inputField input");
const addBtn = document.querySelector(".inputField button");
const todoList = document.querySelector(".todoList");
const deleteAllBtn = document.querySelector(".footer button");
// onkeyup event
inputBox.onkeyup = ()=>{
  let userEnteredValue = inputBox.value; //getting user entered value
  if(userEnteredValue.trim() != 0){ //if the user value isn't only spaces
    addBtn.classList.add("active"); //active the add button
  }else{
    addBtn.classList.remove("active"); //unactive the add button
  }
}

showTasks(); //calling showTask function

function showTasks(){
  $.get("/items", function(data, status){
    var listArray = Object.keys(data).map((key) => data[key]);
    let newLiTag = "";

    listArray.forEach((element, index) => {
      var checked = element.IsDone ? "checked" : ""
      newLiTag += `<li><input type="checkbox" ${checked} id="checkbox${element.Id}" name="checkbox${element.Id}" onclick="updateTasks(${element.Id})">     </input>${element.ItemName}<span class="icon" onclick="deleteTask(${element.Id})">-</span></li>`;
    });
    todoList.innerHTML = newLiTag; //adding new li tag inside ul tag
    inputBox.value = ""; //once task added leave the input field blank
  });
}

addBtn.onclick = ()=>{ //when user click on plus icon button
  let userEnteredValue = inputBox.value; //getting input field value
  $.post("/items", {"itemName":userEnteredValue}, function(data, status){
    addBtn.classList.remove("active"); //unactive the add button once the task added
    showTasks()
    alert(data.message);
  });
}

function updateTasks(id){
  var patch = {
    "isDone" : $(`#checkbox${id}`).is(':checked') ? "1" : "0"
  }
  $.ajax({
    type: 'PATCH',
    url: `/items/${id}`,
    data: JSON.stringify(patch),
    processData: true,
    contentType: 'application/merge-patch+json',
    success: function(data){
      showTasks()
      alert(data.message);
    }
 });
}

function deleteTask(id){
  $.ajax({
    type: 'DELETE',
    url: `/items/${id}`,
    success: function(data){
      showTasks()
      alert(data.message);
    }
 });
}

// delete all tasks function
deleteAllBtn.onclick = ()=>{
  $.ajax({
    type: 'DELETE',
    url: '/items',
    success: function(data){
      showTasks()
      alert(data.message);
    }
 });
}
