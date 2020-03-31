$(document).ready(function () {
    $('.loading').removeClass('d-hide');
    val = window.location.search;
    if (val === "?q=&auto_redirect=true") {
        val = "?q=https://golang.google.cn/doc/effective_go.html"
    }
    console.log("val: ", val);
    $.ajax({
        url: "/search" + val,
        type: "POST",
        error: function (err) {
            console.log("=== 页面加载，出现错误 ===", JSON.stringify(err));
            // window.location.href = "/";
        },
        success: function (data) {
            console.log("解析的data", data);
            $('.loading').addClass('d-hide');
            $('#search-results').html("");
            $('#search-results').append(`
                <div class="tile tile-centered">
                    <div class="tile-content">
                        <p class="tile-title"><b>` + data + `</b></p>
                    </div>
                </div>
            `)
            // for (var i = 0; i < data.results.length; i++) {
            //     $('#search-results').append(`<a href="` + data.results[i].url + `">
            //         <div class="tile tile-centered">
            //             <div class="tile-content">
            //                 <p class="tile-title"><b>` + data.results[i].title + `</b></p>
            //                 <p class="tile-subtitle">` + data.results[i].description + `</p>
            //             </div>
            //         </div>
            //     </a>`);
            //     if (i+1 < data.results.length) {
            //         $('#search-results').append(`<div class="divider"></div>`)
            //     }
            // }
        }
    });
});