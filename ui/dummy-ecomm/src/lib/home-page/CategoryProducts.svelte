<!--
This component is used to display the best products from each category. 
  -->
<script>
	import { onMount } from 'svelte';
	export let filters = {
		category_id: null,
		limit: 10,
		sort: 'asc',
		keywords: '',
		min_price: 0,
		max_price: 10000
	};
	export let displayName = '';
	let products = [];
	onMount(async () => {
		console.log(filters);
		const queryParams = new URLSearchParams();

		if (filters.category_id) queryParams.set('category_id', filters.category_id);
		if (filters.limit) queryParams.set('limit', filters.limit);
		if (filters.sort) queryParams.set('sort', filters.sort);
		if (filters.keywords) queryParams.set('keywords', filters.keywords);
		if (filters.min_price) queryParams.set('min_price', filters.min_price);
		if (filters.max_price) queryParams.set('max_price', filters.max_price);

		const response = await fetch(`http://localhost:1323/products?${queryParams.toString()}`);
		const data = await response.json();
		products = data.products;
	});
</script>

{#if products}
	<div class="space-y-8">
		<div class="p-8">
			<h2 class="text-2xl font-bold mb-4 text-gray-900">{displayName}</h2>
			<div class="flex space-x-4 overflow-x-auto">
				{#each products as product}
					<div class="min-w-[200px] bg-white p-4 rounded-lg shadow-sm flex-shrink-0 border-2">
						<img
							src={'http://localhost:1323/static/images/' + product.thumbnail}
							alt=""
							class="w-full h-48 object-cover"
						/>
						<h3 class="text-lg font-semibold mt-2 text-gray-900">
							<a href={'/product/' + product.id}>{product.name}</a>
						</h3>
						<p class="text-gray-500 mt-1">₹{product.price}</p>
					</div>
				{/each}
			</div>
		</div>
	</div>
{/if}
