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
                page: "CRM-VIEW",
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
        const res = await fetch("/api/crm", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                crm_search: e,
                crm_page : parseInt(page)
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
                console.log(totalpaging)
                totalrecord = totalrecordall;
                let no = 0
                let css_point = "";
                if(page > 1){
                    no = parseInt(page) 
                }
                for (var i = 0; i < record.length; i++) {
                    if(parseInt(record[i]["crmisbtv_point"]) > 0){
                        css_point = "color:blue;font-weight:bold;"
                    }else{  
                        css_point = "color:red;font-weight:bold;"
                    }
                    no = parseInt(no) + 1;
                    listHome = [
                        ...listHome,
                        {
                            news_no: no,
                            crm_id: record[i]["crm_id"],
                            crm_phone: record[i]["crm_phone"],
                            crm_name: record[i]["crm_name"],
                            crm_source: record[i]["crm_source"],
                            crm_status: record[i]["crm_status"],
                            crm_statuscss: record[i]["crm_statuscss"],
                            crm_create: record[i]["crm_create"],
                            crm_update: record[i]["crm_update"],
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
        console.log(e)
        search = e.detail.searchIsbtv;
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
    on:handleNews={handleNews}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {totalrecord}/>
{/if}