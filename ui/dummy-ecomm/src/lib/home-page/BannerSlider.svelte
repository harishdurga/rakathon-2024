<script>
	import { Splide, SplideSlide } from '@splidejs/svelte-splide';
	import '@splidejs/svelte-splide/css';
	import { onMount } from 'svelte';
	import { createEventDispatcher } from 'svelte';
	export let categories = [];
	let banners = [];
	onMount(async () => {
		if (categories.length > 0) {
			const response = await fetch(
				`http://localhost:1323/banners?categories=${categories.join(',')}`
			);
			banners = await response.json();
		}
	});
</script>

{#if banners.length > 0}
	<Splide aria-label="My Favorite Images">
		{#each banners as banner}
			<SplideSlide>
				<img src={banner.url} alt="" />
			</SplideSlide>
		{/each}
	</Splide>
{:else}
	<p>No banners found</p>
{/if}

<style>
	.splide {
		width: 100%;
		height: 100%;
	}
	.splide__slide img {
		width: 100%;
		height: auto;
	}
</style>
