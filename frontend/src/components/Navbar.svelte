<script>
    import { link } from "svelte-spa-router";
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
</script>

<header>
    <nav
        class="navbar navbar-expand-md navbar-dark fixed-top bg-dark"
        style="background-color: #1f2937!important;">
        <div class="container-fluid">
            <a class="navbar-brand" href="/" style="font-size: 13px;">BACKEND</a>
            <button
                class="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarCollapse"
                aria-controls="navbarCollapse"
                aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon" />
            </button>
            <div class="collapse navbar-collapse" id="navbarCollapse">
                <ul class="navbar-nav me-auto mb-2 mb-md-0">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" use:link href="/" style="font-size: 13px;">DASHBOARD</a>
                    </li>
                    <li class="nav-item dropdown">
                        <a
                            class="nav-link dropdown-toggle"
                            id="navbarDropdown"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false" style="font-size: 13px;">
                            CRM
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                            <li>
                                <a class="dropdown-item" use:link href="/crm" style="font-size: 13px;">CRM</a>
                            </li>
                            <li>
                                <a class="dropdown-item" use:link href="/crmisbtv" style="font-size: 13px;">ISBTV</a>
                            </li>
                            <li>
                                <a class="dropdown-item" use:link href="/crmduniafilm" style="font-size: 13px;">DUNIAFILM</a>
                            </li>
                        </ul>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" aria-current="page" use:link href="/news" style="font-size: 13px;">NEWS</a>
                    </li>
                    <li class="nav-item dropdown">
                        <a
                            class="nav-link dropdown-toggle"
                            id="navbarDropdown"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false" style="font-size: 13px;">
                            MOVIE
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                            <li><a class="dropdown-item" use:link href="/movie" style="font-size: 13px;">MOVIE</a></li>
                            <li><a class="dropdown-item" use:link href="/series" style="font-size: 13px;">SERIES</a></li>
                            <li><a class="dropdown-item" use:link href="/movienews" style="font-size: 13px;">MOVIE TRAILER</a></li>
                            <li><a class="dropdown-item" use:link href="/movietrouble" style="font-size: 13px;">MOVIE TROUBLE</a></li>
                            <li><a class="dropdown-item" use:link href="/seriesrouble" style="font-size: 13px;">SERIES TROUBLE</a></li>
                        </ul>
                    </li>
                    <li class="nav-item dropdown">
                        <a
                            class="nav-link dropdown-toggle"
                            id="navbarDropdown"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false" style="font-size: 13px;">
                            KELUARAN
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                            <li>
                                <a class="dropdown-item" use:link href="/pasaran" style="font-size: 13px;">PASARAN</a>
                            </li>
                            <li>
                                <a class="dropdown-item" use:link href="/tafsir-mimpi" style="font-size: 13px;">TAFSIR MIMPI</a>
                            </li>
                        </ul>
                    </li>
                    <li class="nav-item dropdown">
                        <a
                            class="nav-link dropdown-toggle"
                            id="navbarDropdown"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false" style="font-size: 13px;">
                            MASTER
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                            <li>
                                <a class="dropdown-item" use:link href="/domain" style="font-size: 13px;">DOMAIN MANAGEMENT</a>
                                <a class="dropdown-item" use:link href="/game" style="font-size: 13px;">GAME MANAGEMENT</a>
                                <a class="dropdown-item" use:link href="/webagen" style="font-size: 13px;">WEBSITE MANAGEMENT</a>
                                <a class="dropdown-item" use:link href="/album" style="font-size: 13px;">ALBUM MANAGEMENT</a>
                                <a class="dropdown-item" use:link href="/admin" style="font-size: 13px;">ADMIN MANAGEMENT</a>
                                <a class="dropdown-item" use:link href="/adminrule" style="font-size: 13px;">ADMIN RULE</a>
                            </li>
                        </ul>
                    </li>
                </ul>
                <form class="d-flex">
                    <button 
                        on:click={() => {
                        handleLogout();
                        }}
                        class="btn btn-warning btn-sm">LOGOUT</button>
                </form>
            </div>
        </div>
    </nav>
</header>