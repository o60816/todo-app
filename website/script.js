// getting all required elements
const inputBox = document.querySelector(".inputField input");
const addBtn = document.querySelector(".inputField button");
const todoList = document.querySelector(".todoList");
const footer = document.querySelector(".footer");
const pageTag = document.querySelector(".footer a");
let footerDefault = '<button class="btn active" onclick="deleteAllTasks">Clear All</button>';
let currentPage = 1
let totalPage = 1
// onkeyup event
inputBox.onkeyup = ()=>{
  let userEnteredValue = inputBox.value; //getting user entered value
  if(userEnteredValue.trim() != 0){ //if the user value isn't only spaces
    addBtn.classList.add("active"); //active the add button
  }else{
    addBtn.classList.remove("active"); //unactive the add button
  }
}

showAllTasks()

function showAllTasks(){
  $.get("/items", function(data, status){
    let newLiTag = "";
    data.itemlist.forEach((element, index) => {
      let checked = element.IsDone ? "checked" : "";
      newLiTag += `<li><input type="checkbox" ${checked} id="checkbox${element.Id}" name="checkbox${element.Id}" onclick="updateTasks(${element.Id})">     </input>${element.ItemName}<span class="icon" onclick="deleteTask(${element.Id})">-</span></li>`;
    });
    todoList.innerHTML = newLiTag; //adding new li tag inside ul tag
    inputBox.value = ""; //once task added leave the input field blank

    footer.innerHTML = footerDefault
  });
}

function changePagination(pageSize, toLastPage) {
  if (pageSize=="all"){
    showAllTasks()
  }else{
    $.get("/items/pagesize?pageSize=" +pageSize, function(data, status){
      let newFooter = '';
      let tmpTotalPage = parseInt(data.totalpage);
      
      for (i = 1; i!=tmpTotalPage+1; i++) {
        newFooter += `<a href="javascript:showPagination(${i})">${i}</a>`;
      }
      footer.innerHTML = footerDefault + newFooter
      totalPage = tmpTotalPage
      currentPage = 1
      showPagination(toLastPage ? tmpTotalPage : currentPage)
    });
  }
}

function showPagination(page){
  $.get("/items/pagination?page=" +page, function(data, status){
    let newLiTag = "";
    data.itemlist.forEach((element, index) => {
      var checked = element.IsDone ? "checked" : ""
      newLiTag += `<li><input type="checkbox" ${checked} id="checkbox${element.Id}" name="checkbox${element.Id}" onclick="updateTasks(${element.Id})">     </input>${element.ItemName}<span class="icon" onclick="deleteTask(${element.Id})">-</span></li>`;
    });
    todoList.innerHTML = newLiTag; //adding new li tag inside ul tag
    inputBox.value = ""; //once task added leave the input field blank
    currentPage = page
  });
}

addBtn.onclick = ()=>{ //when user click on plus icon button
  let userEnteredValue = inputBox.value; //getting input field value
  $.post("/items", {"itemName":userEnteredValue}, function(data, status){
    addBtn.classList.remove("active"); //unactive the add button once the task added
    changePagination($("#pagesize").get(0).value, true)
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
      showPagination(currentPage)
    }
 });
}

function deleteTask(id){
  $.ajax({
    type: 'DELETE',
    url: `/items/${id}`,
    success: function(data){
      changePagination($("#pagesize").get(0).value)
    }
 });
}

function deleteAllTasks(){
  $.ajax({
    type: 'DELETE',
    url: '/items',
    success: function(data){
      changePagination($("#pagesize").get(0).val)
    }
 });
}
