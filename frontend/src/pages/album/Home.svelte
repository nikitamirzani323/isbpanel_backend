<script>
   
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
    export let listPage = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "ALBUM"
    let sData = "";
    let myModal_newentry = "";
    let listalbumcloudflare = []
    let page_field = "";

    let domain_field = "";
    let status_field = "";
    let idrecord = "";
    let searchAlbum = "";
    let filterAlbum = [];
    let css_loader = "display: none;";
    let msgloader = "";
    let pagingnow = 0;

    $: {
        if (searchAlbum) {
            filterAlbum = listHome.filter(
                (item) =>
                    item.album_movie
                        .toLowerCase()
                        .includes(searchAlbum.toLowerCase())
            );
        } else {
            filterAlbum = [...listHome];
        }
    }
    
    const NewData = () => {
        
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    const handleSelectPaging = (event) => {
        let page = event.target.value
        pagingnow = page
        const movie = {
                page,
        };
        dispatch("handlePaging", movie);
    };
    async function handleGetCloudflare(page) {
        listalbumcloudflare = []
        const res = await fetch("/api/cloudflare", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"ALBUM-VIEW",
                album_page: 1,
                album_pagecloud: parseInt(page),
            }),
        });
        const json = await res.json();
        if (json.status == true) {
            let record = json.record;
            let images = record.images;
            let no = 0;
            let perpage = 100
            if(page > 1){
                no = page * perpage
            }
            if (record != null) {
                for (var i = 0; i < images.length; i++) {
                    let signed = ""
                    if(images[i]["requireSignedURLs"] == true){
                        signed = "LOCK"
                    }
                    no = no + 1;
                    listalbumcloudflare = [
                        ...listalbumcloudflare,
                        {
                            album_no: no,
                            album_filename: images[i]["filename"],
                            album_id: images[i]["id"],
                            album_link_0: images[i]["variants"][0],
                            album_signed: signed,
                        },
                    ];
                }
            }
        }
    }
    async function handleDownload() {
        console.log(listalbumcloudflare.length)
        let flag = true
        let msg = ""
        if(listalbumcloudflare.length < 1){
            flag = false
            msg += "The Album is required\n"
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/albumsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"ALBUM-SAVE",
                    album_page: 1,
                    album_data: listalbumcloudflare,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    
    async function handleUpdateCloudflare(e,id,idalbum) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/moviecloudupdate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: "Edit",
                page:"MOVIE-SAVE",
                movie_id: id,
                movie_tipe: e,
                album_id: parseInt(idalbum),
            }),
        });
        const json = await res.json();
        const status = json.status;
        if(status == true){
            RefreshHalaman()
            msgloader = "Success";
        }
        
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        
    }
    async function handleDeleteCloudflare(idcloudflare,idalbum,idmovie) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/movieclouddelete", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: "Edit",
                page:"MOVIE-DELETE",
                cloudflare_id: idcloudflare,
                album_id: parseInt(idalbum),
                movie_id: parseInt(idmovie),
            }),
        });
        const json = await res.json();
        const status = json.status;
        if(status == true){
            RefreshHalaman()
            msgloader = "Success";
        }
        
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "GETDATA":
                handleGetCloudflare();break;
            case "DOWNLOAD":
                handleDownload();break;
        }
    }
  
    const handleSelectGet = (event) => {
        let page = event.target.value
        handleGetCloudflare(page)
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
                button_function="NEW"
                button_title="Fetch Cloudflare"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-title">
                    <div class="float-end">
                        <select
                            on:change={handleSelectPaging}
                            style="text-align: center;" 
                            class="form-control">
                            {#each listPage as rec}
                                <option value="{rec.page_value}">{rec.page_display}</option>
                            {/each}
                        </select>
                    </div>
                </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchAlbum}
                            type="text"
                            class="form-control"
                            placeholder="Search Album"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped table-hover">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">FILENAME</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LINK</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MOVIE</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MOVIE STATUS</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterAlbum as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        {#if rec.album_name != "placeholder.png"}
                                        <i 
                                            on:click={() => {
                                                handleDeleteCloudflare(rec.album_idcloud,rec.album_id,rec.album_movieid);
                                            }} 
                                            class="bi bi-trash"></i>
                                        {/if}
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.album_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};cursor:pointer;">
                                        {#if rec.album_name != "placeholder.png"}
                                            {#if rec.album_signed == "LOCK"}
                                            <i 
                                                on:click={() => {
                                                    handleUpdateCloudflare("UNLOCK",rec.album_idcloud,rec.album_id);
                                                }} 
                                                class="bi bi-lock-fill"></i>
                                            {:else}
                                            <i 
                                                on:click={() => {
                                                    handleUpdateCloudflare("LOCK",rec.album_idcloud,rec.album_id);
                                                }} 
                                                class="bi bi-unlock"></i>
                                            {/if}
                                        {/if}
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.album_idcloud}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.album_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <a href="{rec.album_varian}" target="_blank">LINK</a>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        ID : {rec.album_movieid}<br>
                                        TITLE : {rec.album_movie}
                                    </td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.album_moviestatuscss}">{rec.album_moviestatus}</td>
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
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered"
	modal_title="Cloudflare"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div style="padding: 10px;">
            <select
            on:change={handleSelectGet} 
            class="form-control required" bind:value="{page_field}">
            <option value="1">PAGE 1</option>
            <option value="2">PAGE 2</option>
            <option value="3">PAGE 3</option>
            <option value="4">PAGE 4</option>
            <option value="5">PAGE 5</option>
            <option value="6">PAGE 6</option>
            <option value="7">PAGE 7</option>
            <option value="8">PAGE 8</option>
            <option value="9">PAGE 9</option>
            <option value="10">PAGE 10</option>
            <option value="11">PAGE 11</option>
            <option value="12">PAGE 12</option>
            <option value="13">PAGE 13</option>
            <option value="14">PAGE 14</option>
            <option value="15">PAGE 15</option>
            <option value="16">PAGE 16</option>
            <option value="17">PAGE 17</option>
            <option value="18">PAGE 18</option>
            <option value="19">PAGE 19</option>
            <option value="20">PAGE 20</option>
        </select>
        </div>
	</slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};"></th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">FILENAME</th>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">LINK</th>
                </tr>
            </thead>
            <tbody>
                {#each listalbumcloudflare as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                        {#if rec.album_signed == "LOCK"}
                        <i class="bi bi-lock-fill"></i>
                        {:else}
                        <i class="bi bi-unlock"></i>
                        {/if}
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.album_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.album_filename}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        <a href="{rec.album_link_0}" target="_blank">LINK</a>
                    </td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="DOWNLOAD"
            button_title="DOWNLOAD"
            button_css="btn-warning"/>
	</slot:template>
</Modal>



