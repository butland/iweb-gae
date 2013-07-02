
$.getJSON('/api/blog/article', function(data) {
	var items = [];
	$.each(data, function(key, val) {
		var strHtml = "<div class=\"tile\">"+
            "<h3 class=\"tile-title\">"+val.MetaData.Title+"</h3>"+
			"<p>"+val.MetaData.PostTime.substr(0,10)+" 标签：";
        $.each(val.MetaData.Tags, function(i, tag){
           	strHtml += "<a href=\"/blog/tag/ta\">"+tag+"</a>    "; 
        });        
        strHtml += "</p>";
          
        strHtml +=val.Text;
        strHtml += "<a class=\"btn btn-info btn-large\" "+
        	"href=\"/api/blog/article/"+val.MetaData.Id+"\">Read More</a>"+
            "</div>"+
            "<br/><br/>";

		if(key%2==0){
			$("#article-left").append(strHtml);
		} else{
			$("#article-right").append(strHtml);
		}
	});

});