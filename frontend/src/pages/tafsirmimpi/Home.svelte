<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "TAFSIR MIMPI"
    let sData = "";
    let myModal_newentry = "";
    let tafsir_field_mimpi = "";
    let tafsir_field_artimimpi = "";
    let tafsir_field_2d = "";
    let tafsir_field_3d = "";
    let tafsir_field_4d = "";
    let tafsir_field_status = "";
    let idrecord = "";
    let searchTafsirMimpi = "";
    let filterTafsirMimpi = [];
    let record = ""
    let record_message = ""
    let totalrecord_keluaran = ""
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchTafsirMimpi) {
            filterTafsirMimpi = listHome.filter(
                (item) =>
                    item.tafsirmimpi_mimpi
                        .toLowerCase()
                        .includes(searchTafsirMimpi.toLowerCase()) ||
                    item.tafsirmimpi_artimimpi
                        .toLowerCase()
                        .includes(searchTafsirMimpi.toLowerCase())
            );
        } else {
            filterTafsirMimpi = [...listHome];
        }
    }
    
    const NewData = () => {
        clearField()
        sData = "New"
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e,id,mimpi,artimimpi, nomor2d,nomor3d,nomor4d,status) => {
        sData = e
        idrecord = parseInt(id);
        tafsir_field_mimpi = mimpi;
        tafsir_field_artimimpi = artimimpi;
        tafsir_field_2d = nomor2d;
        tafsir_field_3d = nomor3d;
        tafsir_field_4d = nomor4d;
        if(tafsir_field_status == "SHOW"){
            status = "Y"
        }
        tafsir_field_status = status;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
    };
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(tafsir_field_mimpi == ""){
                flag = false
                msg += "The Mimpi is required\n"
            }
            if(tafsir_field_artimimpi == ""){
                flag = false
                msg += "The Arti Mimpi is required\n"
            }
            if(tafsir_field_2d == ""){
                flag = false
                msg += "The 2D is required\n"
            }
            if(tafsir_field_3d == ""){
                flag = false
                msg += "The 3D is required\n"
            }
            if(tafsir_field_4d == ""){
                flag = false
                msg += "The 4D is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(tafsir_field_mimpi == ""){
                flag = false
                msg += "The Mimpi is required\n"
            }
            if(tafsir_field_artimimpi == ""){
                flag = false
                msg += "The Arti Mimpi is required\n"
            }
            if(tafsir_field_2d == ""){
                flag = false
                msg += "The 2D is required\n"
            }
            if(tafsir_field_3d == ""){
                flag = false
                msg += "The 3D is required\n"
            }
            if(tafsir_field_4d == ""){
                flag = false
                msg += "The 4D is required\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/tafsirmimpisave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"TAFSIRMIMPI-SAVE",
                    tafsirmimpi_id: parseInt(idrecord),
                    tafsirmimpi_mimpi: tafsir_field_mimpi,
                    tafsirmimpi_artimimpi: tafsir_field_artimimpi,
                    tafsirmimpi_angka2d: tafsir_field_2d,
                    tafsirmimpi_angka3d: tafsir_field_3d,
                    tafsirmimpi_angka4d: tafsir_field_4d,
                    tafsirmimpi_status: tafsir_field_status,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sData=="New"){
                    clearField()
                }
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
    
    function clearField(){
        idrecord = "";
        tafsir_field_mimpi = "";
        tafsir_field_artimimpi = "";
        tafsir_field_2d = "";
        tafsir_field_3d = "";
        tafsir_field_4d = "";
        tafsir_field_status = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
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
                button_function="NEW"
                button_title="New"
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
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchTafsirMimpi}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Tafsir Mimpi + Enter"
                            aria-label="Search"
                        />
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped table-hover">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MIMPI</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterTafsirMimpi as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                EditData("Edit",rec.tafsirmimpi_id,
                                                    rec.tafsirmimpi_mimpi, rec.tafsirmimpi_artimimpi, 
                                                    rec.tafsirmimpi_angka2d, 
                                                    rec.tafsirmimpi_angka3d, 
                                                    rec.tafsirmimpi_angka4d,rec.tafsirmimpi_status);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.tafsirmimpi_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.tafsirmimpi_statuscss}">{rec.tafsirmimpi_status}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {rec.tafsirmimpi_mimpi}<br>
                                        {rec.tafsirmimpi_artimimpi}<br>
                                        4D - {rec.tafsirmimpi_angka4d}<br>
                                        3D - {rec.tafsirmimpi_angka3d}<br>
                                        2D - {rec.tafsirmimpi_angka2d}
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
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Mimpi</label>
                    <Input
                        bind:value={tafsir_field_mimpi}
                        class="required"
                        type="text"
                        placeholder="Mimpi"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Arti Mimpi</label>
                    <textarea
                        style="height: 100px;resize: none;" 
                        bind:value={tafsir_field_artimimpi} class="form-control"></textarea>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">2D</label>
                    <Input
                        bind:value={tafsir_field_2d}
                        class="required"
                        type="text"
                        placeholder="2D"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">3D</label>
                    <Input
                        bind:value={tafsir_field_3d}
                        class="required"
                        type="text"
                        placeholder="3D"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">4D</label>
                    <Input
                        bind:value={tafsir_field_4d}
                        class="required"
                        type="text"
                        placeholder="4D"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select 
                        class="form-control required"
                        bind:value={tafsir_field_status}>
                        <option value="Y">SHOW</option>
                        <option value="">HIDDEN</option>
                    </select>
                </div>
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>



