<!--
This component is used to display the best products from each category. 
  -->
<script>
	const categories = [
		{
			name: 'Electronics',
			products: [
				{
					name: 'OneMinus Bullets 2',
					image: 'https://placehold.co/400x400',
					price: 2400,
					id: 1
				},
				{
					name: 'Unrealme 7 Ultra Pro Max',
					image: 'https://placehold.co/400x400',
					price: 20000,
					id: 2
				},
				{
					name: 'Maomi 10000MAH Powerbank',
					image: 'https://placehold.co/400x400',
					price: 1500,
					id: 3
				},
				{
					name: 'Mawa Earbuds pro',
					image: 'https://placehold.co/400x400',
					price: 9500,
					id: 4
				},
				{
					name: 'Someday Battery Light',
					image: 'https://placehold.co/400x400',
					price: 500,
					id: 5
				}
			]
		},
		{
			name: "Men's Fashion",
			products: [
				{
					name: 'Jonex Shoes',
					image: 'https://placehold.co/400x400',
					price: 550,
					id: 6
				},
				{
					name: 'Mevis Jeans',
					image: 'https://placehold.co/400x400',
					price: 2000,
					id: 7
				},
				{
					name: 'Yolo T-Shirt',
					image: 'https://placehold.co/400x400',
					price: 600,
					id: 8
				},
				{
					name: 'I & N T-Shirt',
					image: 'https://placehold.co/400x400',
					price: 999,
					id: 9
				},
				{
					name: 'LowLander Jeans',
					image: 'https://placehold.co/400x400',
					price: 2199,
					id: 10
				}
			]
		},
		{
			name: 'Recently Viewed',
			products: [
				{
					name: 'I & N T-Shirt',
					image: 'https://placehold.co/400x400',
					price: 999,
					id: 9
				},
				{
					name: 'LowLander Jeans',
					image: 'https://placehold.co/400x400',
					price: 2199,
					id: 10
				},
				{
					name: 'Mawa Earbuds pro',
					image: 'https://placehold.co/400x400',
					price: 9500,
					id: 4
				}
			]
		}
	];
</script>

<div class="space-y-8">
	{#each categories as category}
		<div class="p-8">
			<h2 class="text-2xl font-bold mb-4 text-gray-900">Best From {category.name}</h2>
			<div class="flex space-x-4 overflow-x-auto">
				{#each category.products as product}
					<div class="min-w-[200px] bg-white p-4 rounded-lg shadow-sm flex-shrink-0 border-2">
						<img src={product.image} alt="" class="w-full h-48 object-cover" />
						<h3 class="text-lg font-semibold mt-2 text-gray-900">
							<a href={'/product/' + product.id}>{product.name}</a>
						</h3>
						<p class="text-gray-500 mt-1">₹{product.price}</p>
					</div>
				{/each}
			</div>
		</div>
	{/each}
</div>
