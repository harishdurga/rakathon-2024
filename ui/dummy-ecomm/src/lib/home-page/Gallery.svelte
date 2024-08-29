<script>
	import { Splide, SplideSlide } from '@splidejs/svelte-splide';
	import '@splidejs/svelte-splide/css';

	const images = [
		'https://placehold.co/1920x400',
		'https://placehold.co/1920x400',
		'https://placehold.co/1920x400'
	];
</script>

<Splide aria-label="My Favorite Images">
	{#each images as image}
		<SplideSlide>
			<img src={image} alt="" />
		</SplideSlide>
	{/each}
</Splide>

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
