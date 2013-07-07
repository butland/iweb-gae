var path = window.location.pathname;
if(path=="/"){
	path = "/api/blog/article";
}
var urlHash = window.location.hash;
if(urlHash.length>0){
	readMore();
}else{
	listArticle(path);
}
function listArticle(url){
	$.getJSON(url, function(data) {
		$.each(data, function(key, val) {
			var strHtml = "<div class=\"tile\">"+
			"<h3 class=\"tile-title\">"+val.MetaData.Title+"</h3>"+
			"<p>"+val.MetaData.PostTime.substr(0,10)+" 标签：";
			$.each(val.MetaData.Tags, function(i, tag){
				strHtml += "<a href=\"#!/api/blog/tag/"+tag+"\">"+tag+"</a>    "; 
			});        
			strHtml += "</p>";

			strHtml +=val.Text;
			strHtml += "<a class=\"btn btn-info btn-large readMore\" "+
			"href=\"/blog/"+val.MetaData.PostTime.substr(0,4)+"/"+
			val.MetaData.PostTime.substr(5,2)+"/"+val.MetaData.PostTime.substr(8,2)+"/"+
			val.MetaData.Title+"\">Read More</a>"+
			"</div>"+
			"<br/><br/>";

			if(key%2==0){
				$("#article-left").append(strHtml);
			} else{
				$("#article-right").append(strHtml);
			}
		});

	});
}


$(window).bind('hashchange', function(e){
	var url = window.location.hash;
	//alert(url.substr(2));
	if (url.length>0){
		$("#article-left").empty();
		$("#article-right").empty();
    	listArticle(url.substr(2));
	}else{
		$("#article-left").empty();
		$("#article-right").empty();
		listArticle(path);
	}
});

function readMore(){
	var url = window.location.hash;
	$.getJSON(url.substr(2), function(data) {
		var strHtml = "<div class=\"article\">"+
    		"<h1>"+data.MetaData.Title+"</h1>"+
    		"<p>"+data.MetaData.PostTime.substr(0,10)+" 标签：";
    	$.each(data.MetaData.Tags, function(i, tag){
			strHtml += "<a href=\"/blog/tag/ta\">"+tag+"</a>    "; 
		}); 
    	strHtml += "</p>";
    	strHtml +=data.Text;
		strHtml +="</div>";
		$(".article").append(strHtml);
		$("#article-list").hide();
	});
}


