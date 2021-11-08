<script>
    import { Input } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";

    export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0

    let title_page = "NEWS"
    let sData = "";
    let myModal = "";
    
    let listnews = []
    let totalrecordnews = 0
    
    let tanggal_start_newsfetch = "";
    let tanggal_end_newsfetch = "";
    let page_newsfetch = "";
    let css_loader = "display: none;";
    let msgloader = "";

    
    const ShowFormNewsFetch = () => {
        sData = "Edit"
        myModal = new bootstrap.Modal(document.getElementById("modalfetchnew"));
        myModal.show();
    };
    async function call_news() {
        listnews = [];
        let KEY_NEWS = "apiKey=25ff185c903e49ddba06551850241e06"
        let COUNTRY_NEWS = "country=id"
        let PAGE_NEWS = "page="+page_newsfetch
        let FROM_NEWS = "from=" + tanggal_start_newsfetch
        let TO_NEWS = "to="+tanggal_end_newsfetch
        let URL_NEWS = "https://newsapi.org/v2/top-headlines?"+KEY_NEWS+"&"+COUNTRY_NEWS+"&"+FROM_NEWS+"&"+TO_NEWS+"&"+PAGE_NEWS
        const res = await fetch(URL_NEWS);
        const json = await res.json();
        let status = json.status;
        let message = json.message;
        let record = json.articles;
        let no = 0;
        if(status == "ok"){
            totalrecordnews = record.length;
            for (var i = 0; i < record.length; i++) {
                no = no + 1
                listnews = [
                            ...listnews,
                    {
                        news_no: no,
                        news_author: record[i]["author"],
                        news_title: record[i]["title"],
                        news_description: record[i]["description"],
                        news_url: record[i]["url"],
                        news_urlToImage: record[i]["urlToImage"],
                        news_publishedat: record[i]["publishedAt"],
                        news_content: record[i]["content"],
                    },
                ];
            }
        }else{
            alert(message)
        }
        
       
        
    }
    function callFunction(event){
        switch(event.detail){
            case "CALL_FORMNEWS":
                ShowFormNewsFetch();
                break;
            case "FETCH_NEWS":
                call_news();
                break;
            case "NEW_KELUARAN":
                ShowNewKeluaran();
                break;
            case "SAVE_KELUARAN":
                handleSaveKeluaran();
                break;
            case "NEW_PREDIKSI":
                ShowNewPrediksi();
                break;
            case "SAVE_PREDIKSI":
                handleSavePrediksi();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
   
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-6">   
            <Button
                on:click={callFunction}
                button_function="CALL_FORMNEWS"
                button_title="Fetch"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecordnews}>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NEWS</th>
                                </tr>
                            </thead>
                            {#if totalrecordnews > 0}
                            <tbody>
                                {#each listnews as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    saveData(rec.news_title,rec.news_description,rec.news_url,rec.news_urlToImage);
                                                }} 
                                                class="bi bi-save"></i>
                                        </td>
                                        
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.news_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            <b>AUTHOR : </b> {rec.news_author}<br>
                                            <a href="{rec.news_url}" target="_blank">{rec.news_title}</a><br>
                                            {@html rec.news_description} <br>
                                            <img width="100" src="{rec.news_urlToImage}" class="img-thumbnail" alt="...">
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
        <div class="col-sm-6">
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NEWS</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each listHome as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    EditData(rec.pasaran_id, rec.pasaran_name, rec.pasaran_url, rec.pasaran_diundi, rec.pasaran_jamjadwal,rec.pasaran_display,rec.pasaran_status);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            <a href="{rec.news_url}" target="_blank">{rec.news_title}</a><br>
                                            <img width="100" src="{rec.news_image}" class="img-thumbnail" alt="...">
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
<Modal
	modal_id="modalfetchnew"
	modal_size="modal-dialog-centered"
	modal_title="{title_page}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="example" class="form-label">Start</label>
			<Input
                bind:value={tanggal_start_newsfetch}
                class="required"
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
		</div>
        <div class="mb-3">
            <label for="example" class="form-label">End</label>
			<Input
                bind:value={tanggal_end_newsfetch}
                class="required"
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
		</div>
        <div class="mb-3">
            <label for="example" class="form-label">Page</label>
            <select 
                class="form-control"
                bind:value={page_newsfetch}>
                <option value="1">Page 1</option>
                <option value="2">Page 2</option>
                <option value="3">Page 3</option>
            </select>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FETCH_NEWS"
            button_title="Submit"
            button_css="btn-warning"/>
	</slot:template>
</Modal>