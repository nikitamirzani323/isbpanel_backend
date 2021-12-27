<script>
  import Home from "./Home.svelte";
  import dayjs from "dayjs";

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
        page: "CRMISBTV-VIEW",
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
    const res = await fetch("/api/crmisbtv", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        crmisbtv_search: e,
        crmisbtv_page: parseInt(page),
      }),
    });
    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      perpage = json.perpage;
      totalrecordall = json.totalrecord;
      record_message = json.message;
      if (record != null) {
        totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage));
        totalrecord = totalrecordall;
        let no = 0;
        let css_point = "";
        if (page > 1) {
          no = parseInt(page);
        }
        for (var i = 0; i < record.length; i++) {
          if (record[i]["crmisbtv_name"] === "Viandra Natari")
            console.log(record[i]["crmisbtv_lastlogin"]);
          if (parseInt(record[i]["crmisbtv_point"]) > 0) {
            css_point = "color:blue;font-weight:bold;";
          } else {
            css_point = "color:red;font-weight:bold;";
          }
          no = parseInt(no) + 1;
          listHome = [
            ...listHome,
            {
              news_no: no,
              crmisbtv_username: record[i]["crmisbtv_username"],
              crmisbtv_name: record[i]["crmisbtv_name"],
              crmisbtv_coderef: record[i]["crmisbtv_coderef"],
              crmisbtv_point: record[i]["crmisbtv_point"],
              crmisbtv_csspoint: css_point,
              crmisbtv_status: record[i]["crmisbtv_status"],
              crmisbtv_lastlogin: dayjs(record[i]["crmisbtv_lastlogin"]).format(
                "YYYY-MM-DD HH:mm:ss"
              ),
              crmisbtv_create: dayjs(record[i]["crmisbtv_create"]).format(
                "YYYY-MM-DD HH:mm:ss"
              ),
              crmisbtv_update: dayjs(record[i]["crmisbtv_update"]).format(
                "YYYY-MM-DD HH:mm:ss"
              ),
            },
          ];
        }
        listPage = [];
        for (var i = 1; i < totalpaging; i++) {
          listPage = [
            ...listPage,
            {
              page_id: i,
              page_value: i * perpage - perpage,
              page_display: i + " Of " + perpage * i,
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
    console.log(e);
    search = e.detail.searchIsbtv;
    initHome(search);
  };
  const handlePaging = (e) => {
    page = e.detail.page;
    initHome("");
  };
  initapp();
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
    {totalrecord}
  />
{/if}
