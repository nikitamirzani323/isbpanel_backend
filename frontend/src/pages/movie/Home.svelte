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
	export let listPage = []
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
    let title_page = "MOVIE"
    let sData = "";
    let myModal = "";
    
    let listnews = []
    let listgenre = []
    let listsource = []
    let record = ""
    let totalrecordnews = 0
    let totalrecordcategory = 0
    
    let tanggal_start_newsfetch = "";
    let tanggal_end_newsfetch = "";
    let page_newsfetch = "";
    let genre_field_idrecord = 0;
    let genre_field_name = "";
    let genre_field_display = 0;
    let movie_field_idrecord = 0;
    let movie_field_title = "";
    let movie_field_descp = "";
    let movie_field_genre = [];
    let movie_field_year = 0.0;
    let movie_field_imdb = 0.0;
    let movie_field_image = "";
    let movie_field_cover = "";
    let movie_field_status = "0";
    let searchMovie = "";
    let filterMovie = "";

    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchMovie) {
            filterMovie = listHome.filter(
                (item) =>
                    item.movie_status
                        .toLowerCase()
                        .includes(searchMovie.toLowerCase()) || 
                    item.movie_title
                        .toLowerCase()
                        .includes(searchMovie.toLowerCase()) || 
                    item.movie_year
                        .toLowerCase()
                        .includes(searchMovie.toLowerCase())
            );
        } else {
            filterMovie = [...listHome];
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const ShowFormNewsFetch = () => {
        sData = "Edit"
        myModal = new bootstrap.Modal(document.getElementById("modalfetchnew"));
        myModal.show();
    };
    const ShowGenre = () => {
        sData = ""
        myModal = new bootstrap.Modal(document.getElementById("modalgenre"));
        myModal.show();
        call_genre()
    };
    const ShowFormGenre = (e,id,name,display) => {
        sData = e
        if(e == "Edit"){
            genre_field_idrecord = parseInt(id);
            genre_field_name = name;
            genre_field_display = parseInt(display);
        }else{
            clearfield_genre()
        }
        
        myModal = new bootstrap.Modal(document.getElementById("modalcrudgenre"));
        myModal.show();
    };
    const ShowFormMovie = (e,id,category,title,descp,url,image) => {
        sData = e
        if(e == "Edit"){
            movie_field_idrecord = parseInt(id);
            movie_field_title = title;
            movie_field_descp = descp;
            movie_field_genre = parseInt(category);
            movie_field_image = image;
        }else{
            clearfield_movie()
        }
        
        myModal = new bootstrap.Modal(document.getElementById("modalcrudmovie"));
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
    async function call_genre() {
        listgenre = [];
        const res = await fetch("/api/genremovie", {
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
            if (record != null) {
                totalrecordcategory = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listgenre = [
                        ...listgenre,
                        {
                            genre_no: no,
                            genre_id: record[i]["genre_id"],
                            genre_name: record[i]["genre_name"],
                            genre_display: record[i]["genre_display"],
                            genre_create: record[i]["genre_create"],
                            genre_update: record[i]["genre_update"],
                        },
                    ];
                }
            }
        } 
    }
    async function handleSaveGenre() {
        let flag = true
        let msg = ""
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if(sData == "New"){
            if(genre_field_name == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(genre_field_display == ""){
                flag = false
                msg += "The Display is required\n"
            }
        }
        if(flag){
            
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/genremoviesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"MOVIEGENRE-SAVE",
                    genre_id: parseInt(genre_field_idrecord),
                    genre_name: genre_field_name.toUpperCase(),
                    genre_display: parseInt(genre_field_display),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                myModal.hide()
                call_genre()
                clearfield_genre()
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
    async function handleSave() {
        let flag = true
        let msg = ""
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/newssave", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"NEWS-SAVE",
                news_id: news_field_idrecord,
                news_category: news_field_category,
                news_title: news_field_title,
                news_descp: news_field_descp,
                news_url: news_field_url,
                news_image: news_field_image,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            myModal.hide()
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
    async function handleDeleteNews(e) {
        let flag = true
        let msg = ""
        if(e == ""){
            flag = false
            msg = "The News is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/newsdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"NEWS-DELETE",
                    news_id: parseInt(e),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                RefreshHalaman()
                msgloader = json.message;
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
    async function handleDeleteCategoryNews(e) {
        let flag = true
        let msg = ""
        if(e == ""){
            flag = false
            msg = "The Category is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/categorynewsdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"CATEGORYNEWS-DELETE",
                    category_id: parseInt(e),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                call_category()
                msgloader = json.message;
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
    function callFunction(event){
        switch(event.detail){
            case "CALL_FORMNEWS":
                ShowFormNewsFetch();
                break;
            case "FETCH_NEWS":
                call_news();
                break;
            case "CALL_GENRE":
                ShowGenre();
                break;
            case "FORM_MOVIE":
                ShowFormMovie("New");
                break;
            case "FORMNEW_GENRE":
                ShowFormGenre("New");
                break;
            case "SAVE_GENRE":
                handleSaveGenre();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE_NEWS":
                handleSave();break;
        }
    }
    function clearfield_movie(){
        movie_field_idrecord = 0;
        movie_field_title = "";
        movie_field_descp = "";
        movie_field_genre = "";
        movie_field_image = "";
    }
    function clearfield_genre(){
        genre_field_idrecord = 0;
        genre_field_name = "";
        genre_field_display = 0;
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterMovie = [];
                listHome = [];
                const movie = {
                    searchMovie,
                };
                dispatch("handleMovie", movie);
        }  
    };
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="FORM_MOVIE"
                button_title="New Movie"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="FORM_SERIES"
                button_title="New Series"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="CALL_GENRE"
                button_title="Genre"
                button_css="btn-primary"/>
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
                            style="text-align: center;" 
                            class="form-control">
                            {#each listPage as rec}
                                <option value="{rec.page_id}">{rec.page_display}</option>
                            {/each}
                        </select>
                    </div>
                </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchMovie}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Movie + Tekan Enter"
                            aria-label="Search"
                        />
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped table-hover">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TYPE</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">YEAR</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">GENRE</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">THUMBNAIL</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MOVIE</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">IMDB</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VIEW</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterMovie as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                ShowFormNews("Edit",rec.news_id,rec.news_idcategory,rec.news_title,rec.news_descp,rec.news_url,rec.news_image)
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                handleDeleteNews(rec.news_id);
                                            }} 
                                            class="bi bi-trash"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.movie_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.movie_statuscss}">{rec.movie_status}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.movie_date}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="{rec.movie_csstype}padding:5px 10px 5px 10px;">{rec.movie_type}</span>
                                    </td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.movie_year}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {#each rec.movie_genre as rec2}
                                            {rec2.moviegenre_name}<br>
                                        {/each}
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <img width="50" class="img-thumbnail" src="{rec.movie_thumbnail}" alt="">
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.movie_title}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.movie_imdb}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.movie_view}</td>
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
	modal_id="modalcrudmovie"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Movie/{sData}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Movie</label>
                    <Input
                        bind:value={movie_field_title}
                        class="required"
                        type="text"
                        placeholder="Movie Title"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Deskripsi</label>
                    <textarea
                        style="height: 100px;resize: none;" 
                        bind:value={movie_field_descp} class="form-control required"></textarea>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Year</label>
                    <Input
                        bind:value={movie_field_year}
                        minlength=4
                        maxlength=4
                        style="text-align:right;"
                        class="required"
                        type="text"
                        placeholder="Movie Imdb"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Imdb</label>
                    <Input
                        bind:value={movie_field_imdb}
                        class="required"
                        style="text-align:right;"
                        type="text"
                        placeholder="Movie Imdb"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url Thumbnail</label>
                    <Input
                        bind:value={movie_field_image}
                        class="required"
                        type="text"
                        placeholder="Movie URL Thumbnail"/>
                    <a href="https://id.imgbb.com/" target="_blank">imgbb</a>, 
                    <a href="https://imgur.com/" target="_blank">imgur</a>

                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url Cover</label>
                    <Input
                        bind:value={movie_field_cover}
                        class="required"
                        type="text"
                        placeholder="Movie URL Cover"/>
                    <a href="https://id.imgbb.com/" target="_blank">imgbb</a>,
                    <a href="https://imgur.com/" target="_blank">imgur</a>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select  
                        bind:value={movie_field_status}
                        class="form-control required">
                        <option value="1">SHOW</option>
                        <option value="0">HIDE</option>
                    </select>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Genre - <span style="text-decoration: underline;cursor:pointer;color:blue;">New</span></label>
                    <table>
                        {#each listgenre as rec }
                        <tr>
                            <td>{rec.name}</td>
                        </tr>       
                        {/each}
                    </table>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Source - <span style="text-decoration: underline;cursor:pointer;color:blue;">New</span></label>
                    <table>
                        {#each listsource as rec }
                        <tr>
                            <td>{rec.url}</td>
                        </tr>       
                        {/each}
                    </table>
                </div>
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_NEWS"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalgenre"
	modal_size="modal-dialog-centered"
	modal_title="GENRE"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" colspan="2">&nbsp;</th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">GENRE</th>
                    <th width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DISPLAY</th>
                </tr>
            </thead>
            <tbody>
                {#each listgenre as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i 
                            on:click={() => {
                                ShowFormGenre("Edit",rec.genre_id,rec.genre_name,rec.genre_display);
                            }} 
                            class="bi bi-pencil"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i 
                            on:click={() => {
                                handleDeleteCategoryNews(rec.genre_id);
                            }} 
                            class="bi bi-trash"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.genre_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.genre_name}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.genre_display}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FORMNEW_GENRE"
            button_title="New"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalcrudgenre"
	modal_size="modal-dialog-centered"
	modal_title="GENRE/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
			<Input
                bind:value={genre_field_name}
                class="required"
                type="text"
                placeholder="Genre Name"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Display</label>
			<Input
                bind:value={genre_field_display}
                class="required"
                maxlength=3
                type="text"
                style="text-align:right;"
                placeholder="Genre Display"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_GENRE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>