<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;
    let perpage = 0;
    let page = 0;
    let totalpaging = 0;
    let totalrecordall = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ALBUM-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        listHome = [];
        const res = await fetch("/api/album", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                album_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            record_message = json.message;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = totalrecordall;
                let no = 0
                if(page > 1){
                    no = parseInt(page) 
                }
                for (var i = 0; i < record.length; i++) {
                    no = parseInt(no) + 1;
                    listHome = [
                        ...listHome,
                        {
                            album_no: no,
                            album_id: record[i]["album_id"],
                            album_idcloud: record[i]["album_idcloud"],
                            album_name: record[i]["album_name"],
                            album_signed: record[i]["album_signed"],
                            album_varian: record[i]["album_varian"],
                            album_movieid: record[i]["album_movieid"],
                            album_movie: record[i]["album_movie"],
                            album_moviestatus: record[i]["album_moviestatus"],
                            album_moviestatuscss: record[i]["album_moviestatuscss"],
                            album_createdate: record[i]["album_createdate"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
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
    const handlePaging = (e) => {
        page = e.detail.page
        initHome()
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {totalrecord}/>
{/if}