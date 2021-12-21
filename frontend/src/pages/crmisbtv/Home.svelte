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
  let title_page = "ISBTV";
  let sData = "";
  let myModal = "";

  let listcategory = [];

  let pagingnow = 0;

  let category_field_idrecord = 0;
  let category_field_name = "";
  let category_field_display = 0;
  let category_field_status = "";
  let news_field_idrecord = 0;
  let news_field_title = "";
  let news_field_descp = "";
  let news_field_category = "";
  let news_field_url = "";
  let news_field_image = "";
  let searchIsbtv = "";
  let filterIsbtv = "";

  let css_loader = "display: none;";
  let msgloader = "";

  $: {
    if (searchIsbtv) {
      filterIsbtv = listHome.filter(
        (item) =>
          item.crmisbtv_name
            .toLowerCase()
            .includes(searchIsbtv.toLowerCase()) ||
          item.crmisbtv_username
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
                <th
                  NOWRAP
                  width="10%"
                  style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >REF</th
                >
                <th
                  NOWRAP
                  width="10%"
                  style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >POINT</th
                >
                <th
                  NOWRAP
                  width="20%"
                  style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >LASTLOGIN</th
                >
                <th
                  NOWRAP
                  width="20%"
                  style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >CREATE</th
                >
                <th
                  NOWRAP
                  width="20%"
                  style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};"
                  >UPDATE</th
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
                      >{rec.crmisbtv_name}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: left;vertical-align: top;font-size: {table_body_font};"
                    >
                      <a
                        href="https://wa.me/{rec.crmisbtv_username}"
                        target="_blank">WHATSAPP</a
                      >
                      &nbsp;
                    </td>
                    <td
                      NOWRAP
                      style="text-align: left;vertical-align: top;font-size: {table_body_font};"
                      >{rec.crmisbtv_coderef}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.crmisbtv_csspoint}"
                      >{rec.crmisbtv_point}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: center;vertical-align: top;font-size: {table_body_font};"
                      >{rec.crmisbtv_lastlogin}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: center;vertical-align: top;font-size: {table_body_font};"
                      >{rec.crmisbtv_create}</td
                    >
                    <td
                      NOWRAP
                      style="text-align: center;vertical-align: top;font-size: {table_body_font};"
                      >{rec.crmisbtv_update}</td
                    >
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

<Modal
  modal_id="modalcrudnews"
  modal_size="modal-dialog-centered"
  modal_title="NEWS/{sData}"
  modal_body_css=""
  modal_footer_css="padding:5px;"
  modal_footer={true}
>
  <slot:template slot="body">
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Category</label>
      <select bind:value={news_field_category} class="form-control required">
        {#each listcategory as rec}
          <option value={rec.category_id}>{rec.category_name}</option>
        {/each}
      </select>
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Title</label>
      <Input
        bind:value={news_field_title}
        class="required"
        type="text"
        placeholder="News Title"
      />
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Deskripsi</label>
      <textarea
        style="height: 100px;resize: none;"
        bind:value={news_field_descp}
        class="form-control required"
      />
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Url Image</label>
      <Input
        bind:value={news_field_image}
        class="required"
        type="text"
        placeholder="News Image"
      />
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Url</label>
      <Input
        bind:value={news_field_url}
        class="required"
        type="text"
        placeholder="News URL"
      />
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="SAVE_NEWS"
      button_title="Save"
      button_css="btn-warning"
    />
  </slot:template>
</Modal>
<Modal
  modal_id="modalcategory"
  modal_size="modal-dialog-centered"
  modal_title="CATEGORY"
  modal_body_css="height:500px;"
  modal_footer_css="padding:5px;"
  modal_footer={true}
>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
          <th width="1%" colspan="2">&nbsp;</th>
          <th
            width="1%"
            style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
            >NO</th
          >
          <th
            width="1%"
            style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
            >STATUS</th
          >
          <th
            width="*"
            style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
            >CATEGORY</th
          >
          <th
            width="5%"
            style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
            >NEWS</th
          >
          <th
            width="5%"
            style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};"
            >DISPLAY</th
          >
        </tr>
      </thead>
      <tbody>
        {#each listcategory as rec}
          <tr>
            <td
              NOWRAP
              style="text-align: center;vertical-align: top;cursor:pointer;"
            >
              <i
                on:click={() => {
                  ShowFormCategory(
                    "Edit",
                    rec.category_id,
                    rec.category_name,
                    rec.category_display,
                    rec.category_status
                  );
                }}
                class="bi bi-pencil"
              />
            </td>
            <td
              NOWRAP
              style="text-align: center;vertical-align: top;cursor:pointer;"
            >
              <i
                on:click={() => {
                  handleDeleteCategoryNews(rec.category_id);
                }}
                class="bi bi-trash"
              />
            </td>
            <td
              NOWRAP
              style="text-align: center;vertical-align: top;font-size: {table_body_font};"
              >{rec.category_no}</td
            >
            <td
              NOWRAP
              style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.category_statuscss}"
              >{rec.category_status}</td
            >
            <td
              NOWRAP
              style="text-align: left;vertical-align: top;font-size: {table_body_font};"
              >{rec.category_name}</td
            >
            <td
              NOWRAP
              style="text-align: right;vertical-align: top;font-size: {table_body_font};"
              >{rec.category_totalnews}</td
            >
            <td
              NOWRAP
              style="text-align: right;vertical-align: top;font-size: {table_body_font};"
              >{rec.category_display}</td
            >
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="FORMNEW_CATEGORY"
      button_title="New"
      button_css="btn-warning"
    />
  </slot:template>
</Modal>
<Modal
  modal_id="modalcrudcategory"
  modal_size="modal-dialog-centered"
  modal_title="CATEGORY/{sData}"
  modal_body_css=""
  modal_footer_css="padding:5px;"
  modal_footer={true}
>
  <slot:template slot="body">
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Name</label>
      <Input
        bind:value={category_field_name}
        class="required"
        type="text"
        placeholder="Category Name"
      />
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Display</label>
      <Input
        bind:value={category_field_display}
        class="required"
        maxlength="3"
        type="text"
        style="text-align:right;"
        placeholder="Category Display"
      />
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Status</label>
      <select bind:value={category_field_status} class="form-control required">
        <option value="Y">ACTIVE</option>
        <option value="">DEACTIVE</option>
      </select>
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="SAVE_CATEGORY"
      button_title="Save"
      button_css="btn-warning"
    />
  </slot:template>
</Modal>
