<!-- YOU CAN DELETE EVERYTHING IN THIS PAGE -->
<script>
	import BannerSlider from '$lib/home-page/BannerSlider.svelte';
	import CategoryProducts from '$lib/home-page/CategoryProducts.svelte';
	import { onMount } from 'svelte';
	let pageLayoutJSON = null;
	onMount(async () => {
		const res = await fetch('http://localhost:1323/home-page-layout');
		pageLayoutJSON = await res.json();
		console.log(pageLayoutJSON);
	});
</script>

{#if pageLayoutJSON}
	{#each pageLayoutJSON.layout as layoutItem}
		{#if layoutItem.type === 'main-banner-slider'}
			<BannerSlider categories={layoutItem.filters.categories.split(',').map(Number)} />
		{:else if layoutItem.type === 'category-products'}
			<CategoryProducts filters={layoutItem.filters} displayName={layoutItem.display_title} />
		{/if}
	{/each}
{/if}
