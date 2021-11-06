<script>
    import { Input } from "sveltestrap";
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, set } from "firebase/database";
    import dayjs from "dayjs";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    const firebaseConfig = {
        apiKey: "AIzaSyCnjwV66P7jDLx5A0Hlh7CHKoZ2tg9jmMY",
        authDomain: "united-rope-233010.firebaseapp.com",
        databaseURL: "https://united-rope-233010-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "united-rope-233010",
        storageBucket: "united-rope-233010.appspot.com",
        messagingSenderId: "994050756260",
        appId: "1:994050756260:web:4dee40c4ca0c34a1842031"
    };
    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);
	export let table_header_font
	export let table_body_font
	export let token
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "SDSB4D - NIGHT"
    let sData = "";
    let myModal_newentry = "";
    let tanggal_keluaran = "";
    let date_keluaran ="";
    let idrecord = 0;
    let prize1 = "";
    let prize2 = "";
    let prize3 = "";
    let prize1_flag = false;
    let prize2_flag = false;
    let prize3_flag = false;
    let prize1_save_flag = false;
    let prize2_save_flag = false;
    let prize3_save_flag = false;
    let css_loader = "display: none;";
    let msgloader = "";

    console.log(dayjs().format("DD/MM/YYYY"))
    
    const NewData = () => {
        clearField()
        sData = "New"
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e,tanggal,prize_1,prize_2,prize_3) => {
        sData = "Edit"
        idrecord = parseInt(e)
        prize1 = prize_1
        prize2 = prize_2
        prize3 = prize_3
        if(prize_1 !=""){
            prize1_flag = true;
            prize1_save_flag = true;
        }else{
            prize1_flag = false;
            prize1_save_flag = false;
        }
        if(prize_2 !=""){
            prize2_flag = true;
            prize2_save_flag = true;
        }else{
            prize2_flag = false;
            prize2_save_flag = false;
        }
        if(prize_3 !=""){
            prize3_flag = true;
            prize3_save_flag = true;
        }else{
            prize3_flag = false;
            prize3_save_flag = false;
        }
        tanggal_keluaran = tanggal;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentryedit"));
        myModal_newentry.show();
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(date_keluaran == ""){
            flag = false
            msg = "The Date is required"
        }
        if(flag){
            const res = await fetch("/api/savesdsbnight", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"SDSB4DNIGHT-SAVE",
                    idrecord: parseInt(0),
                    tanggal: date_keluaran,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                set(ref(db, 'sdsb4dnight'), {
                    datedraw: dayjs(date_keluaran).format("DD-MMM-YYYY"),
                    nextdraw: dayjs(date_keluaran).add(1,'day').format("YYYY-MM-DD"),
                    prize1: "",
                    prize2: "",
                    prize3: "",
                });

                msgloader = json.message;
                myModal_newentry.hide()
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
    async function handleSaveGenerator(tipe,prize) {
        const res = await fetch("/api/savegeneratorsdsbnight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"SDSB4DNIGHT-SAVE",
                idrecord: parseInt(idrecord),
                tipe: tipe,
                prize: prize.toString(),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            RefreshHalaman()
            switch(tipe){
                case "prize1":
                    prize1_save_flag = true;
                    prize1_flag = true;
                    set(ref(db, 'sdsb4dnight'), {
                        datedraw: dayjs(tanggal_keluaran).format("DD-MMM-YYYY"),
                        nextdraw: dayjs(tanggal_keluaran).add(1,'day').format("YYYY-MM-DD"),
                        prize1: prize,
                        prize2: prize2,
                        prize3: prize3,
                    });
                    
                    break;
                case "prize2":
                    prize2_save_flag = true;
                    prize2_flag = true;
                    set(ref(db, 'sdsb4dnight'), {
                        datedraw: dayjs(tanggal_keluaran).format("DD-MMM-YYYY"),
                        nextdraw: dayjs(tanggal_keluaran).add(1,'day').format("YYYY-MM-DD"),
                        prize1: prize1,
                        prize2: prize,
                        prize3: prize3,
                    });
                    break;
                case "prize3":
                    prize3_save_flag = true;
                    prize3_flag = true;
                    set(ref(db, 'sdsb4dnight'), {
                        datedraw: dayjs(tanggal_keluaran).format("DD-MMM-YYYY"),
                        nextdraw: dayjs(tanggal_keluaran).add(1,'day').format("YYYY-MM-DD"),
                        prize1: prize1,
                        prize2: prize2,
                        prize3: prize,
                    });
                    break;
                    break;
            }
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function handleGeneratorAutomation(tipe,prize) {
        const res = await fetch("/api/generatornumbernight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
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
    }
    function clearField(){
        date_keluaran = ""
    }
    function generate(field){
        let numbergenerate = (Math.floor(Math.random() * 10000) + 10000).toString().substring(1);
        switch(field){
            case "prize1":
                prize1 = numbergenerate
                break;
            case "prize2":
                prize2 = numbergenerate
                break;
            case "prize3":
                prize3 = numbergenerate
                break;
        }
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "GENERATOR":
                handleGeneratorAutomation();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_format = () => {
		let numbera;
		for (let i = 0; i < prize1.length; i++) {
			numbera = parseInt(prize1[i]);
			if (isNaN(numbera)) {
				prize1 = "";
			}
		}
        for (let i = 0; i < prize2.length; i++) {
			numbera = parseInt(prize2[i]);
			if (isNaN(numbera)) {
				prize2 = "";
			}
		}
        for (let i = 0; i < prize3.length; i++) {
			numbera = parseInt(prize3[i]);
			if (isNaN(numbera)) {
				prize3 = "";
			}
		}
    }
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
            <Button
                on:click={callFunction}
                button_function="GENERATOR"
                button_title="Generator"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 1</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 2</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 3</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each listHome as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    EditData(rec.sdsbnight_id,rec.sdsbnight_date,rec.sdsbnight_prize1,rec.sdsbnight_prize2,rec.sdsbnight_prize3);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_no}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_date}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_prize1}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_prize2}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_prize3}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_create}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.sdsbnight_update}</td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="10">
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
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Date</label>
			<Input
                bind:value={date_keluaran}
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
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

<Modal
	modal_id="modalentryedit"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Date</label>
			<Input
                bind:value={tanggal_keluaran}
                type="text"
                name="date"
                id="exampleDate"
                disabled
                placeholder="Tanggal"/>
		</div>
        
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Prize 1</label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize1}
                    on:keyup={handleKeyboard_format}    
                    disabled='{prize1_flag}'
                    type="text"
                    minlength="4"
                    maxlength="4"
                    placeholder="Prize 1"/>
                <button
                    on:click={() => {
                        generate("prize1");
                    }}  
                    disabled='{prize1_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize1",prize1);
                    }} 
                    disabled='{prize1_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Prize 2</label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize2}
                    on:keyup={handleKeyboard_format}
                    disabled='{prize2_flag}'
                    type="text"
                    minlength="4"
                    maxlength="4"
                    placeholder="Prize 2"/>
                <button
                    on:click={() => {
                        generate("prize2");
                    }} 
                    disabled='{prize2_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize2",prize2);
                    }} 
                    disabled='{prize2_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Prize 3</label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize3}
                    on:keyup={handleKeyboard_format}
                    disabled='{prize3_flag}'
                    type="text"
                    minlength="4"
                    maxlength="4"
                    placeholder="Prize 3"/>
                <button
                    on:click={() => {
                        generate("prize3");
                    }} 
                    disabled='{prize3_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize3",prize3);
                    }} 
                    disabled='{prize3_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
	</slot:template>
</Modal>