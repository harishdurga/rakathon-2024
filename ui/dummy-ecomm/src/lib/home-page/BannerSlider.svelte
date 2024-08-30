<script>
	import { Splide, SplideSlide } from '@splidejs/svelte-splide';
	import '@splidejs/svelte-splide/css';
	import { onMount } from 'svelte';
	export let categories = [];
	let banners = null;
	onMount(async () => {
		if (categories.length > 0) {
			const response = await fetch(
				`http://localhost:1323/banners?category_ids=${categories.join(',')}`
			);
			banners = await response.json();
			console.log(banners);
		}
	});
</script>

{#if banners && banners.banners && banners.banners.length > 0}
	<Splide aria-label="My Favorite Images">
		{#each banners.banners as banner}
			<SplideSlide>
				<img src={'http://localhost:1323/static/images/' + banner.url} alt="" />
			</SplideSlide>
		{/each}
	</Splide>
{:else}
	<p>No banners found</p>
{/if}
