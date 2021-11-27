<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let sData = "";
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
                page: "KELUARAN-VIEW",
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
        const res = await fetch("/api/pasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
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
                            pasaran_no: no,
                            pasaran_id: record[i]["pasaran_id"],
                            pasaran_name: record[i]["pasaran_name"],
                            pasaran_url: record[i]["pasaran_url"],
                            pasaran_diundi: record[i]["pasaran_diundi"],
                            pasaran_jamjadwal: record[i]["pasaran_jamjadwal"],
                            pasaran_display: record[i]["pasaran_display"],
                            pasaran_status: record[i]["pasaran_status"],
                            pasaran_statuscss: record[i]["pasaran_statuscss"],
                            pasaran_keluaran: record[i]["pasaran_keluaran"],
                            pasaran_prediksi: record[i]["pasaran_prediksi"],
                            pasaran_create: record[i]["pasaran_create"],
                            pasaran_update: record[i]["pasaran_update"],
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
    const handleEditData = (e) => {
        admin_username = e.detail.e;
        sData = "Edit";
        alert(admin_username)
        // editAdmin(admin_username);
    };
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleEditData={handleEditData}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}/>
{/if}