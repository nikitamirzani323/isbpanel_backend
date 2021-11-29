<script>
    import Home from "./Home.svelte";
   
    export let table_header_font = "";
    export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let sData = "";
    let search = "";
    let record = "";
    let record_message = "";
    let perpage = 0;
    let page = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "SERIES-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome("");
        }
    }
    async function initHome(e) {
        listHome = [];
        const res = await fetch("/api/movieseries", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                movie_search: e,
                movie_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            totalrecord = json.totalrecord;
            record_message = json.message;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                console.log("PERPAGE :" + perpage)
                console.log("PAGE :" + page)
                console.log("TOTAL PAGING :" + totalpaging)
                console.log("TOTAL ALL :" + totalrecordall)
                let no = 0
                if(page > 1){
                    no = parseInt(page) 
                }
                for (var i = 0; i < record.length; i++) {
                    let genre = record[i]["movie_genre"]
                    let season = record[i]["movie_season"]
                    if(record[i]["movie_genre"] == null){
                        genre = []
                    }
                    if(record[i]["movie_season"] == null){
                        season = []
                    }
                    let css_type = "background-color:#0dcaf0;font-weight:bold;"
                    if(record[i]["movie_type"] == "movie"){
                        css_type = "background-color:#ffc107;font-weight:bold;"
                    }
                    let css_cdn = "background-color:#dc3545;font-weight:bold;color:white;"
                    if(record[i]["movie_imgcdn"] == "Y"){
                        css_cdn = "background-color:#ffc107;font-weight:bold;"
                    }
                    let css_imdb = "font-weight:bold;color:red;"
                    if(parseInt(record[i]["movie_imdb"]) > 0){
                        css_imdb = "font-weight:bold;color:blue;"
                    }
                    let css_view = "font-weight:bold;color:red;"
                    if(parseInt(record[i]["movie_view"]) > 0){
                        css_view = "font-weight:bold;color:blue;"
                    }
                    let css_comment = "font-weight:bold;color:red;"
                    if(parseInt(record[i]["movie_comment"]) > 0){
                        css_comment = "font-weight:bold;color:blue;"
                    }
                    no = parseInt(no) + 1;
                    listHome = [
                        ...listHome,
                        {
                            movie_no: no,
                            movie_id: record[i]["movie_id"],
                            movie_date: record[i]["movie_date"],
                            movie_type: record[i]["movie_type"].toUpperCase(),
                            movie_csstype: css_type,
                            movie_title: record[i]["movie_title"],
                            movie_label: record[i]["movie_label"],
                            movie_slug: record[i]["movie_slug"],
                            movie_descp: record[i]["movie_descp"],
                            movie_imgcdn: record[i]["movie_imgcdn"],
                            movie_css_cdn: css_cdn,
                            movie_thumbnail: record[i]["movie_thumbnail"], 
                            movie_year: record[i]["movie_year"].toString(),
                            movie_rating: record[i]["movie_rating"],
                            movie_imdb: record[i]["movie_imdb"],
                            movie_imdbcss: css_imdb,
                            movie_view: record[i]["movie_view"],
                            movie_viewcss: css_view,
                            movie_comment: record[i]["movie_comment"],
                            movie_commentcss: css_comment,
                            movie_genre: genre,
                            movie_season: season,
                            movie_status: record[i]["movie_status"],
                            movie_statuscss: record[i]["movie_statuscss"],
                            movie_create: record[i]["movie_create"],
                            movie_update: record[i]["movie_update"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                    console.log((i*perpage)-perpage)
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    const handleMovie = (e) => {
        search = e.detail.searchMovie;
        initHome(search)
    };
    const handlePaging = (e) => {
        page = e.detail.page
        initHome("")
    };
    initapp()
</script>
{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleMovie={handleMovie}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {totalrecord}/>
{/if}