/* attach a submit handler to the form */
$("#postComment").submit(function(event) {

  /* stop form from submitting normally */
  event.preventDefault();

  /* get some values from elements on the page: */
  var $form = $( this ),
  urlStr = $form.find( 'input[name="urlStr"]' ).val(),
  articleId = $form.find( 'input[name="articleId"]' ).val(),
  name = $form.find( 'input[name="name"]' ).val(),
  email = $form.find( 'input[name="email"]' ).val(),
  website = $form.find( 'input[name="website"]' ).val(),
  content = $form.find( 'textarea[name="content"]' ).val(),
  url = $form.attr( 'action' );

  /* Send the data using post */
  var posting = $.post( url, { 
    urlStr: urlStr,
    articleId:articleId,
    name:name,
    email:email,
    website:website,
    content:content,
  } );

  /* Put the results in a div */
  posting.done(function( data ) {
    var content = "<p align=\"center\" id=\"commentMessage\" data-toggle=\"tooltip\" "+
      "title=\" " + data.Text+"\"></p>";
      
    $("#commentTip").empty();
    $("#commentTip").append(content);
    $("#commentMessage").tooltip("show");
    function doFinished(){
      $("#commentMessage").tooltip("hide");
    }
    setTimeout(doFinished,"5000");//启动一个定时器，5秒后隐藏
    if(data.Ok){
      $('#commentList').empty();
      loadComment(articleId);
    }
  });
});

function loadComment(articleId){
  var url = "/api/blog/comment?id="+articleId;
  $.getJSON(url, function(data) {
    $.each(data, function(key, val) {
      var strHtml = "<li class=\"media\">"+
              "<a class=\"pull-left\" href=\"#\">"+
              "<img class=\"media-object\" src=\"/themes/flat/bootstrap/img/holder.png\">"+
              "</a>"+
              "<div class=\"media-body\">"+
              "<h6 class=\"media-heading\">Written by "+val.Author+
              " on "+val.PostTime.substr(0,10)+" </h6>"+val.Content+    
              "</div> </li>";

      $("#commentList").append(strHtml);
    });
      
  });

}