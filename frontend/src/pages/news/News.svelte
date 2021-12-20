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
        page: "NEWS-VIEW",
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
    const res = await fetch("/api/news", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        news_search: e,
        news_page: parseInt(page),
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
        if (page > 1) {
          no = parseInt(page);
        }
        for (var i = 0; i < record.length; i++) {
          no = parseInt(no) + 1;
          listHome = [
            ...listHome,
            {
              news_no: no,
              news_id: record[i]["news_id"],
              news_idcategory: record[i]["news_idcategory"],
              news_category: record[i]["news_category"],
              news_title: record[i]["news_title"],
              news_descp: record[i]["news_descp"],
              news_url: record[i]["news_url"],
              news_image: record[i]["news_image"],
              news_create: record[i]["news_create"],
              news_update: record[i]["news_update"],
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
    search = e.detail.searchNews;
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
