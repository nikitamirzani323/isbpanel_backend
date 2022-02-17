<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./components/Navbar.svelte";
	import Dashboard from "./pages/dashboard/Dashboard.svelte";
	import Admin from "./pages/admin/Admin.svelte";
	import Adminrule from "./pages/adminrule/Adminrule.svelte";
	import Pasaran from "./pages/pasaran/Pasaran.svelte";
	import Tafsirmimpi from "./pages/tafsirmimpi/Tafsirmimpi.svelte";
	import News from "./pages/news/News.svelte";
	import Movie from "./pages/movie/Movie.svelte";
	import Series from "./pages/series/Series.svelte";
	import Movienews from "./pages/movienews/Movienews.svelte";
	import Movietrouble from "./pages/movietrouble/Movietrouble.svelte";
	import Seriestrouble from "./pages/seriestrouble/Seriestrouble.svelte";
	import Domain from "./pages/domain/Domain.svelte";
	import Album from "./pages/album/Album.svelte";
	import Crm from "./pages/crm/Crm.svelte";
	import Crmisbtv from "./pages/crmisbtv/Crmisbtv.svelte";
	import Crmduniafilm from "./pages/crmduniafilm/Crmduniafilm.svelte";
	import Websiteagen from "./pages/websiteagen/Websiteagen.svelte";
	import Game from "./pages/game/Game.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/NotFound.svelte";
	export let table_header_font;
	export let table_body_font;
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
				component: Dashboard,
			}),
			"/webagen": wrap({
				component: Websiteagen,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/game": wrap({
				component: Game,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/crm": wrap({
				component: Crm,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/crmduniafilm": wrap({
				component: Crmduniafilm,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/crmisbtv": wrap({
				component: Crmisbtv,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/album": wrap({
				component: Album,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/seriesrouble": wrap({
				component: Seriestrouble,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/movietrouble": wrap({
				component: Movietrouble,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/domain": wrap({
				component: Domain,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/movienews": wrap({
				component: Movienews,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/series": wrap({
				component: Series,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/movie": wrap({
				component: Movie,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/pasaran": wrap({
				component: Pasaran,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/tafsir-mimpi": wrap({
				component: Tafsirmimpi,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/news": wrap({
				component: News,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/admin": wrap({
				component: Admin,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/adminrule": wrap({
				component: Adminrule,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navbar />
{/if}
<Router {routes} />
