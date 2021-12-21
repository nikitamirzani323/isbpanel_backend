<script>
  import { Input } from "sveltestrap";
  import Panel from "../../components/Panel.svelte";
  import Loader from "../../components/Loader.svelte";
  import Button from "../../components/Button.svelte";
  import Modal from "../../components/Modal.svelte";
  import { createEventDispatcher } from "svelte";

  export let table_header_font = "";
  export let table_body_font = "";
  export let token = "";
  export let listHome = [];
  export let listPage = [];
  export let totalrecord = 0;
  let dispatch = createEventDispatcher();
  let title_page = "DUNIA FILM";
  let sData = "";
  let myModal = "";

  let listcategory = [];

  let pagingnow = 0;

  let searchIsbtv = "";
  let filterIsbtv = "";

  let css_loader = "display: none;";
  let msgloader = "";

  $: {
    if (searchIsbtv) {
      filterIsbtv = listHome.filter(
        (item) =>
          item.crmduniafilm_name
            .toLowerCase()
            .includes(searchIsbtv.toLowerCase()) ||
          item.crmduniafilm_username
            .toLowerCase()
            .includes(searchIsbtv.toLowerCase())
      );
    } else {
      filterIsbtv = [...listHome];
    }
  }
  const RefreshHalaman = () => {
    dispatch("handleRefreshData", "call");
  };
  const handleSelectPaging = (event) => {
    let page = event.target.value;
    pagingnow = page;
    const movie = {
      page,
    };
    dispatch("handlePaging", movie);
  };

  function callFunction(event) {
    switch (event.detail) {
      case "REFRESH":
        RefreshHalaman();
        break;
    }
  }

  const handleKeyboard_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      filterIsbtv = [];
      listHome = [];
      const news = {
        searchIsbtv,
      };
      dispatch("handleNews", news);
    }
  };
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
  {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
  <div class="row">
    <div class="col-sm-12">
      <Button
        on:click={callFunction}
        button_function="REFRESH"
        button_title="Refresh"
        button_css="btn-primary"
      />

      <Panel
        card_search={true}
        card_title={title_page}
        card_footer={totalrecord}
      >
        <slot:template slot="card-title">
          <div class="float-end">
            <select
              on:change={handleSelectPaging}
              style="text-align: center;"
              class="form-control"
            >
              {#each listPage as rec}
                <option value={rec.page_value}>{rec.page_display}</option>
              {/each}
            </select>
          </div>
        </slot:template>
        <slot:template slot="card-search">
          <div class="col-lg-12" style="padding: 5px;">
            <input
              bind:value={searchIsbtv}
              on:keypress={handleKeyboard_checkenter}
              type="text"
              class="form-control"
              placeholder="Search Username"
              aria-label="Search"
            />
          </div>
        </slot:template>
        <slot:template slot="card-body">
          <table class="table table-striped table-hover table-sm">
            <thead>
              <tr>
                <th
                  NOWRAP
                  width="1%"
                  style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
                  >NO</th
                >
                <th
                  NOWRAP
                  width="*"
                  style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >NAME</th
                >
                <th
                  NOWRAP
                  width="10%"
                  style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >USERNAME</th
                >
              </tr>
            </thead>
            {#if totalrecord > 0}
              <tbody>
                {#each filterIsbtv as rec}
                  <tr>
                    <td
                      NOWRAP
                      style="text-align: center;vertical-align: top;font-size: {table_body_font};"
                      >{rec.news_no}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: left;vertical-align: top;font-size: {table_body_font};"
                      >{rec.crmduniafilm_name}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: left;vertical-align: top;font-size: {table_body_font};"
                    >
                      <a
                        href="https://wa.me/{rec.crmduniafilm_username}"
                        target="_blank">WHATSAPP</a
                      >
                      &nbsp;
                    </td>
                  </tr>
                {/each}
              </tbody>
            {:else}
              <tbody>
                <tr>
                  <td colspan="20">
                    <center>
                      <Loader />
                    </center>
                  </td>
                </tr>
              </tbody>
            {/if}
          </table>
        </slot:template>
      </Panel>
    </div>
  </div>
</div>
