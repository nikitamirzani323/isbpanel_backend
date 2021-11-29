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
    let title_page = "MOVIE - SERIES TROUBLE"
    let sData = "";
    let sDataRoot = "";
    let myModal = "";
    
    let listalbum = []
    let listgenre = []
    let listseason = []
    let listepisode = []
    let record = ""
    let totalrecordcategory = 0
    
    
    let genre_field_idrecord = 0;
    let genre_field_name = "";
    let genre_field_display = 0;

    let season_field_idrecord = 0;
    let season_field_idmovie = 0;
    let season_field_titlemovie = "";
    let season_field_name = "";
    let season_field_display = 0;

    let episode_field_idrecord = 0;
    let episode_field_seasonid = 0;
    let episode_field_titleseason = "";
    let episode_field_name = "";
    let episode_field_source = "";
    let episode_field_display = 0;

    let movie_field_idrecord = 0;
    let movie_field_title = "";
    let movie_field_label = "";
    let movie_field_slug = "";
    let movie_field_descp = "";
    let movie_field_urlvideo = "";
    let movie_field_genre = [];
    let movie_field_source_count = 0;
    let movie_field_source = [];
    let movie_field_series = [];
    let movie_field_episode = [];
    let movie_field_year = 0;
    let movie_field_imdb = 0.0;
    let movie_field_image = "";
    let movie_field_status = "0";

    let album_field_name = "";
    let searchMovie = "";
    let filterMovie = "";
    let album_flagclick = false;
    let genre_flagclick = false;
    let album_css = "";
    let genre_css = "";
    let css_loader = "display: none;";
    let msgloader = "";
    let pagingnow = 0;
    
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
    const handleSelectPaging = (event) => {
        let page = event.target.value
        pagingnow = page
        const movie = {
                page,
        };
        dispatch("handlePaging", movie);
    };
    
    const ShowGenre = (e) => {
        myModal = new bootstrap.Modal(document.getElementById("modalgenre"));
        myModal.show();
        genre_flagclick = e
        if(genre_flagclick){
            genre_css = "text-decoration:underline;color:blue;cursor:pointer;"
        }else{
            genre_css = "";
            genre_flagclick = false
        }
        call_genre()
    };
    const ShowAlbum = (e) => {
        myModal = new bootstrap.Modal(document.getElementById("modalalbum"));
        myModal.show();
        album_flagclick = e
        if(album_flagclick){
            album_css = "text-decoration:underline;color:blue;cursor:pointer;"
        }else{
            album_css = "";
            album_flagclick = false
        }
        call_album();
    };
    const ShowSeason = (e,title) => {
        myModal = new bootstrap.Modal(document.getElementById("modalseason"));
        myModal.show();
        call_season(e,title);
    };
    const ShowEpisode = (e,title) => {
        episode_field_seasonid = e
        episode_field_titleseason = title
        myModal = new bootstrap.Modal(document.getElementById("modalepisode"));
        myModal.show();
        call_episode(e);
    };
    const ShowFormAlbum = (e,id,name,display) => {
        sData = e
        if(e == "Edit"){
            genre_field_idrecord = parseInt(id);
            genre_field_name = name;
            genre_field_display = parseInt(display);
        }else{
            clearfield_genre()
        }
        
        myModal = new bootstrap.Modal(document.getElementById("modalcrudalbum"));
        myModal.show();
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
    const ShowFormSeason = (e,id,name,display) => {
        sData = e
        if(e == "Edit"){
            season_field_idrecord = parseInt(id);
            season_field_name = name;
            season_field_display = parseInt(display);
        }else{
            clearfield_season()
        }
        
        myModal = new bootstrap.Modal(document.getElementById("modalcrudseason"));
        myModal.show();
    };
    const ShowFormEpisode = (e,id,name,display) => {
        sData = e
        if(e == "Edit"){
            episode_field_idrecord = parseInt(id);
            episode_field_name = name;
            episode_field_display = parseInt(display);
        }else{
            clearfield_season()
        }
        
        myModal = new bootstrap.Modal(document.getElementById("modalcrudepisode"));
        myModal.show();
    };
    const ShowFormMovie = (e,id,title,label,descp,image,year,imdb,slug,status,genre) => {
        sDataRoot = e
        if(e == "Edit"){
            movie_field_idrecord = parseInt(id);
            movie_field_title = title;
            movie_field_label = label;
            movie_field_descp = descp;
            movie_field_image = image;
            movie_field_slug = slug;
            movie_field_year = parseInt(year);
            movie_field_imdb = parseFloat(imdb);
            movie_field_status = "0";
            if(status == "SHOW"){
                movie_field_status = "1";
            }
            movie_field_genre = [];
            for (var i = 0; i < genre.length; i++) {
                movie_field_genre = [
                    ...movie_field_genre,
                    {
                        movie_genre_id: genre[i]['moviegenre_id'],
                        movie_genre_name: genre[i]['moviegenre_name'],
                    },
                ];         
            }
        }else{
            clearfield_movie()
        }
        myModal = new bootstrap.Modal(document.getElementById("modalcrudmovie"));
        myModal.show();
    };
    
    async function call_album(){
        listalbum = []
        const res = await fetch("/api/moviecloudualbum", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"MOVIEALBUM-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status == true) {
            record = json.record;
            let images = record.images;
            let no = 0;
            if (record != null) {
                for (var i = 0; i < images.length; i++) {
                    let signed = ""
                    if(images[i]["requireSignedURLs"] == true){
                        signed = "LOCKED"
                    }
                    no = no + 1;
                    listalbum = [
                        ...listalbum,
                        {
                            album_no: no,
                            album_filename: images[i]["filename"],
                            album_id: images[i]["id"],
                            album_variant_0: images[i]["variants"][0],
                            album_signed: signed,
                        },
                    ];
                }
            }
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
    async function call_episode(e) {
        listepisode = [];
        episode_field_seasonid = e
        const res = await fetch("/api/movieseriesepisode", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                movie_id: parseInt(season_field_idmovie),
                season_id: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    let source = record[i]["movieepisode_source"]
                    if(record[i]["movieepisode_source"] == null){
                        source = []
                    }
                    no = no + 1;
                    listepisode = [
                        ...listepisode,
                        {
                            movieepisode_no: no,
                            movieepisode_id: record[i]["movieepisode_id"],
                            movieepisode_seasonid: record[i]["movieepisode_seasonid"],
                            movieepisode_name: record[i]["movieepisode_name"],
                            movieepisode_display: record[i]["movieepisode_display"],
                            movieepisode_source: source,
                        },
                    ];
                }
            }
        } 
    }
    async function call_season(e,title) {
        listseason = [];
        season_field_idmovie = e;
        season_field_titlemovie = title;
        const res = await fetch("/api/movieseriesseason", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                movie_id: e,
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
                    listseason = [
                        ...listseason,
                        {
                            movieseason_no: no,
                            movieseason_id: record[i]["movieseason_id"],
                            movieseason_name: record[i]["movieseason_name"],
                            movieseason_episodetotal: record[i]["movieseason_episodetotal"],
                            movieseason_display: record[i]["movieseason_display"],
                        },
                    ];
                }
            }
        } 
    }
    async function handleSaveEpisode() {
        let flag = true
        let msg = ""
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if(sData == "New"){
            if(episode_field_name == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(episode_field_source == ""){
                flag = false
                msg += "The Source is required\n"
            }
            if(episode_field_display == "" || parseInt(episode_field_display) == 0){
                flag = false
                msg += "The Display is required\n"
            }
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/movieseriesepisodesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"MOVIE-SAVE",
                    movie_page: parseInt(pagingnow),
                    movie_id: parseInt(season_field_idmovie),
                    movieseason_id: parseInt(episode_field_seasonid),
                    movieepisode_id: parseInt(episode_field_idrecord),
                    movieepisode_name: episode_field_name.toUpperCase(),
                    movieepisode_source: episode_field_source,
                    movieepisode_display: parseInt(episode_field_display),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                myModal.hide()
                call_episode(episode_field_seasonid)
                call_season(season_field_idmovie,season_field_titlemovie)
                clearfield_episode()
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
    async function handleSaveSeason() {
        let flag = true
        let msg = ""
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if(sData == "New"){
            if(season_field_name == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(season_field_display == "" || parseInt(season_field_display) == 0){
                flag = false
                msg += "The Display is required\n"
            }
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/movieseriesseasonsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"MOVIE-SAVE",
                    movie_page: parseInt(pagingnow),
                    movie_id: parseInt(season_field_idmovie),
                    movieseason_id: parseInt(season_field_idrecord),
                    movieseason_name: season_field_name.toUpperCase(),
                    movieseason_display: parseInt(season_field_display),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                myModal.hide()
                call_season(season_field_idmovie,season_field_titlemovie)
                clearfield_season()
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
                    page:"MOVIE-SAVE",
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
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/movieseriessave", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sDataRoot,
                page:"MOVIE-SAVE",
                movie_page: parseInt(pagingnow),
                movie_id: movie_field_idrecord,
                movie_name: movie_field_title,
                movie_label: movie_field_label,
                movie_slug: movie_field_slug,
                movie_tipe: "serie",
                movie_descp: movie_field_descp,
                movie_urlmovie: movie_field_image,
                movie_year: parseInt(movie_field_year),
                movie_imdb: parseFloat(movie_field_imdb),
                movie_status: parseInt(movie_field_status),
                movie_gender:movie_field_genre,
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
    async function handleNewMovieImage(file) {
        movie_field_image = file
        myModal.hide();
    }
    async function handleNewMovieGenre(id,name) {
        movie_field_genre = [
            ...movie_field_genre,
            {
                movie_genre_id: id,
                movie_genre_name: name,
            },
        ];
    }
    async function handleNewMovieSource() {
        if(movie_field_urlvideo != ""){
            movie_field_source_count = movie_field_source_count + 1
            movie_field_source = [
                ...movie_field_source,
                {
                    movie_source_id: parseInt(movie_field_source_count),
                    movie_source_name: movie_field_urlvideo,
                },
            ];
        }else{
            alert("The URL Video is required")
        }
        movie_field_urlvideo = ""
    }
    let files;
    async function handleNewCloudflare() {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const formData = new FormData();
        formData.append('sdata', sData);
        formData.append('page', "MOVIE-SAVE");
        formData.append('file', files[0]);
        const res = await fetch("/api/moviecloudupload", {
            method: "POST",
            headers: {
                Authorization: "Bearer " + token,
            },
            body: formData,
        });
        const json = await res.json();
        const status = json.status;
        if(status == true){
            album_field_name = ""
            myModal.hide();
            msgloader = "Success";
        }
        
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        call_album()
    }
    async function handleUpdateCloudflare(e,id) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
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
            }),
        });
        const json = await res.json();
        const status = json.status;
        if(status == true){
            call_album()
            msgloader = "Success";
        }
        
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        
    }
    async function handleDeleteCloudflare(id) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
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
                movie_id: id,
            }),
        });
        const json = await res.json();
        const status = json.status;
        if(status == true){
            call_album()
            msgloader = "Success";
        }
        
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        
    }
   
    async function handleDeleteMovieGenre(e) {
        let temp = movie_field_genre.filter(item => item.movie_genre_id !== parseInt(e))
        movie_field_genre = []
        for(var i=0;i<temp.length;i++){
            movie_field_genre = [
                ...movie_field_genre,
                {
                    movie_genre_id: parseInt(temp[i].movie_genre_id),
                    movie_genre_name: temp[i].movie_genre_name,
                },
            ];
        }
    }
    async function handleDeleteMovie(e) {
        let flag = true
        let msg = ""
        if(e == ""){
            flag = false
            msg = "The News is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/moviedelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"MOVIE-DELETE",
                    movie_id: parseInt(e),
                    movie_page: parseInt(pagingnow),
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
    async function handleDeleteSeason(e) {
        let flag = true
        let msg = ""
        if(e == "" || parseInt(e) < 1){
            flag = false
            msg = "The Season is required"
        }
        if(parseInt(season_field_idmovie) < 1){
            flag = false
            msg = "The Movie ID is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/movieseriesseasondelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"MOVIE-DELETE",
                    movie_page: parseInt(pagingnow),
                    movie_id: parseInt(season_field_idmovie),
                    movieseason_id: parseInt(e),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                call_season(season_field_idmovie,season_field_titlemovie)
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
    async function handleDeleteEpisode(e,idseason) {
        let flag = true
        let msg = ""
        if(e == "" || parseInt(e) < 1){
            flag = false
            msg = "The Season is required"
        }
        if(parseInt(season_field_idmovie) < 1){
            flag = false
            msg = "The Movie ID is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/movieseriesepisodedelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"MOVIE-DELETE",
                    movie_page: parseInt(pagingnow),
                    movie_id: parseInt(season_field_idmovie),
                    season_id: parseInt(idseason),
                    episode_id: parseInt(e),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                call_episode(idseason)
                listseason = [];
                call_season(season_field_idmovie,season_field_titlemovie)
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
    function callFunction(event){
        switch(event.detail){
            case "FORMNEW_EPISODE":
                ShowFormEpisode("New");break;
            case "SAVE_EPISODE":
                handleSaveEpisode();break;
            case "FORMNEW_SEASON":
                ShowFormSeason("New");break;
            case "SAVE_SEASON":
                handleSaveSeason();break;
            case "CALL_ALBUM":
                ShowAlbum();break;
            case "FORMNEW_ALBUM":
                ShowFormAlbum("New");break;
            case "CALL_GENRE":
                ShowGenre(false);break;
            case "FORM_MOVIE":
                ShowFormMovie("New");break;
            case "FORMNEW_GENRE":
                ShowFormGenre("New");break;
            case "SAVE_GENRE":
                handleSaveGenre();break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE_SOURCE":
                handleNewMovieSource();break;
            case "SAVE_CLOUDFLARE":
                handleNewCloudflare();break;
            case "SAVE_ALBUM":
                handleNewCloudflare();break;
            case "SAVE_MOVIE":
                handleSave();break;
        }
    }
    function clearfield_movie(){
        movie_field_idrecord = 0;
        movie_field_title = "";
        movie_field_label = "";
        movie_field_slug = "";
        movie_field_descp = "";
        movie_field_urlvideo = "";
        movie_field_genre = [];
        movie_field_source = [];
        movie_field_year = 0.0;
        movie_field_imdb = 0.0;
        movie_field_image = "";
        movie_field_status = "0";
    }
    function clearfield_genre(){
        genre_field_idrecord = 0;
        genre_field_name = "";
        genre_field_display = 0;
    }
    function clearfield_season(){
        season_field_idrecord = 0;
        season_field_name = "";
        season_field_display = 0;
    }
    function clearfield_episode(){
        episode_field_idrecord = 0;
        episode_field_name = "";
        episode_field_source = "";
        episode_field_display = 0;
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
    function uppercase(element) {
		function onInput(event) {
			element.value = element.value.toUpperCase();
		}
		element.addEventListener('input', onInput);
		return {
			destroy() {
				element.removeEventListener('input', onInput);
			}
		}
	}
    function lowercase(element) {
		function onInput(event) {
			element.value = element.value.toLowerCase();
		}
		element.addEventListener('input', onInput);
		return {
			destroy() {
				element.removeEventListener('input', onInput);
			}
		}
	}
    function popupwindow(e){
        window.open(e, "", "width=640, height=480");
    }
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
                button_function="CALL_ALBUM"
                button_title="Album"
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
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CDN</th>
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TYPE</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">YEAR</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">GENRE</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">SEASON</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">THUMBNAIL</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MOVIE</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">IMDB</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VIEW</th>
                                <th NOWRAP width="2%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">COMMENT</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterMovie as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                        on:click={() => {
                                            ShowFormMovie("Edit",rec.movie_id,rec.movie_title,rec.movie_label,rec.movie_descp,
                                            rec.movie_thumbnail,rec.movie_year,rec.movie_imdb,rec.movie_slug,rec.movie_status,rec.movie_genre)
                                        }} 
                                        class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        {#if rec.movie_season.length < 1}
                                        <i 
                                            on:click={() => {
                                                handleDeleteMovie(rec.movie_id);
                                            }} 
                                            class="bi bi-trash"></i>
                                        {/if}
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.movie_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.movie_statuscss}">{rec.movie_status}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        CREATE : {rec.movie_date}<br>
                                        UPDATE : {rec.movie_update}
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="{rec.movie_css_cdn}padding:5px 10px 5px 10px;">{rec.movie_imgcdn}</span>
                                    </td>
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
                                        {#each rec.movie_season as rec2}
                                            {rec2.movieseason_name} - {rec2.movieseason_episodetotal}<br>
                                        {/each}
                                        <button 
                                            on:click={() => {
                                                ShowSeason(rec.movie_id,rec.movie_title);
                                            }} 
                                            type="button" class="btn btn-info btn-sm">New Season</button>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <img width="50" class="img-thumbnail" src="{rec.movie_thumbnail}" alt="">
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {rec.movie_title}<br>
                                        <b>LABEL</b> : {rec.movie_label}<br>
                                        <b>SLUG</b> : {rec.movie_slug}
                                    </td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.movie_imdbcss}">{rec.movie_imdb}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.movie_viewcss}">{rec.movie_view}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.movie_commentcss}">{rec.movie_comment}</td>
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
                    <label for="exampleForm" class="form-label">Label</label>
                    <Input
                        bind:value={movie_field_label}
                        class="required"
                        type="text"
                        placeholder="Movie Label"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Slug</label>
                    <input
                        use:lowercase
                        bind:value={movie_field_slug}
                        class="required form-control"
                        type="text"
                        placeholder="Movie Slug"/>
                    <p class="text-muted">
                        Ex : dragon-ball-2021
                    </p>
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
                    <label for="exampleForm" class="form-label">Url Image</label>
                    <div class="input-group mb-3">
                        <Input
                            bind:value={movie_field_image}
                            class="required"
                            type="text"
                            placeholder="Movie URL Image"/>
                        <button
                            on:click={() => {
                                ShowAlbum(true);
                            }}  
                            type="button" class="btn btn-info">Album</button>
                    </div>
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
                    <label for="exampleForm" class="form-label">Genre - 
                        <span
                            on:click={() => {
                                ShowGenre(true)
                            }} 
                            style="text-decoration: underline;cursor:pointer;color:blue;">New</span>
                    </label>
                    <table class="table table-sm">
                        <tbody>
                            {#each movie_field_genre as rec }
                            <tr>
                                <td width="1%" style="cursor: pointer;text-align:center;vertical-align:top;">
                                    <i 
                                        on:click={() => {
                                            handleDeleteMovieGenre(rec.movie_genre_id);
                                        }} 
                                        class="bi bi-trash"></i>
                                </td>
                                <td width="*" style="text-align:left;vertical-align:top;font-size:12px;">{rec.movie_genre_name}</td>
                            </tr>       
                            {/each}
                        </tbody>
                    </table>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Preview</label><br>
                    <img class="img-thumbnail" src="{movie_field_image}" alt="">
                </div>
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_MOVIE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalalbum"
	modal_size="modal-dialog-centered"
	modal_title="Album"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" colspan="2">&nbsp;</th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">FILENAME</th>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">LINK</th>
                </tr>
            </thead>
            <tbody>
                {#each listalbum as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        {#if rec.album_signed == "LOCKED"}
                        <i 
                            on:click={() => {
                                handleUpdateCloudflare("UNLOCK",rec.album_id);
                            }} 
                            class="bi bi-lock-fill"></i>
                        {:else}
                        <i 
                            on:click={() => {
                                handleUpdateCloudflare("LOCK",rec.album_id);
                            }} 
                            class="bi bi-unlock"></i>
                        {/if}
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i 
                            on:click={() => {
                                handleDeleteCloudflare(rec.album_id);
                            }} 
                            class="bi bi-trash"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.album_no}</td>
                    {#if album_flagclick == true}
                        <td 
                            on:click={() => {
                                handleNewMovieImage(rec.album_variant_0);
                            }} 
                            NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};{album_css}">{rec.album_filename}</td>
                    {:else}
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.album_filename}</td>
                    {/if}
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        <a href="{rec.album_variant_0}" target="_blank">LINK</a>
                    </td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FORMNEW_ALBUM"
            button_title="New"
            button_css="btn-primary"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalcrudalbum"
	modal_size="modal-dialog-centered"
	modal_title="ALBUM/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <input id="fileUpload" type="file" bind:files>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_ALBUM"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalformsource"
	modal_size="modal-dialog-centered"
	modal_title="Source"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">URL Video</label>
			<Input
                bind:value={movie_field_urlvideo}
                class="required"
                type="text"
                placeholder="Url VIdeo"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_SOURCE"
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
	modal_footer={false}>
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
                    {#if genre_flagclick == true}
                        <td 
                            on:click={() => {
                                handleNewMovieGenre(rec.genre_id,rec.genre_name);
                            }} 
                            NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};{genre_css}">{rec.genre_name}</td>
                    {:else}
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};{genre_css}">{rec.genre_name}</td>
                    {/if}
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

<Modal
	modal_id="modalseason"
	modal_size="modal-dialog-centered"
	modal_title="{season_field_titlemovie}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" colspan="2">&nbsp;</th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SEASON</th>
                    <th width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">EPISODE</th>
                    <th width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DISPLAY</th>
                </tr>
            </thead>
            <tbody>
                {#each listseason as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i 
                            on:click={() => {
                                ShowFormSeason("Edit",rec.movieseason_id,rec.movieseason_name,rec.movieseason_display);
                            }} 
                            class="bi bi-pencil"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        {#if rec.movieseason_episodetotal < 1}
                        <i 
                            on:click={() => {
                                handleDeleteSeason(rec.movieseason_id);
                            }} 
                            class="bi bi-trash"></i>
                        {/if}
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.movieseason_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};{genre_css}">{rec.movieseason_name}</td>
                    <td 
                        on:click={() => {
                            ShowEpisode(rec.movieseason_id,rec.movieseason_name);
                        }}
                        NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};cursor:pointer;text-decoration:underline;color:blue;">{rec.movieseason_episodetotal}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.movieseason_display}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FORMNEW_SEASON"
            button_title="New"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalcrudseason"
	modal_size="modal-dialog-centered"
	modal_title="SEASON/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
			<input
                use:uppercase
                bind:value={season_field_name}
                class="required form-control"
                type="text"
                placeholder="Season Name"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Display</label>
			<Input
                bind:value={season_field_display}
                class="required"
                maxlength=3
                type="text"
                style="text-align:right;"
                placeholder="Season Display"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_SEASON"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalepisode"
	modal_size="modal-dialog-centered"
	modal_title="{episode_field_titleseason}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%">&nbsp;</th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">EPISODE</th>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                    <th width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DISPLAY</th>
                </tr>
            </thead>
            <tbody>
                {#each listepisode as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i 
                            on:click={() => {
                                handleDeleteEpisode(rec.movieepisode_id,rec.movieepisode_seasonid);
                            }} 
                            class="bi bi-trash"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.movieepisode_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};{genre_css}">{rec.movieepisode_name}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};cursor:pointer;text-decoration:underline;color:blue;">
                        {#each rec.movieepisode_source as rec2}
                            <span
                            style="text-decoration: underline;color:blue;cursor:pointer;" 
                            on:click={() => {
                                popupwindow(rec2.moviesource_url)
                            }} >{rec2.moviesource_stream}</span><br>
                        {/each}
                    </td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.movieepisode_display}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FORMNEW_EPISODE"
            button_title="New"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalcrudepisode"
	modal_size="modal-dialog-centered"
	modal_title="EPISODE/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
			<input
                use:uppercase
                bind:value={episode_field_name}
                class="required form-control"
                type="text"
                placeholder="Episode Name"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Url Source</label>
			<input
                bind:value={episode_field_source}
                class="required form-control"
                type="text"
                placeholder="Episode Source"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Display</label>
			<Input
                bind:value={episode_field_display}
                class="required"
                maxlength=3
                type="text"
                style="text-align:right;"
                placeholder="Episode Display"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_EPISODE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>