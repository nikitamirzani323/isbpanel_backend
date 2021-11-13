<script>
    import Home from "./Home.svelte";
   
    export let table_header_font = "";
    export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = true;
    let listHome = [];
    let sData = "";
    let search = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "MOVIE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            initHome("");
        }
    }
    async function initHome(e) {
        listHome = [];
        const res = await fetch("/api/movie", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                movie_search: e
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            movie_no: no,
                            movie_id: record[i]["movie_id"],
                            movie_type: record[i]["movie_type"].toUpperCase(),
                            movie_title: record[i]["movie_title"],
                            movie_descp: record[i]["movie_descp"],
                            movie_year: record[i]["movie_year"],
                            movie_rating: record[i]["movie_rating"],
                            movie_imdb: record[i]["movie_imdb"],
                            movie_view: record[i]["movie_view"],
                            movie_status: record[i]["movie_status"],
                            movie_create: record[i]["movie_create"],
                            movie_update: record[i]["movie_update"],
                        },
                    ];
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
    const handleNews = (e) => {
        search = e.detail.searchNews;
        initHome(search)
   };
    initapp()
</script>
<Home
    on:handleNews={handleNews}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}
/>