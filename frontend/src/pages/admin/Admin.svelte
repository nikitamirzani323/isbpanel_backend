<script>
    import Home from "../admin/Home.svelte";
    import Entry from "../admin/Entry.svelte";
    let listAdmin = [];
    let listAdminrule = [];
    let sData = "";
    let record = "";
    let totalrecord = 0;
    let admin_username = "";
    export let table_header_font = "";
	export let table_body_font = "";
    let token = localStorage.getItem("token");
    let akses_page = true;
    
    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMIN-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            initAdmin();
        }
    }
    async function initAdmin() {
        const res = await fetch("/api/alladmin", {
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
            totalrecord = record.length;
            let recordlistrule = json.listruleadmin;
            let css_status = "";
            let no = 0
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if (record[i]["admin_status"] == "ACTIVE") {
                        css_status = "background:#8BC34A;color:black;font-weight:bold;";
                    } else {
                        css_status = "background:#E91E63;font-size:12px;font-weight:bold;color:white;";
                    }
                    listAdmin = [
                        ...listAdmin,
                        {
                            admin_no: no,
                            admin_username: record[i]["admin_username"],
                            admin_nama: record[i]["admin_nama"],
                            admin_rule: record[i]["admin_rule"],
                            admin_timezone: record[i]["admin_timezone"],
                            admin_joindate: record[i]["admin_joindate"],
                            admin_lastlogin: record[i]["admin_lastlogin"],
                            admin_lastipaddres: record[i]["admin_lastipaddres"],
                            admin_status: record[i]["admin_status"],
                            admin_statuscss: css_status,
                        },
                    ];
                }
            }
            if (recordlistrule != null) {
                for (let i = 0; i < recordlistrule.length; i++) {
                    listAdminrule = [
                        ...listAdminrule,
                        {
                            adminrule_idruleadmin:recordlistrule[i]["adminrule_idruleadmin"],
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
        listAdmin = [];
        listAdminrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdmin();
        }, 500);
    };
    initapp()
</script>
<Home
    on:handleEditData={handleEditData}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listAdmin}
    {listAdminrule}
    {totalrecord}
/>